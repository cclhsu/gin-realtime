package _interface

import (
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/gin-gonic/gin"
)

type HealthControllerInterface interface {
	Health(ginContext *gin.Context) // Get the health/status of the server
}

type HealthServiceInterface interface {
	Health() (model.HealthResponseDTO, error) // Get the health/status of the server
}
