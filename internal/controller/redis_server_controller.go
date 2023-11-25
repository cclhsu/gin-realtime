package controller

import (
	"context"
	"net/http"

	_interface "github.com/cclhsu/gin-realtime/internal/interface"
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RedisServerControllerInterface interface {
	NewRedisServerController(ctx context.Context, logger *logrus.Logger, redisServerService *service.RedisServerService) *RedisServerController
	_interface.ProducerControllerInterface
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

// curl -s -X 'POST' -H 'accept: application/json' 'http://0.0.0.0:3002/redis/create-topic' -d '{"topic":"test"}' | jq
// @Summary redis server create topic
// @Description redis server create topic
// @Tags redis
// @Accept json
// @Produce json
// @Param topic body string true "topic"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /redis/create-topic [post]
func (wcc *RedisServerController) CreateTopic(c *gin.Context) {
	wcc.logger.Info("RedisServerController CreateTopicHandler")

	var topic model.TopicDTO
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := wcc.redisServerService.CreateTopic(topic.Topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"topic": topic,
	})
}

// curl -s -X 'POST' -H 'accept: application/json' 'http://0.0.0.0:3002/redis/delete-topic' -d '{"topic":"test"}' | jq
// @Summary redis server delete topic
// @Description redis server delete topic
// @Tags redis
// @Accept json
// @Produce json
// @Param topic body string true "topic"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /redis/delete-topic [post]
func (wcc *RedisServerController) DeleteTopic(c *gin.Context) {
	wcc.logger.Info("RedisServerController DeleteTopicHandler")

	var topic model.TopicDTO
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := wcc.redisServerService.DeleteTopic(topic.Topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"topic": topic,
	})
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/redis/produce' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 1, "stage": 0, "action": 0, "environment": 0, "sender": "00000000-0000-0000-0000-000000000000", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/redis/produce' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 2, "stage": 0, "action": 0, "environment": 0, "sender": "00000000-0000-0000-0000-000000000000", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/redis/produce' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 3, "stage": 0, "action": 0, "environment": 0, "sender": "13eca4f1-91ca-4ff9-bdd8-edb9cb63affd", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// @Summary redis server produce message
// @Description redis server produce message
// @Tags redis
// @Accept json
// @Produce json
// @Param topic body string true "topic"
// @Param message body string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /redis/produce [post]
func (wcc *RedisServerController) Produce(c *gin.Context) {
	wcc.logger.Info("RedisServerController ProduceHandler")

	var message model.MessageDTO
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wcc.logger.Infof("RedisServerController ProduceHandler message: %v", message)

	if err := wcc.redisServerService.Produce(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
