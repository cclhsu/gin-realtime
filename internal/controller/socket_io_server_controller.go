package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SocketIOServerControllerInterface interface {
	SocketIOHandler(ginContext *gin.Context)
}

type SocketIOServerController struct {
	ctx					  context.Context
	logger				  *logrus.Logger
	socketIOServerService *service.SocketIOServerService
}

func NewSocketIOServerController(ctx context.Context, logger *logrus.Logger, socketIOServerService *service.SocketIOServerService) *SocketIOServerController {
	return &SocketIOServerController{
		ctx:				   ctx,
		logger:				   logger,
		socketIOServerService: socketIOServerService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/socketio/handler' | jq
// @Summary socketio handler
// @Description socketio handler
// @Tags socketio
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /socket-io/handler [get]
func (wsc *SocketIOServerController) SocketIOHandler(ginContext *gin.Context) {
	wsc.logger.Info("SocketIOHandler")
	// wsc.socketIOServerService.SocketIOHandler(ginContext)
}
