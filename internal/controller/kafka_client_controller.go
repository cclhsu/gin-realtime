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
	Send(c *gin.Context)
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

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/kafka-client/health' | jq
// @Summary kafka client health
// @Description kafka client health
// @Tags kafka-client
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /kafka-client/health [get]
func (wcc *KafkaClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("KafkaClientController HealthHandler")

	ginContext.JSON(http.StatusOK, gin.H{
		"message": wcc.kafkaClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/kafka-client/send?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/kafka-client/send' -d '{"message":"hello"}' | jq
// @Summary kafka client send message
// @Description kafka client send message
// @Tags kafka-client
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /kafka-client/send [get]
func (wcc *KafkaClientController) Send(ginContext *gin.Context) {
	wcc.logger.Info("KafkaClientController SendHandler")

	message := ginContext.Query("message")
	message, err := wcc.kafkaClientService.Send(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.KafkaMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	//	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	// }

	// message, err	 := wcc.kafkaClientService.Send(data)
	// if err != nil {
	//	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	// }

	ginContext.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
