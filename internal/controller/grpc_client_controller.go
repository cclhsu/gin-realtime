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
	Send(c *gin.Context)
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

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/grpc-client/health' | jq
// @Summary grpc client health
// @Description grpc client health
// @Tags grpc-client
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /grpc-client/health [get]
func (wcc *GrpcClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("GrpcClientController HealthHandler")

	ginContext.JSON(http.StatusOK, gin.H{
		"message": wcc.grpcClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/grpc-client/send?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/grpc-client/send' -d '{"message":"hello"}' | jq
// @Summary grpc client send message
// @Description grpc client send message
// @Tags grpc-client
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /grpc-client/send [get]
func (wcc *GrpcClientController) Send(ginContext *gin.Context) {
	wcc.logger.Info("GrpcClientController SendHandler")

	message := ginContext.Query("message")
	message, err := wcc.grpcClientService.Send(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.GrpcMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	//	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	// }

	// message, err	 := wcc.grpcClientService.Send(data)
	// if err != nil {
	//	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	// }

	ginContext.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
