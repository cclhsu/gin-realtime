package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebRTCServerControllerInterface interface {
	WebRTCHandler(ginContext *gin.Context)
}

type WebRTCServerController struct {
	ctx					context.Context
	logger				*logrus.Logger
	webrtcServerService *service.WebRTCServerService
}

func NewWebRTCServerController(ctx context.Context, logger *logrus.Logger, webrtcServerService *service.WebRTCServerService) *WebRTCServerController {
	return &WebRTCServerController{
		ctx:				 ctx,
		logger:				 logger,
		webrtcServerService: webrtcServerService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/webrtc/handler' | jq
// @Summary webrtc handler
// @Description webrtc handler
// @Tags webrtc
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webrtc/handler [get]
func (wsc *WebRTCServerController) WebRTCHandler(ginContext *gin.Context) {
	wsc.logger.Info("WebRTCHandler")
	// wsc.webrtcServerService.WebRTCHandler(ginContext)
}
