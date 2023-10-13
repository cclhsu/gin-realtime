package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebhookServerControllerInterface interface {
	RegisterWebhook(c *gin.Context)
	HandleWebhookEvent(c *gin.Context)
	UnregisterWebhook(c *gin.Context)
	ListRegisteredWebhooks(c *gin.Context)
	UpdateWebhook(c *gin.Context)
}

type WebhookServerController struct {
	ctx                  context.Context
	logger               *logrus.Logger
	webhookServerService *service.WebhookServerService
}

func NewWebhookServerController(ctx context.Context, logger *logrus.Logger, webhookServerService *service.WebhookServerService) *WebhookServerController {
	return &WebhookServerController{
		ctx:                  ctx,
		logger:               logger,
		webhookServerService: webhookServerService,
	}
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"id":"1234567890","url":"http://localhost:3002/webhook/1234567890","expiryDate":"2021-08-31T00:00:00Z","config":{"secret":"1234567890","isActive":true,"type":"test"}}' 'http://0.0.0.0:3001/webhook/register' | jq
// @Summary register webhook
// @Description register webhook
// @Tags webhook
// @Accept json
// @Produce json
// @Param request body model.WebhookInfoDTO true "Webhook Info"
// @Success 200 {object} model.WebhookRegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/register [post]
func (wsc *WebhookServerController) RegisterWebhook(c *gin.Context) {
	var request model.WebhookInfoDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := wsc.webhookServerService.RegisterWebhook(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/webhook/handle-event' -d '{"data": "hello world","id": "0123456789","type": "test"}'
// @Summary list registered webhooks
// @Description list registered webhooks
// @Tags webhook
// @Accept json
// @Produce json
// @Success 200 {object} []model.WebhookInfoDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/handle-event [post]
func (wsc *WebhookServerController) HandleWebhookEvent(c *gin.Context) {
	var payload model.EventDataDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := wsc.webhookServerService.HandleWebhookEvent(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// curl -s -X 'DELETE' -H 'accept: application/json' 'http://0.0.0.0:3001/webhook/1234567890' | jq
// @Summary unregister webhook
// @Description unregister webhook
// @Tags webhook
// @Accept json
// @Produce json
// @Param webhookId path string true "Webhook ID"
// @Success 200 {object} model.WebhookRegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/{webhookId} [delete]
func (wsc *WebhookServerController) UnregisterWebhook(c *gin.Context) {
	webhookID := c.Param("webhookId")

	response, err := wsc.webhookServerService.UnregisterWebhook(webhookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// curl -s -X 'GET' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/webhook/list' | jq
// @Summary update webhook
// @Description update webhook
// @Tags webhook
// @Accept json
// @Produce json
// @Param webhookId path string true "Webhook ID"
// @Param request body model.WebhookInfoDTO true "Webhook Info"
// @Success 200 {object} model.WebhookRegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/list [get]
func (wsc *WebhookServerController) ListRegisteredWebhooks(c *gin.Context) {
	webhooks, err := wsc.webhookServerService.ListRegisteredWebhooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, webhooks)
}

// curl -s -X 'PUT' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"id":"1234567890","url":"http://localhost:3002/webhook/1234567890","expiryDate":"2021-08-31T00:00:00Z","config":{"secret":"1234567890","isActive":true,"type":"test"}}' 'http://0.0.0.0:3001/webhook/1234567890' | jq
// @Summary update webhook
// @Description update webhook
// @Tags webhook
// @Accept json
// @Produce json
// @Param webhookId path string true "Webhook ID"
// @Param request body model.WebhookInfoDTO true "Webhook Info"
// @Success 200 {object} model.WebhookRegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/{webhookId} [put]
func (wsc *WebhookServerController) UpdateWebhook(c *gin.Context) {
	webhookID := c.Param("webhookId")
	var request model.WebhookInfoDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := wsc.webhookServerService.UpdateWebhook(webhookID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
