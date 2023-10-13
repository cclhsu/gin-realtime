package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type WebpushClientControllerInterface interface {
}

type WebpushClientController struct {
	ctx                  context.Context
	logger               *logrus.Logger
	webpushClientService *service.WebpushClientService
}

func NewWebpushClientController(ctx context.Context, logger *logrus.Logger, webpushClientService *service.WebpushClientService) *WebpushClientController {
	return &WebpushClientController{
		ctx:                  ctx,
		logger:               logger,
		webpushClientService: webpushClientService,
	}
}
