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
	Send(c *gin.Context)
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

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/server-sent-event-client/health' | jq
// @Summary server-sent-event client health
// @Description server-sent-event client health
// @Tags server-sent-event-client
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /server-sent-event-client/health [get]
func (wcc *ServerSentEventClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("ServerSentEventClientController HealthHandler")

	ginContext.JSON(http.StatusOK, gin.H{
		"message": wcc.serverSentEventClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/server-sent-event-client/send?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/server-sent-event-client/send' -d '{"message":"hello"}' | jq
// @Summary server-sent-event client send message
// @Description server-sent-event client send message
// @Tags server-sent-event-client
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /server-sent-event-client/send [get]
func (wcc *ServerSentEventClientController) Send(ginContext *gin.Context) {
	wcc.logger.Info("ServerSentEventClientController SendHandler")

	message := ginContext.Query("message")
	message, err := wcc.serverSentEventClientService.Send(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.ServerSentEventMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	//	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	// }

	// message, err	 := wcc.serverSentEventClientService.Send(data)
	// if err != nil {
	//	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	// }

	ginContext.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
