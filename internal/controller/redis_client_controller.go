package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RedisClientControllerInterface interface {
	Connect(c *gin.Context)
	Disconnect(c *gin.Context)
	Trigger(c *gin.Context)
	// Echo(c *gin.Context)
	// Broadcast(c *gin.Context)
	Health(c *gin.Context)
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

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/redis/health' | jq
// @Summary redis client health
// @Description redis client health
// @Tags redis
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /redis/health [get]
func (wcc *RedisClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("RedisClientController HealthHandler")

	ginContext.JSON(200, gin.H{
		"message": wcc.redisClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/redis/trigger?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/redis/trigger' -d '{"message":"hello"}' | jq
// @Summary redis client trigger message
// @Description redis client trigger message
// @Tags redis
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /redis/trigger [get]
func (wcc *RedisClientController) Trigger(ginContext *gin.Context) {
	wcc.logger.Info("RedisClientController TriggerHandler")

	message := ginContext.Query("message")
	message, err := wcc.redisClientService.Trigger(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.RedisMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	// 	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// message, err  := wcc.redisClientService.Trigger(data)
	// if err != nil {
	// 	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	ginContext.JSON(200, gin.H{
		"message": message,
	})
}
