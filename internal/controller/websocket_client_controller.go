package controller

import (
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type WebsocketClientControllerInterface interface {
}

type WebsocketClientController struct {
	// ctx	  context.Context
	logger                 *logrus.Logger
	websocketClientService *service.WebsocketClientService
}

func NewWebsocketClientController(logger *logrus.Logger, websocketClientService *service.WebsocketClientService) *WebsocketClientController {
	return &WebsocketClientController{
		// ctx:	   ctx,
		logger:                 logger,
		websocketClientService: websocketClientService,
	}
}
