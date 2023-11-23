package _interface

import (
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/gin-gonic/gin"
)

type ConnectionControllerInterface interface {
	Connect(ginContext *gin.Context)         // Connect to the server
	Disconnect(ginContext *gin.Context)      // Disconnect from the server
	ListConnections(ginContext *gin.Context) // List all connections
}

type ConnectionServiceInterface interface {
	Connect(connection model.ConnectionDTO) (model.ConnectionResponseDTO, error) // Connect to the server
	Disconnect(connectionID string) error                                        // Disconnect from the server
	ListConnections() ([]model.ConnectionDTO, error)                             // List all connections
}
