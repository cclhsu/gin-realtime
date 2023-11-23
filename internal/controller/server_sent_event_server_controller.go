package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ServerSentEventServerControllerInterface interface {
	ServerSentEventHandler(ginContext *gin.Context)
}

type ServerSentEventServerController struct {
	ctx							 context.Context
	logger						 *logrus.Logger
	serverSentEventServerService *service.ServerSentEventServerService
}

func NewServerSentEventServerController(ctx context.Context, logger *logrus.Logger, serverSentEventServerService *service.ServerSentEventServerService) *ServerSentEventServerController {
	return &ServerSentEventServerController{
		ctx:						  ctx,
		logger:						  logger,
		serverSentEventServerService: serverSentEventServerService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/server-sent-event/handler' | jq
// @Summary server-sent-event handler
// @Description server-sent-event handler
// @Tags server-sent-event
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /server-sent-event/handler [get]
func (wsc *ServerSentEventServerController) ServerSentEventHandler(ginContext *gin.Context) {
	wsc.logger.Info("ServerSentEventHandler")
	// wsc.ServerSentEventServerService.ServerSentEventHandler(ginContext)
}
