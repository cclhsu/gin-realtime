package service

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/cclhsu/gin-realtime/internal/model"
)

type WebhookClientInterface interface {
	RegisterWebhook(webhookData model.WebhookInfoDTO) (model.WebhookRegistrationResponseDTO, error)
	TriggerEvent(eventDataDTO model.EventDataDTO) (model.EventDataResponseDTO, error)
	HandleTriggeredEvent(eventDataDTO model.EventDataDTO) (model.EventDataResponseDTO, error)
	DisconnectWebhook(webhookId string) (model.WebhookRegistrationResponseDTO, error)
}

type WebhookClientService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewWebhookClientService(ctx context.Context, logger *logrus.Logger) *WebhookClientService {
	return &WebhookClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (wcs *WebhookClientService) RegisterWebhook(webhookData model.WebhookInfoDTO) (model.WebhookRegistrationResponseDTO, error) {
	return model.WebhookRegistrationResponseDTO{}, nil
}

func (wcs *WebhookClientService) TriggerEvent(eventDataDTO model.EventDataDTO) (model.EventDataResponseDTO, error) {
	return model.EventDataResponseDTO{}, nil
}

func (wcs *WebhookClientService) HandleTriggeredEvent(eventDataDTO model.EventDataDTO) (model.EventDataResponseDTO, error) {
	return model.EventDataResponseDTO{}, nil
}

func (wcs *WebhookClientService) DisconnectWebhook(webhookId string) (model.WebhookRegistrationResponseDTO, error) {
	return model.WebhookRegistrationResponseDTO{}, nil
}
