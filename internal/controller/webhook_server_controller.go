package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type WebhookServerControllerInterface interface {
}

type WebhookServerController struct {
	// ctx                  context.Context
	logger               *logrus.Logger
	webhookServerService *service.WebhookServerService
}

func NewWebhookServerController(ctx context.Context, logger *logrus.Logger, webhookServerService *service.WebhookServerService) *WebhookServerController {
	return &WebhookServerController{
		// ctx:                  ctx,
		logger:               logger,
		webhookServerService: webhookServerService,
	}
}
