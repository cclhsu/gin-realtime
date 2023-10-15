package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/utils"
	"github.com/google/uuid"

	// "github.com/cclhsu/gin-realtime/internal/utils"
	"github.com/sirupsen/logrus"
)

type WebhookClientInterface interface {
	RegisterWebhook(webhookData model.WebhookInfoDTO) (model.WebhookRegistrationResponseDTO, error)
	TriggerEvent(eventDataDTO model.EventDataDTO) (model.EventDataResponseDTO, error)
	HandleTriggeredEvent(eventDataDTO model.EventDataDTO) (model.EventDataResponseDTO, error)
	ListRegisteredWebhooks() ([]model.WebhookInfoDTO, error)
	DisconnectWebhook(webhookId string) (model.WebhookRegistrationResponseDTO, error)
}

type WebhookClientService struct {
	ctx                     context.Context
	logger                  *logrus.Logger
	webhookServerServiceURL string
	// registeredWebhooks []model.WebhookInfoDTO
}

func NewWebhookClientService(ctx context.Context, logger *logrus.Logger) *WebhookClientService {
	return &WebhookClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (wcs *WebhookClientService) Initialize() {
	wcs.logger.Info("WebhookClientService Initialize")
	wcs.webhookServerServiceURL = wcs.initializewebhookServerServiceURL()
	wcs.logger.Infof("Webhook Server URL: %s\n", wcs.webhookServerServiceURL)

	webhookData := model.WebhookInfoDTO{
		ID:         uuid.New().String(),
		URL:        wcs.initializeWebhookClientURL(),
		ExpiryDate: time.Now(),
		Config: model.WebhookConfigDTO{
			Secret:   "1234567890",
			IsActive: true,
			Type:     model.Test,
		},
	}
	// wcs.registeredWebhooks = []model.WebhookInfoDTO{
	//	webhookData,
	// }
	// Initialize webhook server URL and register the webhook
	wcs.RegisterWebhook(webhookData)
}

func (wcs *WebhookClientService) initializewebhookServerServiceURL() string {
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
	return fmt.Sprintf("http://%s:%s/webhook-client/handle-payload", CLIENT_HOST, CLIENT_PORT)
}

func (wcs *WebhookClientService) RegisterWebhook(webhookData model.WebhookInfoDTO) (model.WebhookRegistrationResponseDTO, error) {
	wcs.logger.Infof("Registering webhook: %+v\n", webhookData)

	// Simulate sending the request
	// In a real implementation, you would make an HTTP POST request here
	response, err := utils.SendRequest(wcs.logger, wcs.webhookServerServiceURL+"/register", "POST", webhookData)
	if err != nil {
		wcs.logger.Errorf("Webhook registration error: %s\n", err.Error())
		return model.WebhookRegistrationResponseDTO{}, err
	}

	regResponse, ok := response.(model.WebhookRegistrationResponseDTO)
	if !ok {
		return model.WebhookRegistrationResponseDTO{}, fmt.Errorf("unexpected response type: %T", response)
	}

	wcs.logger.Infof("Webhook registration response: %+v\n", regResponse)
	return regResponse, nil
}

func (wcs *WebhookClientService) TriggerEvent(eventDataDTO model.EventDataDTO) (model.EventDataResponseDTO, error) {
	wcs.logger.Infof("Triggering event: %+v\n", eventDataDTO)

	response, err := utils.SendRequest(wcs.logger, wcs.webhookServerServiceURL+"/handle-event", "POST", eventDataDTO)
	if err != nil {
		wcs.logger.Errorf("Event triggering error: %s\n, %+v", err.Error(), response)
		return model.EventDataResponseDTO{}, err
	}

	wcs.logger.Infof("Event triggering response: %+v\n", response)

	eventDataResponse, ok := response.(model.EventDataResponseDTO)
	if !ok {
		wcs.logger.Errorf("unexpected response type: %T\n, %+v", response, response)
		return model.EventDataResponseDTO{}, fmt.Errorf("unexpected response type: %T", response)
	}

	return eventDataResponse, nil
}

func (wcs *WebhookClientService) HandleTriggeredEvent(eventDataDTO model.EventDataDTO) (model.EventDataResponseDTO, error) {
	wcs.logger.Infof("Handling triggered event: %+v\n", eventDataDTO)

	return model.EventDataResponseDTO{
		Success: true,
		Message: fmt.Sprintf("Event %s handled successfully", eventDataDTO.ID),
		Data:    nil,
	}, nil
}

func (ws *WebhookClientService) ListRegisteredWebhooks() ([]model.WebhookInfoDTO, error) {
	ws.logger.Infof("Listing registered webhooks\n")

	response, err := utils.SendRequest(ws.logger, ws.webhookServerServiceURL+"/list", "GET", nil)
	if err != nil {
		ws.logger.Errorf("Webhook listing error: %s\n", err.Error())
		return nil, err
	}

	webhookList, ok := response.([]model.WebhookInfoDTO)
	if !ok {
		ws.logger.Errorf("unexpected response type: %T\n", webhookList)
		return nil, fmt.Errorf("unexpected response type: %T", response)
	}

	ws.logger.Infof("Webhook listing response: %+v\n", webhookList)
	return webhookList, nil
}

func (wcs *WebhookClientService) DisconnectWebhook(webhookID string) (model.WebhookRegistrationResponseDTO, error) {
	wcs.logger.Infof("Disconnecting webhook: %s\n", webhookID)

	response, err := utils.SendRequest(wcs.logger, wcs.webhookServerServiceURL+"/unregister/"+webhookID, "DELETE", nil)
	if err != nil {
		wcs.logger.Errorf("Webhook disconnection error: %s\n", err.Error())
		return model.WebhookRegistrationResponseDTO{}, err
	}

	wcs.logger.Infof("Webhook disconnection response: %+v\n", response)
	return response.(model.WebhookRegistrationResponseDTO), nil
}
