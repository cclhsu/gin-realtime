package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebpushClientControllerInterface interface {
	Connect(c *gin.Context)
	Disconnect(c *gin.Context)
	Trigger(c *gin.Context)
	// Echo(c *gin.Context)
	// Broadcast(c *gin.Context)
	Health(c *gin.Context)
}

type WebpushClientController struct {
	ctx                  context.Context
	logger               *logrus.Logger
	webpushClientService *service.WebpushClientService
}

func NewWebpushClientController(ctx context.Context, logger *logrus.Logger, webpushClientService *service.WebpushClientService) *WebpushClientController {
	return &WebpushClientController{
		ctx:                  ctx,
		logger:               logger,
		webpushClientService: webpushClientService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/webpush/health' | jq
// @Summary webpush client health
// @Description webpush client health
// @Tags webpush
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webpush/health [get]
func (wcc *WebpushClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("WebpushClientController HealthHandler")

	ginContext.JSON(200, gin.H{
		"message": wcc.webpushClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/webpush/trigger?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/webpush/trigger' -d '{"message":"hello"}' | jq
// @Summary webpush client trigger message
// @Description webpush client trigger message
// @Tags webpush
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webpush/trigger [get]
func (wcc *WebpushClientController) Trigger(ginContext *gin.Context) {
	wcc.logger.Info("WebpushClientController TriggerHandler")

	message := ginContext.Query("message")
	message, err := wcc.webpushClientService.Trigger(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.WebpushMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	// 	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// message, err  := wcc.webpushClientService.Trigger(data)
	// if err != nil {
	// 	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	ginContext.JSON(200, gin.H{
		"message": message,
	})
}
