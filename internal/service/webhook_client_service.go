package service

import (
	"context"
	"fmt"
	"os"

	_interface "github.com/cclhsu/gin-realtime/internal/interface"
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/types"
	"github.com/cclhsu/gin-realtime/internal/utils"
	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
)

type WebhookClientInterface interface {
	NewWebhookClientService(ctx context.Context, logger *logrus.Logger) *WebhookClientService
	_interface.RegistrationServiceInterface
	_interface.MessageServiceInterface
	// _interface.HealthServiceInterface
	Initialize() error
	initializeWebhookServerServiceURL() string
	initializeWebhookClientURL() string
}

type WebhookClientService struct {
	ctx                     context.Context
	logger                  *logrus.Logger
	webhookServerServiceURL string
	clientID                string
	registeredWebhooks      []model.RegistrationDTO
}

func NewWebhookClientService(ctx context.Context, logger *logrus.Logger) *WebhookClientService {
	return &WebhookClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (wcs *WebhookClientService) Initialize() error {
	wcs.logger.Info("WebhookClientService Initialize")
	wcs.webhookServerServiceURL = wcs.initializeWebhookServerServiceURL()
	wcs.logger.Infof("Webhook Server URL: %s\n", wcs.webhookServerServiceURL)

	wcs.clientID = uuid.New().String()
	webhookRegistration := model.RegistrationDTO{
		UUID:        wcs.clientID,
		Type:        types.MESSAGE_TYPES_REGISTRATION,
		Stage:       types.STAGE_TYPES_UNSPECIFIED,
		Environment: types.ENVIRONMENT_TYPES_UNSPECIFIED,
		Sender:      wcs.clientID,
		// Timestamp:     time.Now(),
		CallbackURL:   wcs.initializeWebhookClientURL(),
		Subscriptions: []string{},
		// Expires:       time.Now().Add(1 * time.Hour * 24 * 7), // 1 week
		Secret: "1234567890",
		State:  types.GENERAL_STATE_TYPES_UNSPECIFIED,
	}

	wcs.registeredWebhooks = []model.RegistrationDTO{
		webhookRegistration,
	}
	// Initialize webhook server URL and register the webhook
	wcs.Register(webhookRegistration)
	return nil
}

func (wcs *WebhookClientService) initializeWebhookServerServiceURL() string {
	SERVER_HOST := os.Getenv("SERVER_HOST")
	if SERVER_HOST == "" {
		SERVER_HOST = "0.0.0.0"
	}
	SERVER_PORT := os.Getenv("SERVER_PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "3001"
	}
	return fmt.Sprintf("http://%s:%s/webhook", SERVER_HOST, SERVER_PORT)
}

func (wcs *WebhookClientService) initializeWebhookClientURL() string {
	CLIENT_HOST := os.Getenv("SERVICE_HOST")
	if CLIENT_HOST == "" {
		CLIENT_HOST = "0.0.0.0"
	}
	CLIENT_PORT := os.Getenv("SERVICE_PORT")
	if CLIENT_PORT == "" {
		CLIENT_PORT = "3002"
	}
	return fmt.Sprintf("http://%s:%s/webhook-client/message/receive", CLIENT_HOST, CLIENT_PORT)
}

func (wcs *WebhookClientService) Register(registration model.RegistrationDTO) (model.RegistrationResponseDTO, error) {
	wcs.logger.Infof("Registering webhook: %+v\n", registration)

	response, err := utils.SendRequest(wcs.logger, wcs.webhookServerServiceURL+"/register", "POST", registration)
	if err != nil {
		wcs.logger.Errorf("Webhook registration error: %s\n", err.Error())
		return model.RegistrationResponseDTO{}, err
	}

	regResponse, ok := model.ConvertToRegistrationResponseDTO(response)
	if ok != nil {
		wcs.logger.Errorf("unexpected response type: %T\n", response)
		return model.RegistrationResponseDTO{}, fmt.Errorf("unexpected response type: %T", response)
	}

	wcs.logger.Infof("Webhook registration response: %+v\n", regResponse)
	return regResponse, nil
}

func (wcs *WebhookClientService) Unregister(registrationID string) error {
	wcs.logger.Infof("Unregistering webhook: %+v\n", registrationID)

	response, err := utils.SendRequest(wcs.logger, wcs.webhookServerServiceURL+"/register/"+registrationID, "DELETE", nil)
	if err != nil {
		wcs.logger.Errorf("Webhook unregistration error: %s\n", err.Error())
		return err
	}

	wcs.logger.Infof("Webhook unregistration response: %+v\n", response)
	return nil
}

func (wcs *WebhookClientService) ListRegistrations() ([]model.RegistrationDTO, error) {
	wcs.logger.Infof("Listing registered webhooks")

	response, err := utils.SendRequest(wcs.logger, wcs.webhookServerServiceURL+"/register", "GET", nil)
	if err != nil {
		wcs.logger.Errorf("Webhook listing error: %s\n", err.Error())
		return nil, err
	}
	wcs.logger.Infof("Webhook listing response: %+v\n", response)

	// Type assertion to get the underlying slice
	var registrations []model.RegistrationDTO
	if responseSlice, ok := response.([]interface{}); ok {
		// Convert []interface{} to []model.RegistrationDTO
		for _, item := range responseSlice {
			registration, err := model.ConvertToRegistrationDTO(item)
			if err != nil {
				fmt.Println("Error converting to RegistrationDTO:", err)
				return nil, err
			}
			registrations = append(registrations, registration)
		}

		// Now registrations is []model.RegistrationDTO
		fmt.Println(registrations)
		wcs.registeredWebhooks = registrations
	} else if responseSlice, ok := response.([]model.RegistrationDTO); ok {
		// Type assertion to get the underlying slice
		wcs.registeredWebhooks = responseSlice
	} else {
		wcs.logger.Errorf("Unexpected type in Webhook listing response: %T", response)
		return nil, fmt.Errorf("Unexpected type in Webhook listing response: %T", response)
	}

	wcs.logger.Infof("Webhook listing response: %+v\n", wcs.registeredWebhooks)
	return wcs.registeredWebhooks, nil
}

func (wcs *WebhookClientService) UpdateRegistration(registrationID string, request model.RegistrationDTO) (model.RegistrationResponseDTO, error) {
	wcs.logger.Infof("Updating webhook: %+v\n", registrationID)

	request = model.RegistrationDTO{
		UUID:        registrationID,
		Type:        types.MESSAGE_TYPES_REGISTRATION,
		Stage:       types.STAGE_TYPES_UNSPECIFIED,
		Environment: types.ENVIRONMENT_TYPES_UNSPECIFIED,
		Sender:      registrationID,
		// Timestamp:     time.Now(),
		CallbackURL:   wcs.initializeWebhookClientURL(),
		Subscriptions: []string{},
		// Expires:       time.Now().Add(1 * time.Hour * 24 * 7), // 1 week
	}

	response, err := utils.SendRequest(wcs.logger, wcs.webhookServerServiceURL+"/register/"+registrationID, "PUT", request)
	if err != nil {
		wcs.logger.Errorf("Webhook update error: %s\n", err.Error())
		return model.RegistrationResponseDTO{}, err
	}

	wcs.logger.Infof("Webhook update response: %+v\n", response)
	return response.(model.RegistrationResponseDTO), nil
}

func (wcs *WebhookClientService) Send(message model.MessageDTO) error {
	wcs.logger.Infof("Sending message: %+v\n", message)

	response, err := utils.SendRequest(wcs.logger, wcs.webhookServerServiceURL+"/message/send", "POST", message)
	if err != nil {
		wcs.logger.Errorf("Message sending error: %s\n", err.Error())
		return err
	}

	wcs.logger.Infof("Message sending response: %+v\n", response)
	return nil
}

func (wcs *WebhookClientService) Receive(message model.MessageDTO) error {
	wcs.logger.Infof("Receiving message: %+v\n", message)

	response, err := utils.SendRequest(wcs.logger, wcs.webhookServerServiceURL+"/message/receive", "POST", message)
	if err != nil {
		wcs.logger.Errorf("Message receiving error: %s\n", err.Error())
		return err
	}

	wcs.logger.Infof("Message receiving response: %+v\n", response)
	return nil
}

func (wcs *WebhookClientService) ListMessages() ([]model.MessageDTO, error) {
	wcs.logger.Infof("Listing messages")

	response, err := utils.SendRequest(wcs.logger, wcs.webhookServerServiceURL+"/message", "GET", nil)
	if err != nil {
		wcs.logger.Errorf("Message listing error: %s\n", err.Error())
		return nil, err
	}
	wcs.logger.Infof("Message listing response: %+v\n", response)

	// Type assertion to get the underlying slice
	var messageList []model.MessageDTO
	if responseSlice, ok := response.([]interface{}); ok {
		// Convert []interface{} to []model.MessageDTO
		for _, item := range responseSlice {
			message, err := model.ConvertToMessageDTO(item)
			if err != nil {
				fmt.Println("Error converting to MessageDTO:", err)
				return nil, err
			}
			messageList = append(messageList, message)
		}

		// Now messageList is []model.MessageDTO
		fmt.Println(messageList)
	} else if responseSlice, ok := response.([]model.MessageDTO); ok {
		// Type assertion to get the underlying slice
		messageList = responseSlice
	} else {
		wcs.logger.Errorf("Unexpected type in Message listing response: %T", response)
		return nil, fmt.Errorf("Unexpected type in Message listing response: %T", response)
	}

	wcs.logger.Infof("Message listing response: %+v\n", messageList)
	return messageList, nil
}
