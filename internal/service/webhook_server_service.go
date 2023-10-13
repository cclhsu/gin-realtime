package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/utils"
	"github.com/sirupsen/logrus"
)

type WebhookServerInterface interface {
	RegisterWebhook(request model.WebhookInfoDTO) (model.WebhookRegistrationResponseDTO, error)
	HandleWebhookEvent(payload model.EventDataDTO) (model.WebhookRegistrationResponseDTO, error)
	UnregisterWebhook(webhookId string) (model.WebhookRegistrationResponseDTO, error)
	ListRegisteredWebhooks() ([]model.WebhookInfoDTO, error)
	UpdateWebhook(webhookId string, request model.WebhookInfoDTO) (model.WebhookRegistrationResponseDTO, error)
}

type WebhookServerService struct {
	ctx                   context.Context
	logger                *logrus.Logger
	WebhookConfigs        sync.Map
	EventDataDTOs         sync.Map
	WebhookConnections    sync.Map
	WebhookDisconnections sync.Map
}

func NewWebhookServerService(ctx context.Context, logger *logrus.Logger) *WebhookServerService {
	return &WebhookServerService{
		ctx:    ctx,
		logger: logger,
	}
}

func (wss *WebhookServerService) RegisterWebhook(request model.WebhookInfoDTO) (model.WebhookRegistrationResponseDTO, error) {
	wss.logger.Infof("Registering webhook: %+v\n", request)
	// Implement webhook registration logic here in the webhook server
	wss.WebhookConfigs.Store(request.ID, request)
	wss.WebhookConnections.Store(request.ID, true)
	wss.WebhookDisconnections.Store(request.ID, false)

	return model.WebhookRegistrationResponseDTO{
		Success: true,
		Message: fmt.Sprintf("Webhook %s registered successfully", request.ID),
		Data:    request,
	}, nil
}

func (wss *WebhookServerService) HandleWebhookEvent(payload model.EventDataDTO) (model.WebhookRegistrationResponseDTO, error) {
	wss.logger.Infof("Handling webhook event: %+v\n", payload)
	// Implement webhook event handling logic here in the webhook server
	wss.EventDataDTOs.Store(payload.ID, payload)
	postRequests := []chan struct{}{}

	// Loop through webhookConfigs and create an array of channels
	wss.WebhookConfigs.Range(func(key, value interface{}) bool {
		webhookConfigInfo := value.(model.WebhookInfoDTO)

		if webhookConfigInfo.Config.Type == model.Exclusive {
			if !webhookConfigInfo.Config.IsActive {
				return true
			}
			if webhookConfigInfo.ID != payload.ID {
				return true
			}
			disconnected, _ := wss.WebhookDisconnections.Load(webhookConfigInfo.ID)
			if disconnected.(bool) {
				return true
			}
		}

		webhookURL := webhookConfigInfo.URL
		wss.logger.Infof("Sending webhook event to: %s with payload: %+v\n", webhookURL, payload)

		postChan := make(chan struct{})

		// Define a function to send the request and handle retries
		go func(url string, payload model.EventDataDTO, retries int) {
			defer close(postChan)
			for retries > 0 {
				_, err := utils.SendRequest(wss.logger, url, "POST", payload)
				if err == nil {
					return
				}
				retries--
				time.Sleep(time.Second)
			}
		}(webhookURL, payload, 3)

		postRequests = append(postRequests, postChan)

		return true
	})

	// Wait for all post requests to complete
	for _, ch := range postRequests {
		<-ch
	}

	return model.WebhookRegistrationResponseDTO{
		Success: true,
		Message: fmt.Sprintf("Webhook event %s received successfully", payload.ID),
		Data:    nil,
	}, nil
}

func (wss *WebhookServerService) UnregisterWebhook(webhookID string) (model.WebhookRegistrationResponseDTO, error) {
	wss.logger.Infof("Unregistering webhook: %+v\n", webhookID)
	// Implement webhook unregistration logic here in the webhook server
	wss.WebhookConfigs.Delete(webhookID)
	wss.WebhookConnections.Delete(webhookID)
	wss.WebhookDisconnections.Delete(webhookID)

	return model.WebhookRegistrationResponseDTO{
		Success: true,
		Message: fmt.Sprintf("Webhook %s unregistered successfully", webhookID),
		Data:    nil,
	}, nil
}

func (wss *WebhookServerService) ListRegisteredWebhooks() ([]model.WebhookInfoDTO, error) {
	wss.logger.Infof("Listing registered webhooks\n")
	// Implement webhook list logic here in the webhook server
	webhookList := []model.WebhookInfoDTO{}
	wss.WebhookConfigs.Range(func(key, value interface{}) bool {
		configInfo := value.(model.WebhookInfoDTO)
		webhookList = append(webhookList, configInfo)
		return true
	})

	return webhookList, nil
}

func (wss *WebhookServerService) UpdateWebhook(webhookID string, request model.WebhookInfoDTO) (model.WebhookRegistrationResponseDTO, error) {
	wss.logger.Infof("Updating webhook: %+v\n", webhookID)
	// Implement webhook update logic here in the webhook server
	wss.WebhookConfigs.Store(webhookID, request)

	return model.WebhookRegistrationResponseDTO{
		Success: true,
		Message: fmt.Sprintf("Webhook %s updated successfully", webhookID),
		Data:    nil,
	}, nil
}
