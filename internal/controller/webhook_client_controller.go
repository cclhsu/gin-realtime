package controller

import (
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type WebhookClientControllerInterface interface {
}

type WebhookClientController struct {
	// ctx	  context.Context
	logger               *logrus.Logger
	webhookClientService *service.WebhookClientService
}

func NewWebhookClientController(logger *logrus.Logger, webhookClientService *service.WebhookClientService) *WebhookClientController {
	return &WebhookClientController{
		// ctx:	   ctx,
		logger:               logger,
		webhookClientService: webhookClientService,
	}
}
