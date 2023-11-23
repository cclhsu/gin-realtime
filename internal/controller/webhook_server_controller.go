package controller

import (
	"context"
	"net/http"

	_interface "github.com/cclhsu/gin-realtime/internal/interface"
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type WebhookServerControllerInterface interface {
	NewWebhookServerController(ctx context.Context, logger *logrus.Logger, webhookServerService *service.WebhookServerService) *WebhookServerController
	_interface.RegistrationControllerInterface
	_interface.MessageControllerInterface
	// _interface.HealthControllerInterface
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

// func (wsc *WebhookServerController) Initialize() error {
// 	wsc.logger.Info("WebhookServerController Initialize")
// }

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/webhook/' -d '{}' | jq
// @Summary Register webhook
// @Description Register webhook
// @Tags webhook
// @Accept json
// @Produce json
// @Param request body model.RegistrationDTO true "Webhook Info"
// @Success 200 {object} model.RegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/register [post]
func (wsc *WebhookServerController) Register(ginContext *gin.Context) {
	wsc.logger.Info("WebhookServerController Register")

	var request model.RegistrationDTO
	if err := ginContext.ShouldBindJSON(&request); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := wsc.webhookServerService.Register(request)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wsc.logger.Info("WebhookServerController Register: ", response)
	ginContext.JSON(http.StatusOK, response)
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/webhook/' -d '{}' | jq
// @Summary Unregister webhook
// @Description Unregister webhook
// @Tags webhook
// @Accept json
// @Produce json
// @Param webhookId path string true "Webhook ID"
// @Success 200 {object} model.RegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/register/{webhookId} [delete]
func (wsc *WebhookServerController) Unregister(ginContext *gin.Context) {
	wsc.logger.Info("WebhookServerController Unregister")

	webhookID := ginContext.Param("webhookId")

	err := wsc.webhookServerService.Unregister(webhookID)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wsc.logger.Info("WebhookServerController Unregister", webhookID)
	ginContext.JSON(http.StatusOK, gin.H{})
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/webhook/' -d '{}' | jq
// @Summary List registered webhooks
// @Description List registered webhooks
// @Tags webhook
// @Accept json
// @Produce json
// @Success 200 {object} []model.RegistrationDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/register [get]
func (wsc *WebhookServerController) ListRegistrations(ginContext *gin.Context) {
	wsc.logger.Info("WebhookServerController ListRegistrations")

	response, err := wsc.webhookServerService.ListRegistrations()
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wsc.logger.Info("WebhookServerController ListRegistrations: ", response)
	ginContext.JSON(http.StatusOK, response)
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/webhook/' -d '{}' | jq
// @Summary Update webhook registration
// @Description Update webhook registration
// @Tags webhook
// @Accept json
// @Produce json
// @Param webhookId path string true "Webhook ID"
// @Param request body model.RegistrationDTO true "Webhook Info"
// @Success 200 {object} model.RegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/register/{webhookId} [put]
func (wsc *WebhookServerController) UpdateRegistration(ginContext *gin.Context) {
	wsc.logger.Info("WebhookServerController UpdateRegistration")

	webhookID := ginContext.Param("webhookId")
	var request model.RegistrationDTO
	if err := ginContext.ShouldBindJSON(&request); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := wsc.webhookServerService.UpdateRegistration(webhookID, request)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wsc.logger.Info("WebhookServerController UpdateRegistration response: ", response)
	ginContext.JSON(http.StatusOK, response)
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/webhook/' -d '{}' | jq
// @Summary Send message
// @Description Send message
// @Tags webhook
// @Accept json
// @Produce json
// @Param request body model.MessageDTO true "Message Info"
// @Success 200 {object} model.MessageResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/message/send [post]
func (wsc *WebhookServerController) Send(ginContext *gin.Context) {
	wsc.logger.Info("WebhookServerController Send Message")

	var message model.MessageDTO
	if err := ginContext.ShouldBindJSON(&message); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := wsc.webhookServerService.Send(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wsc.logger.Info("WebhookServerController Send Message: ", message)
	ginContext.JSON(http.StatusOK, gin.H{})
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/webhook/' -d '{}' | jq
// @Summary Receive message
// @Description Receive message
// @Tags webhook
// @Accept json
// @Produce json
// @Param request body model.MessageDTO true "Message Info"
// @Success 200 {object} model.MessageResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/message/receive [post]
func (wsc *WebhookServerController) Receive(ginContext *gin.Context) {
	wsc.logger.Info("WebhookServerController Receive Message")

	var message model.MessageDTO
	if err := ginContext.ShouldBindJSON(&message); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := wsc.webhookServerService.Receive(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wsc.logger.Info("WebhookServerController Receive Message: ", message)
	ginContext.JSON(http.StatusOK, message)
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3001/webhook/' -d '{}' | jq
// @Summary List messages
// @Description List messages
// @Tags webhook
// @Accept json
// @Produce json
// @Success 200 {object} []model.MessageDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook/message [get]
func (wsc *WebhookServerController) ListMessages(ginContext *gin.Context) {
	wsc.logger.Info("WebhookServerController ListMessages")

	messages, err := wsc.webhookServerService.ListMessages()
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wsc.logger.Info("WebhookServerController ListMessages: ", messages)
	ginContext.JSON(http.StatusOK, messages)
}
