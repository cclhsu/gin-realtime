package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RedisServerControllerInterface interface {
	RedisHandler(ginContext *gin.Context)
}

type RedisServerController struct {
	ctx				   context.Context
	logger			   *logrus.Logger
	redisServerService *service.RedisServerService
}

func NewRedisServerController(ctx context.Context, logger *logrus.Logger, redisServerService *service.RedisServerService) *RedisServerController {
	return &RedisServerController{
		ctx:				ctx,
		logger:				logger,
		redisServerService: redisServerService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/redis/handler' | jq
// @Summary redis handler
// @Description redis handler
// @Tags redis
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /redis/handler [get]
func (wsc *RedisServerController) RedisHandler(ginContext *gin.Context) {
	wsc.logger.Info("RedisHandler")
	// wsc.redisServerService.RedisHandler(ginContext)
}
