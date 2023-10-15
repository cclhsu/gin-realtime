package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HealthControllerInterface interface {
	IsHealthy(c *gin.Context)
	IsALive(c *gin.Context)
	IsReady(c *gin.Context)
}

type HealthController struct {
	ctx			  context.Context
	logger		  *logrus.Logger
	healthService *service.HealthService
}

func NewHealthController(ctx context.Context, logger *logrus.Logger, healthService *service.HealthService) *HealthController {
	return &HealthController{
		ctx:		   ctx,
		logger:		   logger,
		healthService: healthService,
	}
}

// curl -s -X GET -H 'Content-Type: application/json' http://0.0.0.0:3001/health/healthy | jq
// curl -s -X GET -H 'Content-Type: application/json' http://0.0.0.0:3001/health/healthy?service=cache | jq
// curl -s -X GET -H 'Content-Type: application/json' "http://0.0.0.0:3001/health/healthy" -d "service: cache" | jq
// @Summary Get Health Check
// @Description Get Health Check
// @Tags health
// @Accept json
// @Produce json
// @Param service query string true "Service name" default(service)
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/healthy [get]
func (hcc *HealthController) IsHealthy(c *gin.Context) {
	// Get the 'service' query parameter
	serviceValue := c.Query("service")

	// Create a health request
	request := model.HealthRequest{
		Service: serviceValue,
	}

	// Call the health service to check if the service is healthy
	response, err := hcc.healthService.IsHealthy(request)
	if err != nil {
		// Handle the error, e.g., log it or return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a JSON response with the health status
	c.JSON(http.StatusOK, response)
}

// curl -s -X GET -H 'Content-Type: application/json' http://0.0.0.0:3001/health/live | jq
// curl -s -X GET -H 'Content-Type: application/json' http://0.0.0.0:3001/health/live?service=cache | jq
// curl -s -X GET -H 'Content-Type: application/json' "http://0.0.0.0:3001/health/live" -d "service: cache" | jq
// @Summary Get Health Check
// @Description Get Health Check
// @Tags health
// @Accept json
// @Produce json
// @Param service query string true "Service name" default(service)
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/live [get]
func (hcc *HealthController) IsALive(c *gin.Context) {
	// Get the 'service' query parameter
	serviceValue := c.Query("service")

	// Create a health request
	request := model.HealthRequest{
		Service: serviceValue,
	}

	// Call the health service to check if the service is healthy
	response, err := hcc.healthService.IsALive(request)
	if err != nil {
		// Handle the error, e.g., log it or return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a JSON response with the health status
	c.JSON(http.StatusOK, response)
}

// curl -s -X GET -H 'Content-Type: application/json' http://0.0.0.0:3001/health/ready | jq
// curl -s -X GET -H 'Content-Type: application/json' http://0.0.0.0:3001/health/ready?service=cache | jq
// curl -s -X GET -H 'Content-Type: application/json' "http://0.0.0.0:3001/health/ready" -d "service: cache" | jq
// @Summary Get Health Check
// @Description Get Health Check
// @Tags health
// @Accept json
// @Produce json
// @Param service query string true "Service name" default(service)
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /health/ready [get]
func (hcc *HealthController) IsReady(c *gin.Context) {
	// Get the 'service' query parameter
	serviceValue := c.Query("service")

	// Create a health request
	request := model.HealthRequest{
		Service: serviceValue,
	}

	// Call the health service to check if the service is healthy
	response, err := hcc.healthService.IsReady(request)
	if err != nil {
		// Handle the error, e.g., log it or return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a JSON response with the health status
	c.JSON(http.StatusOK, response)
}
