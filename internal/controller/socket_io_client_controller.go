package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SocketIOClientControllerInterface interface {
	Connect(c *gin.Context)
	Disconnect(c *gin.Context)
	Send(c *gin.Context)
	// Echo(c *gin.Context)
	// Broadcast(c *gin.Context)
	Health(c *gin.Context)
}

type SocketIOClientController struct {
	ctx                   context.Context
	logger                *logrus.Logger
	socketIOClientService *service.SocketIOClientService
}

func NewSocketIOClientController(ctx context.Context, logger *logrus.Logger, socketIOClientService *service.SocketIOClientService) *SocketIOClientController {
	return &SocketIOClientController{
		ctx:                   ctx,
		logger:                logger,
		socketIOClientService: socketIOClientService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/socket-io-client/health' | jq
// @Summary socketIO client health
// @Description socketIO client health
// @Tags socket-io-client
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /socket-io-client/health [get]
func (wcc *SocketIOClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("SocketIOClientController HealthHandler")

	ginContext.JSON(http.StatusOK, gin.H{
		"message": wcc.socketIOClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/socket-io-client/send?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/socket-io-client/send' -d '{"message":"hello"}' | jq
// @Summary socketIO client send message
// @Description socketIO client send message
// @Tags socket-io-client
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /socket-io-client/send [get]
func (wcc *SocketIOClientController) Send(ginContext *gin.Context) {
	wcc.logger.Info("SocketIOClientController SendHandler")

	message := ginContext.Query("message")
	message, err := wcc.socketIOClientService.Send(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.SocketIOMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	//	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	// }

	// message, err	 := wcc.socketIOClientService.Send(data)
	// if err != nil {
	//	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	// }

	ginContext.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
