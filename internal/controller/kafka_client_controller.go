package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type KafkaClientControllerInterface interface {
	Connect(c *gin.Context)
	Disconnect(c *gin.Context)
	Trigger(c *gin.Context)
	// Echo(c *gin.Context)
	// Broadcast(c *gin.Context)
	Health(c *gin.Context)
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

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/kafka/health' | jq
// @Summary kafka client health
// @Description kafka client health
// @Tags kafka
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /kafka/health [get]
func (wcc *KafkaClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("KafkaClientController HealthHandler")

	ginContext.JSON(200, gin.H{
		"message": wcc.kafkaClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/kafka/trigger?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/kafka/trigger' -d '{"message":"hello"}' | jq
// @Summary kafka client trigger message
// @Description kafka client trigger message
// @Tags kafka
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /kafka/trigger [get]
func (wcc *KafkaClientController) Trigger(ginContext *gin.Context) {
	wcc.logger.Info("KafkaClientController TriggerHandler")

	message := ginContext.Query("message")
	message, err := wcc.kafkaClientService.Trigger(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.KafkaMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	// 	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// message, err  := wcc.kafkaClientService.Trigger(data)
	// if err != nil {
	// 	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	ginContext.JSON(200, gin.H{
		"message": message,
	})
}
