package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebsocketClientControllerInterface interface {
	Connect(c *gin.Context)
	Disconnect(c *gin.Context)
	Trigger(c *gin.Context)
	// Echo(c *gin.Context)
	// Broadcast(c *gin.Context)
	Health(c *gin.Context)
}

type WebsocketClientController struct {
	ctx                    context.Context
	logger                 *logrus.Logger
	websocketClientService *service.WebsocketClientService
}

func NewWebsocketClientController(ctx context.Context, logger *logrus.Logger, websocketClientService *service.WebsocketClientService) *WebsocketClientController {
	return &WebsocketClientController{
		ctx:                    ctx,
		logger:                 logger,
		websocketClientService: websocketClientService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/websocket/health' | jq
// @Summary websocket client health
// @Description websocket client health
// @Tags websocket
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /websocket/health [get]
func (wcc *WebsocketClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("WebsocketClientController HealthHandler")

	ginContext.JSON(200, gin.H{
		"message": wcc.websocketClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/websocket/trigger?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/websocket/trigger' -d '{"message":"hello"}' | jq
// @Summary websocket client trigger message
// @Description websocket client trigger message
// @Tags websocket
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /websocket/trigger [get]
func (wcc *WebsocketClientController) Trigger(ginContext *gin.Context) {
	wcc.logger.Info("WebsocketClientController TriggerHandler")

	message := ginContext.Query("message")
	_, message, err := wcc.websocketClientService.Trigger(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.WebsocketMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	// 	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// _, message, err  := wcc.websocketClientService.Trigger(data)
	// if err != nil {
	// 	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	ginContext.JSON(200, gin.H{
		"message": message,
	})
}
