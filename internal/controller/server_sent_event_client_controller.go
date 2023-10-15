package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ServerSentEventClientControllerInterface interface {
	Connect(c *gin.Context)
	Disconnect(c *gin.Context)
	Trigger(c *gin.Context)
	// Echo(c *gin.Context)
	// Broadcast(c *gin.Context)
	Health(c *gin.Context)
}

type ServerSentEventClientController struct {
	ctx                          context.Context
	logger                       *logrus.Logger
	serverSentEventClientService *service.ServerSentEventClientService
}

func NewServerSentEventClientController(ctx context.Context, logger *logrus.Logger, ServerSentEventClientService *service.ServerSentEventClientService) *ServerSentEventClientController {
	return &ServerSentEventClientController{
		ctx:                          ctx,
		logger:                       logger,
		serverSentEventClientService: ServerSentEventClientService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/serversentevent/health' | jq
// @Summary serversentevent client health
// @Description serversentevent client health
// @Tags serversentevent
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /serversentevent/health [get]
func (wcc *ServerSentEventClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("ServerSentEventClientController HealthHandler")

	ginContext.JSON(200, gin.H{
		"message": wcc.serverSentEventClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/serversentevent/trigger?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/serversentevent/trigger' -d '{"message":"hello"}' | jq
// @Summary serversentevent client trigger message
// @Description serversentevent client trigger message
// @Tags serversentevent
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /serversentevent/trigger [get]
func (wcc *ServerSentEventClientController) Trigger(ginContext *gin.Context) {
	wcc.logger.Info("ServerSentEventClientController TriggerHandler")

	message := ginContext.Query("message")
	message, err := wcc.serverSentEventClientService.Trigger(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.ServerSentEventMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	// 	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// message, err  := wcc.serverSentEventClientService.Trigger(data)
	// if err != nil {
	// 	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	ginContext.JSON(200, gin.H{
		"message": message,
	})
}
