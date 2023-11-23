package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebpushServerControllerInterface interface {
	WebpushHandler(ginContext *gin.Context)
}

type WebpushServerController struct {
	ctx					 context.Context
	logger				 *logrus.Logger
	webpushServerService *service.WebpushServerService
}

func NewWebpushServerController(ctx context.Context, logger *logrus.Logger, webpushServerService *service.WebpushServerService) *WebpushServerController {
	return &WebpushServerController{
		ctx:				  ctx,
		logger:				  logger,
		webpushServerService: webpushServerService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/webpush/handler' | jq
// @Summary webpush handler
// @Description webpush handler
// @Tags webpush
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webpush/handler [get]
func (wsc *WebpushServerController) WebpushHandler(ginContext *gin.Context) {
	wsc.logger.Info("WebpushHandler")
	// wsc.webpushServerService.WebpushHandler(ginContext)
}
