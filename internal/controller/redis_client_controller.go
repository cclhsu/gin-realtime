package controller

import (
	"context"

	_interface "github.com/cclhsu/gin-realtime/internal/interface"
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type RedisClientControllerInterface interface {
	NewRedisClientController(ctx context.Context, logger *logrus.Logger, redisClientService *service.RedisClientService) *RedisClientController
	_interface.ConsumerControllerInterface
}

type RedisClientController struct {
	ctx                context.Context
	logger             *logrus.Logger
	redisClientService *service.RedisClientService
}

func NewRedisClientController(ctx context.Context, logger *logrus.Logger, redisClientService *service.RedisClientService) *RedisClientController {
	return &RedisClientController{
		ctx:                ctx,
		logger:             logger,
		redisClientService: redisClientService,
	}
}
