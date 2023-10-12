package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type KafkaClientInterface interface {
}

type KafkaClientService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewKafkaClientService(ctx context.Context, logger *logrus.Logger) *KafkaClientService {
	return &KafkaClientService{
		ctx:    ctx,
		logger: logger,
	}
}
