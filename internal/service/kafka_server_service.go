package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type KafkaServerInterface interface {
}

type KafkaServerService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewKafkaServerService(ctx context.Context, logger *logrus.Logger) *KafkaServerService {
	return &KafkaServerService{
		ctx:    ctx,
		logger: logger,
	}
}
