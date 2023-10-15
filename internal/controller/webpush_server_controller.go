package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type WebpushServerControllerInterface interface {
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
