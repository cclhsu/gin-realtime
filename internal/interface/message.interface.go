package _interface

import (
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/gin-gonic/gin"
)

type MessageControllerInterface interface {
	Send(ginContext *gin.Context)         // Send a message to the server
	Receive(ginContext *gin.Context)      // Receive a message from the server
	ListMessages(ginContext *gin.Context) // List messages
}

type MessageServiceInterface interface {
	Send(message model.MessageDTO) error       // Send a message to the server
	Receive(message model.MessageDTO) error    // Receive a message from the server
	ListMessages() ([]model.MessageDTO, error) // List messages
}
