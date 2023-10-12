package service

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/model"
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
	ctx    context.Context
	logger *logrus.Logger
}

func NewWebhookServerService(ctx context.Context, logger *logrus.Logger) *WebhookServerService {
	return &WebhookServerService{
		ctx:    ctx,
		logger: logger,
	}
}

func (wss *WebhookServerService) RegisterWebhook(request model.WebhookInfoDTO) (model.WebhookRegistrationResponseDTO, error) {
	return model.WebhookRegistrationResponseDTO{}, nil
}

func (wss *WebhookServerService) HandleWebhookEvent(payload model.EventDataDTO) (model.WebhookRegistrationResponseDTO, error) {
	return model.WebhookRegistrationResponseDTO{}, nil
}

func (wss *WebhookServerService) UnregisterWebhook(webhookId string) (model.WebhookRegistrationResponseDTO, error) {
	return model.WebhookRegistrationResponseDTO{}, nil
}

func (wss *WebhookServerService) ListRegisteredWebhooks() ([]model.WebhookInfoDTO, error) {
	return []model.WebhookInfoDTO{}, nil
}

func (wss *WebhookServerService) UpdateWebhook(webhookId string, request model.WebhookInfoDTO) (model.WebhookRegistrationResponseDTO, error) {
	return model.WebhookRegistrationResponseDTO{}, nil
}
