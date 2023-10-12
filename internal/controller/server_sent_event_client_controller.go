package controller

import (
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type ServerSentEventClientControllerInterface interface {
}

type ServerSentEventClientController struct {
	// ctx	  context.Context
	logger                       *logrus.Logger
	ServerSentEventClientService *service.ServerSentEventClientService
}

func NewServerSentEventClientController(logger *logrus.Logger, ServerSentEventClientService *service.ServerSentEventClientService) *ServerSentEventClientController {
	return &ServerSentEventClientController{
		// ctx:	   ctx,
		logger:                       logger,
		ServerSentEventClientService: ServerSentEventClientService,
	}
}
