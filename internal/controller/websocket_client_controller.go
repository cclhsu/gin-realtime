package controller

import (
	"context"
	"fmt"
	"net/http"

	_interface "github.com/cclhsu/gin-realtime/internal/interface"
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebsocketClientControllerInterface interface {
	NewWebsocketClientController(ctx context.Context, logger *logrus.Logger, websocketClientService *service.WebsocketClientService) *WebsocketClientController
	// _interface.ConnectionControllerInterface
	_interface.MessageControllerInterface
	_interface.HealthControllerInterface
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

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/websocket-client/message/send?message=hello' | jq
// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/websocket-client/message/send' -d '{"message":"hello"}' | jq
// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/websocket-client/message/send' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 1, "stage": 0, "action": 0, "environment": 0, "sender": "00000000-0000-0000-0000-000000000000", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/websocket-client/message/send' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 2, "stage": 0, "action": 0, "environment": 0, "sender": "00000000-0000-0000-0000-000000000000", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/websocket-client/message/send' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 3, "stage": 0, "action": 0, "environment": 0, "sender": "13eca4f1-91ca-4ff9-bdd8-edb9cb63affd", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// @Summary websocket client send message
// @Description websocket client send message
// @Tags websocket-client
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /websocket-client/send [post]
func (wcc *WebsocketClientController) Send(ginContext *gin.Context) {
	wcc.logger.Info("WebsocketClientController SendHandler")

	// // Method 1: Send message to server via websocket connection
	// message := ginContext.Param("message")
	// _, msg, err := wcc.websocketClientService.Send(message)
	// if err != nil {
	// 	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// // Method 2: Send message to server via websocket connection
	// message := ginContext.Query("message")
	// _, msg, err := wcc.websocketClientService.Send(message)
	// if err != nil {
	// 	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// Method 3: Send message to server via websocket connection
	var message model.MessageDTO
	if err := ginContext.ShouldBindJSON(&message); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := wcc.websocketClientService.Send(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginContext.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Message sent: %s", message),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/websocket-client/message/receive' | jq
// @Summary websocket client receive message
// @Description websocket client receive message
// @Tags websocket-client
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /websocket-client/message/receive [post]
func (wcc *WebRTCClientController) receive(ginContext *gin.Context) {
	wcc.logger.Info("WebRTCClientController receive")

	// Example data for an unimplemented response
	responseData := map[string]interface{}{
		"error": "Method not implemented",
	}

	// Send JSON response with a status code of 501 (Not Implemented)
	ginContext.JSON(http.StatusNotImplemented, responseData)
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/websocket-client/message' | jq
// @Summary websocket client list messages
// @Description websocket client list messages
// @Tags websocket-client
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /websocket-client/message [get]
func (wcc *WebsocketClientController) ListMessages(ginContext *gin.Context) {
	wcc.logger.Info("WebsocketClientController ListMessages")

	// Example data for an unimplemented response
	responseData := map[string]interface{}{
		"error": "Method not implemented",
	}

	// Send JSON response with a status code of 501 (Not Implemented)
	ginContext.JSON(http.StatusNotImplemented, responseData)
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/websocket-client/health' | jq
// @Summary websocket client health
// @Description websocket client health
// @Tags websocket-client
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /websocket-client/health [get]
func (wcc *WebsocketClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("WebsocketClientController HealthHandler")

	ginContext.JSON(http.StatusOK, gin.H{
		"message": wcc.websocketClientService.Health(),
	})
}
