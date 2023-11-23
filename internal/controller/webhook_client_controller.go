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

type WebhookClientControllerInterface interface {
	NewWebhookClientController(ctx context.Context, logger *logrus.Logger, webhookClientService *service.WebhookClientService) *WebhookClientController
	_interface.RegistrationControllerInterface
	_interface.MessageControllerInterface
	// _interface.HealthControllerInterface
}

type WebhookClientController struct {
	ctx                  context.Context
	logger               *logrus.Logger
	webhookClientService *service.WebhookClientService
}

func NewWebhookClientController(ctx context.Context, logger *logrus.Logger, webhookClientService *service.WebhookClientService) *WebhookClientController {
	return &WebhookClientController{
		ctx:                  ctx,
		logger:               logger,
		webhookClientService: webhookClientService,
	}
}

// func (wcc *WebhookClientController) Initialize() error {
// 	wcc.logger.Info("WebhookClientController Initialize")
// }

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/webhook-client/register' -d '{}' | jq
// @Summary Register webhook
// @Description Register webhook
// @Tags webhook-client
// @Accept json
// @Produce json
// @Param request body model.RegistrationDTO true "Webhook Info"
// @Success 200 {object} model.RegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook-client/register [post]
func (wcc *WebhookClientController) Register(ginContext *gin.Context) {
	wcc.logger.Info("WebhookClientController Register")

	var request model.RegistrationDTO
	if err := ginContext.ShouldBindJSON(&request); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := wcc.webhookClientService.Register(request)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wcc.logger.Info("WebhookClientController Register: ", response)
	ginContext.JSON(http.StatusOK, response)
}

// curl -s -X 'DELETE' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/webhook-client/register' -d '{}' | jq
// @Summary Unregister webhook
// @Description Unregister webhook
// @Tags webhook-client
// @Accept json
// @Produce json
// @Param webhookId path string true "Webhook ID"
// @Success 200 {object} model.RegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook-client/register/{webhookId} [delete]
func (wcc *WebhookClientController) Unregister(ginContext *gin.Context) {
	wcc.logger.Info("WebhookClientController Unregister")

	webhookID := ginContext.Param("webhookId")

	err := wcc.webhookClientService.Unregister(webhookID)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wcc.logger.Info("WebhookClientController Unregister response: ", webhookID)
	ginContext.JSON(http.StatusOK, gin.H{"webhookId": webhookID})
}

// curl -s -X 'GET' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/webhook-client/register' -d '{}' | jq
// @Summary list webhook registrations
// @Description list webhook registrations
// @Tags webhook-client
// @Accept json
// @Produce json
// @Success 200 {object} []model.RegistrationDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook-client/register [get]
func (wcc *WebhookClientController) ListRegistrations(ginContext *gin.Context) {
	wcc.logger.Info("WebhookClientController ListRegistrations")

	response, err := wcc.webhookClientService.ListRegistrations()
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wcc.logger.Info("WebhookClientController ListRegistrations: ", response)
	ginContext.JSON(http.StatusOK, response)
}

// curl -s -X 'PUT' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/webhook-client/register' -d '{}' | jq
// @Summary Update webhook registration
// @Description Update webhook registration
// @Tags webhook-client
// @Accept json
// @Produce json
// @Param webhookId path string true "Webhook ID"
// @Param request body model.RegistrationDTO true "Webhook Info"
// @Success 200 {object} model.RegistrationResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook-client/register/{webhookId} [put]
func (wcc *WebhookClientController) UpdateRegistration(ginContext *gin.Context) {
	wcc.logger.Info("WebhookClientController UpdateRegistration")

	webhookID := ginContext.Param("webhookId")
	var request model.RegistrationDTO
	if err := ginContext.ShouldBindJSON(&request); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := wcc.webhookClientService.UpdateRegistration(webhookID, request)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wcc.logger.Info("WebhookClientController UpdateRegistration response: ", response)
	ginContext.JSON(http.StatusOK, response)
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/webhook-client/message/send' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 1, "stage": 0, "action": 0, "environment": 0, "sender": "00000000-0000-0000-0000-000000000000", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/webhook-client/message/send' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 2, "stage": 0, "action": 0, "environment": 0, "sender": "00000000-0000-0000-0000-000000000000", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/webhook-client/message/send' -d '{ "UUID": "00000000-0000-0000-0000-000000000000", "type": 3, "stage": 0, "action": 0, "environment": 0, "sender": "13eca4f1-91ca-4ff9-bdd8-edb9cb63affd", "recipient": "00000000-0000-0000-0000-000000000000", "recipientType": 0, "recipients": [], "data": { "additionalProp1": {} }, "metadata": { "additionalProp1": {} } }' | jq
// @Summary Send message
// @Description Send message
// @Tags webhook-client
// @Accept json
// @Produce json
// @Param request body model.MessageDTO true "Message Info"
// @Success 200 {object} model.MessageResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook-client/message/send [post]
func (wcc *WebhookClientController) Send(ginContext *gin.Context) {
	wcc.logger.Info("WebhookClientController Send Message")

	var message model.MessageDTO
	if err := ginContext.ShouldBindJSON(&message); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := wcc.webhookClientService.Send(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wcc.logger.Info("WebhookClientController Send Message: ", message)
	ginContext.JSON(http.StatusOK, gin.H{})
}

// curl -s -X 'POST' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/webhook-client/message/receive' -d '{}' | jq
// @Summary Receive message
// @Description Receive message
// @Tags webhook-client
// @Accept json
// @Produce json
// @Param request body model.MessageDTO true "Message Info"
// @Success 200 {object} model.MessageResponseDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook-client/message/receive [post]
func (wcc *WebhookClientController) Receive(ginContext *gin.Context) {
	wcc.logger.Info("WebhookClientController Receive Message")

	var message model.MessageDTO
	if err := ginContext.ShouldBindJSON(&message); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := wcc.webhookClientService.Receive(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wcc.logger.Info("WebhookClientController Receive Message: ", message)
	ginContext.JSON(http.StatusOK, message)
}

// curl -s -X 'GET' -H 'accept: application/json' -H 'Content-Type: application/json' 'http://0.0.0.0:3002/webhook-client/message' -d '{}' | jq
// @Summary List messages
// @Description List messages
// @Tags webhook-client
// @Accept json
// @Produce json
// @Success 200 {object} []model.MessageDTO
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /webhook-client/message [get]
func (wcc *WebhookClientController) ListMessages(ginContext *gin.Context) {
	wcc.logger.Info("WebhookClientController ListMessages")

	messages, err := wcc.webhookClientService.ListMessages()
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	wcc.logger.Info("WebhookClientController ListMessages: ", messages)
	ginContext.JSON(http.StatusOK, messages)
}
