package controller

import (
	"context"

	_interface "github.com/cclhsu/gin-realtime/internal/interface"
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type KafkaClientControllerInterface interface {
	NewKafkaClientController(ctx context.Context, logger *logrus.Logger, kafkaClientService *service.KafkaClientService) *KafkaClientController
	_interface.ConsumerControllerInterface
}

type KafkaClientController struct {
	ctx                context.Context
	logger             *logrus.Logger
	kafkaClientService *service.KafkaClientService
}

func NewKafkaClientController(ctx context.Context, logger *logrus.Logger, kafkaClientService *service.KafkaClientService) *KafkaClientController {
	return &KafkaClientController{
		ctx:                ctx,
		logger:             logger,
		kafkaClientService: kafkaClientService,
	}
}
