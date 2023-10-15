package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebhookClientControllerInterface interface {
	RegisterWebhook(c *gin.Context)
	TriggerEvent(c *gin.Context)
	HandleTriggeredEvent(c *gin.Context)
	DisconnectWebhook(c *gin.Context)
}

type WebhookClientController struct {
	ctx					 context.Context
	logger				 *logrus.Logger
	webhookClientService *service.WebhookClientService
}

func NewWebhookClientController(ctx context.Context, logger *logrus.Logger, webhookClientService *service.WebhookClientService) *WebhookClientController {
	return &WebhookClientController{
		ctx:				  ctx,
		logger:				  logger,
		webhookClientService: webhookClientService,
	}
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"id":"1234567890","url":"http://localhost:3002/webhook-client/handle-payload","expiryDate":"2021-08-31T00:00:00Z","config":{"secret":"1234567890","isActive":true,"type":"test"}}' 'http://0.0.0.0:3002/webhook-client/register' | jq
// @Summary register webhook
// @Description register webhook
// @Tags webhook-client
// @Accept json
// @Produce json
// @Param request body model.WebhookInfoDTO true "Webhook Info"
// @Success 200 {object} model.WebhookRegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook-client/register [post]
func (wcc *WebhookClientController) RegisterWebhook(c *gin.Context) {
	var webhookData model.WebhookInfoDTO
	if err := c.ShouldBindJSON(&webhookData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := wcc.webhookClientService.RegisterWebhook(webhookData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"data": "hello world", "id": "1234567890", "type": "test"}' 'http://0.0.0.0:3002/webhook-client/trigger-event' | jq
// @Summary trigger event
// @Description trigger event
// @Tags webhook-client
// @Accept json
// @Produce json
// @Param request body model.EventDataDTO true "Event Data"
// @Success 200 {object} model.EventDataResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook-client/trigger-event [post]
func (wcc *WebhookClientController) TriggerEvent(c *gin.Context) {
	var eventDataDTO model.EventDataDTO
	if err := c.ShouldBindJSON(&eventDataDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := wcc.webhookClientService.TriggerEvent(eventDataDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"id":"1234567890","url":"http://localhost:3002/webhook-client/1234567890","expiryDate":"2021-08-31T00:00:00Z","config":{"secret":"1234567890","isActive":true,"type":"test"}}' 'http://0.0.0.0:3002/webhook-client/handle-payload' | jq
// @Summary handle triggered event
// @Description handle triggered event
// @Tags webhook-client
// @Accept json
// @Produce json
// @Param request body model.EventDataDTO true "Event Data"
// @Success 200 {object} model.EventDataResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook-client/handle-payload [post]
func (wcc *WebhookClientController) HandleTriggeredEvent(c *gin.Context) {
	var eventDataDTO model.EventDataDTO
	if err := c.ShouldBindJSON(&eventDataDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := wcc.webhookClientService.HandleTriggeredEvent(eventDataDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{}' 'http://0.0.0.0:3002/webhook-client/disconnect/1234567890' | jq
// @Summary disconnect webhook
// @Description disconnect webhook
// @Tags webhook-client
// @Accept json
// @Produce json
// @Param webhookId path string true "Webhook ID"
// @Success 200 {object} model.WebhookRegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook-client/disconnect/{webhookId} [put]
func (wcc *WebhookClientController) DisconnectWebhook(c *gin.Context) {
	webhookID := c.Param("webhookId")

	response, err := wcc.webhookClientService.DisconnectWebhook(webhookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
