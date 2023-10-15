package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type RedisServerControllerInterface interface {
}

type RedisServerController struct {
	ctx                context.Context
	logger             *logrus.Logger
	redisServerService *service.RedisServerService
}

func NewRedisServerController(ctx context.Context, logger *logrus.Logger, redisServerService *service.RedisServerService) *RedisServerController {
	return &RedisServerController{
		ctx:                ctx,
		logger:             logger,
		redisServerService: redisServerService,
	}
}
