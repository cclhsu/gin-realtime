package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type KafkaServerControllerInterface interface {
	KafkaHandler(ginContext *gin.Context)
}

type KafkaServerController struct {
	ctx				   context.Context
	logger			   *logrus.Logger
	kafkaServerService *service.KafkaServerService
}

func NewKafkaServerController(ctx context.Context, logger *logrus.Logger, kafkaServerService *service.KafkaServerService) *KafkaServerController {
	return &KafkaServerController{
		ctx:				ctx,
		logger:				logger,
		kafkaServerService: kafkaServerService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/kafka/handler' | jq
// @Summary kafka handler
// @Description kafka handler
// @Tags kafka
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /kafka/handler [get]
func (wsc *KafkaServerController) KafkaHandler(ginContext *gin.Context) {
	wsc.logger.Info("KafkaHandler")
	// wsc.kafkaServerService.KafkaHandler(ginContext)
}
