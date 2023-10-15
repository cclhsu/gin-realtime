package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GrpcClientControllerInterface interface {
	Connect(c *gin.Context)
	Disconnect(c *gin.Context)
	Trigger(c *gin.Context)
	// Echo(c *gin.Context)
	// Broadcast(c *gin.Context)
	Health(c *gin.Context)
}

type GrpcClientController struct {
	ctx               context.Context
	logger            *logrus.Logger
	grpcClientService *service.GrpcClientService
}

func NewGrpcClientController(ctx context.Context, logger *logrus.Logger, grpcClientService *service.GrpcClientService) *GrpcClientController {
	return &GrpcClientController{
		ctx:               ctx,
		logger:            logger,
		grpcClientService: grpcClientService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/grpc/health' | jq
// @Summary grpc client health
// @Description grpc client health
// @Tags grpc
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /grpc/health [get]
func (wcc *GrpcClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("GrpcClientController HealthHandler")

	ginContext.JSON(200, gin.H{
		"message": wcc.grpcClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/grpc/trigger?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/grpc/trigger' -d '{"message":"hello"}' | jq
// @Summary grpc client trigger message
// @Description grpc client trigger message
// @Tags grpc
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /grpc/trigger [get]
func (wcc *GrpcClientController) Trigger(ginContext *gin.Context) {
	wcc.logger.Info("GrpcClientController TriggerHandler")

	message := ginContext.Query("message")
	message, err := wcc.grpcClientService.Trigger(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.GrpcMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	// 	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// message, err  := wcc.grpcClientService.Trigger(data)
	// if err != nil {
	// 	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	ginContext.JSON(200, gin.H{
		"message": message,
	})
}
