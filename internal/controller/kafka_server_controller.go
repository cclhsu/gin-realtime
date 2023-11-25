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

type KafkaServerControllerInterface interface {
	NewKafkaServerController(ctx context.Context, logger *logrus.Logger, kafkaServerService *service.KafkaServerService) *KafkaServerController
	_interface.ProducerControllerInterface
}

type KafkaServerController struct {
	ctx                context.Context
	logger             *logrus.Logger
	kafkaServerService *service.KafkaServerService
}

func NewKafkaServerController(ctx context.Context, logger *logrus.Logger, kafkaServerService *service.KafkaServerService) *KafkaServerController {
	return &KafkaServerController{
		ctx:                ctx,
		logger:             logger,
		kafkaServerService: kafkaServerService,
	}
}

// curl -s -X 'POST' -H 'accept: application/json' 'http://0.0.0.0:3002/kafka/create-topic' -d '{"topic":"test"}' | jq
// @Summary kafka server create topic
// @Description kafka server create topic
// @Tags kafka
// @Accept json
// @Produce json
// @Param topic body string true "topic"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /kafka/create-topic [post]
func (wcc *KafkaServerController) CreateTopic(c *gin.Context) {
	wcc.logger.Info("KafkaServerController CreateTopicHandler")

	var topic model.TopicDTO
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := wcc.kafkaServerService.CreateTopic(topic.Topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"topic": topic,
	})
}

// curl -s -X 'POST' -H 'accept: application/json' 'http://0.0.0.0:3002/kafka/delete-topic' -d '{"topic":"test"}' | jq
// @Summary kafka server delete topic
// @Description kafka server delete topic
// @Tags kafka
// @Accept json
// @Produce json
// @Param topic body string true "topic"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /kafka/delete-topic [post]
func (wcc *KafkaServerController) DeleteTopic(c *gin.Context) {
	wcc.logger.Info("KafkaServerController DeleteTopicHandler")

	var topic model.TopicDTO
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := wcc.kafkaServerService.DeleteTopic(topic.Topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"topic": topic,
	})
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/kafka/produce' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 1, "stage": 0, "action": 0, "environment": 0, "sender": "00000000-0000-0000-0000-000000000000", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/kafka/produce' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 2, "stage": 0, "action": 0, "environment": 0, "sender": "00000000-0000-0000-0000-000000000000", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/kafka/produce' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 3, "stage": 0, "action": 0, "environment": 0, "sender": "13eca4f1-91ca-4ff9-bdd8-edb9cb63affd", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// @Summary kafka server produce message
// @Description kafka server produce message
// @Tags kafka
// @Accept json
// @Produce json
// @Param topic body string true "topic"
// @Param message body string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /kafka/produce [post]
func (wcc *KafkaServerController) Produce(c *gin.Context) {
	wcc.logger.Info("KafkaServerController ProduceHandler")

	var message model.MessageDTO
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wcc.logger.Infof("KafkaServerController ProduceHandler message: %v", message)

	if err := wcc.kafkaServerService.Produce(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
