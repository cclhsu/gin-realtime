package controller

import (
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type ServerSentEventServerControllerInterface interface {
}

type ServerSentEventServerController struct {
	// ctx	  context.Context
	logger                       *logrus.Logger
	ServerSentEventServerService *service.ServerSentEventServerService
}

func NewServerSentEventServerController(logger *logrus.Logger, ServerSentEventServerService *service.ServerSentEventServerService) *ServerSentEventServerController {
	return &ServerSentEventServerController{
		// ctx:	   ctx,
		logger:                       logger,
		ServerSentEventServerService: ServerSentEventServerService,
	}
}
