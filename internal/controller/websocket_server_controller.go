package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type WebsocketServerControllerInterface interface {
}

type WebsocketServerController struct {
	// ctx                    context.Context
	logger                 *logrus.Logger
	websocketServerService *service.WebsocketServerService
}

func NewWebsocketServerController(ctx context.Context, logger *logrus.Logger, websocketServerService *service.WebsocketServerService) *WebsocketServerController {
	return &WebsocketServerController{
		// ctx:                    ctx,
		logger:                 logger,
		websocketServerService: websocketServerService,
	}
}
