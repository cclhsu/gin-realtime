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
	Send(c *gin.Context)
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

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/webpush-client/health' | jq
// @Summary webpush client health
// @Description webpush client health
// @Tags webpush-client
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webpush-client/health [get]
func (wcc *WebpushClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("WebpushClientController HealthHandler")

	ginContext.JSON(http.StatusOK, gin.H{
		"message": wcc.webpushClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/webpush-client/send?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/webpush-client/send' -d '{"message":"hello"}' | jq
// @Summary webpush client send message
// @Description webpush client send message
// @Tags webpush-client
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webpush-client/send [get]
func (wcc *WebpushClientController) Send(ginContext *gin.Context) {
	wcc.logger.Info("WebpushClientController SendHandler")

	message := ginContext.Query("message")
	message, err := wcc.webpushClientService.Send(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.WebpushMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	//	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	// }

	// message, err	 := wcc.webpushClientService.Send(data)
	// if err != nil {
	//	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	// }

	ginContext.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
