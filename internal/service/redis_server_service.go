package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type RedisServerInterface interface {
}

type RedisServerService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewRedisServerService(ctx context.Context, logger *logrus.Logger) *RedisServerService {
	return &RedisServerService{
		ctx:    ctx,
		logger: logger,
	}
}
