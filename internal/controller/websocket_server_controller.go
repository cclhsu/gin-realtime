package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebsocketServerControllerInterface interface {
	NewWebsocketServerController(ctx context.Context, logger *logrus.Logger, websocketServerService *service.WebsocketServerService) *WebsocketServerController
	WebsocketHandler(ginContext *gin.Context)
}

type WebsocketServerController struct {
	// ctx					  context.Context
	logger                 *logrus.Logger
	websocketServerService *service.WebsocketServerService
}

func NewWebsocketServerController(ctx context.Context, logger *logrus.Logger, websocketServerService *service.WebsocketServerService) *WebsocketServerController {
	return &WebsocketServerController{
		// ctx:					   ctx,
		logger:                 logger,
		websocketServerService: websocketServerService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/websocket/ws' | jq
// @Summary websocket handler
// @Description websocket handler
// @Tags websocket
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /websocket/ws [get]
func (wsc *WebsocketServerController) WebsocketHandler(ginContext *gin.Context) {
	wsc.logger.Info("WebsocketHandler")
	wsc.websocketServerService.WebsocketHandler(ginContext)
}
