package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	_interface "github.com/cclhsu/gin-realtime/internal/interface"
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/types"
	"github.com/sirupsen/logrus"
)

type WebhookServerInterface interface {
	NewWebhookServerService(ctx context.Context, logger *logrus.Logger) *WebhookServerService
	_interface.RegistrationServiceInterface
	_interface.MessageServiceInterface
	// _interface.HealthServiceInterface
	// Initialize() error
	broadcast(message model.MessageDTO) error
	echo(message model.MessageDTO) error
	unicast(message model.MessageDTO) error
	sendRequest(url string, payload model.MessageDTO) error
}

type WebhookServerService struct {
	ctx                   context.Context
	logger                *logrus.Logger
	webhookRegistrations  sync.Map
	MessageDataDTOs       sync.Map
	WebhookConnections    sync.Map
	WebhookDisconnections sync.Map
}

func NewWebhookServerService(ctx context.Context, logger *logrus.Logger) *WebhookServerService {
	return &WebhookServerService{
		ctx:    ctx,
		logger: logger,
	}
}

// func (wss *WebhookServerService) Initialize() error {
// 	wss.logger.Infof("Initializing webhook server")
// 	return errors.New("Method not implemented.")
// }

func (wss *WebhookServerService) Register(registration model.RegistrationDTO) (model.RegistrationResponseDTO, error) {
	wss.logger.Infof("Registering webhook: %+v\n", registration)

	wss.webhookRegistrations.Store(registration.UUID, registration)
	wss.WebhookConnections.Store(registration.UUID, true)
	wss.WebhookDisconnections.Store(registration.UUID, false)

	wss.logger.Infof("Webhook registered successfully: %+v\n", registration)
	return model.RegistrationResponseDTO{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Webhook registered successfully",
		Data: map[string]interface{}{
			"registration": registration,
		},
	}, nil
}

func (wss *WebhookServerService) Unregister(registrationID string) error {
	wss.logger.Infof("Unregistering webhook: %+v\n", registrationID)

	wss.webhookRegistrations.Delete(registrationID)
	wss.WebhookConnections.Delete(registrationID)
	wss.WebhookDisconnections.Delete(registrationID)

	wss.logger.Infof("Webhook unregistered successfully: %+v\n", registrationID)
	return nil
}

func (wss *WebhookServerService) ListRegistrations() ([]model.RegistrationDTO, error) {
	wss.logger.Infof("Listing registered webhooks")

	registrationList := []model.RegistrationDTO{}
	wss.webhookRegistrations.Range(func(key, value interface{}) bool {
		registration := value.(model.RegistrationDTO)
		registrationList = append(registrationList, registration)
		return true
	})

	wss.logger.Infof("Registration list: %+v\n", registrationList)
	return registrationList, nil
}

func (wss *WebhookServerService) UpdateRegistration(registrationID string, request model.RegistrationDTO) (model.RegistrationResponseDTO, error) {
	wss.logger.Infof("Updating webhook: %+v\n", registrationID)

	wss.webhookRegistrations.Store(registrationID, request)

	wss.logger.Infof("Webhook updated successfully: %+v\n", registrationID)
	return model.RegistrationResponseDTO{
		Status:  "success",
		Code:    http.StatusOK,
		Message: "Webhook updated successfully",
		Data: map[string]interface{}{
			"registration": request,
		},
	}, nil
}

func (wss *WebhookServerService) Send(message model.MessageDTO) error {
	wss.logger.Infof("Sending message: %+v\n", message)

	wss.MessageDataDTOs.Store(message.UUID, message)

	switch message.Type {
	case types.MESSAGE_TYPES_BROADCAST:
		wss.logger.Infof("Broadcasting message: %+v\n", message)
		wss.broadcast(message)
		break
	case types.MESSAGE_TYPES_ECHO:
		wss.logger.Infof("Echoing message: %+v\n", message)
		wss.echo(message)
		break
	case types.MESSAGE_TYPES_UNICAST:
		wss.logger.Infof("Unicasting message: %+v\n", message)
		wss.unicast(message)
		break
	default:
		wss.logger.Infof("Unknown message type: %+v\n", message)
		break
	}

	wss.logger.Infof("Message sent successfully: %+v\n", message)
	return nil
}

func (wss *WebhookServerService) Receive(message model.MessageDTO) error {
	wss.logger.Infof("Receiving message: %+v\n", message)

	wss.logger.Infof("Message received successfully: %+v\n", message)
	return nil
}

func (wss *WebhookServerService) ListMessages() ([]model.MessageDTO, error) {
	wss.logger.Infof("Listing messages")

	messageList := []model.MessageDTO{}
	wss.MessageDataDTOs.Range(func(key, value interface{}) bool {
		message := value.(model.MessageDTO)
		messageList = append(messageList, message)
		return true
	})

	wss.logger.Infof("Message list: %+v\n", messageList)
	return messageList, nil
}

func (wss *WebhookServerService) broadcast(message model.MessageDTO) error {
	wss.logger.Infof("Broadcasting message: %+v\n", message)
	defer func() {
		if r := recover(); r != nil {
			wss.logger.Errorf("[broadcast] Recovered from panic: %v", r)
		}
	}()

	wss.MessageDataDTOs.Store(message.UUID, message)

	var wg sync.WaitGroup

	wss.webhookRegistrations.Range(func(key, value interface{}) bool {
		webhookRegistration, ok := value.(model.RegistrationDTO)
		if !ok {
			wss.logger.Error("Invalid webhook configuration")
			return true
		}

		wg.Add(1)

		go func(url string, msg model.MessageDTO) {
			defer wg.Done()

			wss.logger.Infof("[broadcast] Sending webhook event to: %s with payload: %v", url, msg)

			if err := wss.sendRequest(url, msg); err != nil {
				wss.logger.Errorf("[broadcast] Webhook error: %v", err)
			}
		}(webhookRegistration.CallbackURL, message)

		return true
	})

	wg.Wait()

	return nil
}

func (wss *WebhookServerService) echo(message model.MessageDTO) error {
	wss.logger.Infof("Echoing message: %+v\n", message)
	defer func() {
		if r := recover(); r != nil {
			wss.logger.Errorf("[echo] Recovered from panic: %v", r)
		}
	}()

	webhookRegistration, ok := wss.webhookRegistrations.Load(message.Sender)
	if !ok {
		wss.logger.Errorf("Webhook %s not registered", message.Sender)
		return fmt.Errorf("Webhook %s not registered", message.Sender)
	}

	wss.logger.Infof("[echo] Sending webhook event to: %s with payload: %v", webhookRegistration.(model.RegistrationDTO).CallbackURL, message)

	return wss.sendRequest(webhookRegistration.(model.RegistrationDTO).CallbackURL, message)
}

func (wss *WebhookServerService) unicast(message model.MessageDTO) error {
	wss.logger.Infof("Unicasting message: %+v\n", message)
	defer func() {
		if r := recover(); r != nil {
			wss.logger.Errorf("[unicast] Recovered from panic: %v", r)
		}
	}()

	webhookRegistration, ok := wss.webhookRegistrations.Load(message.Recipient)
	if !ok {
		wss.logger.Errorf("Webhook %s not registered", message.Recipient)
		return fmt.Errorf("Webhook %s not registered", message.Recipient)
	}

	wss.logger.Infof("[unicast] Sending webhook event to: %s with payload: %v", webhookRegistration.(model.RegistrationDTO).CallbackURL, message)

	return wss.sendRequest(webhookRegistration.(model.RegistrationDTO).CallbackURL, message)
}

func (wss *WebhookServerService) sendRequest(url string, payload model.MessageDTO) error {
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", strings.NewReader(string(reqBody)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error! Status: %d", resp.StatusCode)
	}

	return nil
}
