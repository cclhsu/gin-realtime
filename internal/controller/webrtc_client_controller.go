package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebRTCClientControllerInterface interface {
	Connect(c *gin.Context)
	Disconnect(c *gin.Context)
	Send(c *gin.Context)
	// Echo(c *gin.Context)
	// Broadcast(c *gin.Context)
	Health(c *gin.Context)
}

type WebRTCClientController struct {
	ctx                 context.Context
	logger              *logrus.Logger
	webrtcClientService *service.WebRTCClientService
}

func NewWebRTCClientController(ctx context.Context, logger *logrus.Logger, webrtcClientService *service.WebRTCClientService) *WebRTCClientController {
	return &WebRTCClientController{
		ctx:                 ctx,
		logger:              logger,
		webrtcClientService: webrtcClientService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/webrtc-client/health' | jq
// @Summary webrtc client health
// @Description webrtc client health
// @Tags webrtc-client
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webrtc-client/health [get]
func (wcc *WebRTCClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("WebRTCClientController HealthHandler")

	ginContext.JSON(http.StatusOK, gin.H{
		"message": wcc.webrtcClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/webrtc-client/send?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/webrtc-client/send' -d '{"message":"hello"}' | jq
// @Summary webrtc client send message
// @Description webrtc client send message
// @Tags webrtc-client
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webrtc-client/send [get]
func (wcc *WebRTCClientController) Send(ginContext *gin.Context) {
	wcc.logger.Info("WebRTCClientController SendHandler")

	message := ginContext.Query("message")
	message, err := wcc.webrtcClientService.Send(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.WebRTCMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	//	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	// }

	// message, err	 := wcc.webrtcClientService.Send(data)
	// if err != nil {
	//	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	// }

	ginContext.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
