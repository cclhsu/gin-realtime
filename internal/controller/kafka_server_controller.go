package controller

import (
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type KafkaServerControllerInterface interface {
}

type KafkaServerController struct {
	// ctx	  context.Context
	logger             *logrus.Logger
	kafkaServerService *service.KafkaServerService
}

func NewKafkaServerController(logger *logrus.Logger, kafkaServerService *service.KafkaServerService) *KafkaServerController {
	return &KafkaServerController{
		// ctx:	   ctx,
		logger:             logger,
		kafkaServerService: kafkaServerService,
	}
}
