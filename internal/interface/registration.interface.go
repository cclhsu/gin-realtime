package _interface

import (
	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/gin-gonic/gin"
)

type RegistrationControllerInterface interface {
	Register(ginContext *gin.Context)           // Register with the server
	Unregister(ginContext *gin.Context)         // Unregister from the server
	ListRegistrations(ginContext *gin.Context)  // List registrations with the server
	UpdateRegistration(ginContext *gin.Context) // Update registration with the server
}

type RegistrationServiceInterface interface {
	Register(registration model.RegistrationDTO) (model.RegistrationResponseDTO, error)                                  // Register with the server
	Unregister(registrationID string) error                                                                              // Unregister from the server
	ListRegistrations() ([]model.RegistrationDTO, error)                                                                 // List registrations with the server
	UpdateRegistration(registrationID string, registration model.RegistrationDTO) (model.RegistrationResponseDTO, error) // Update registration with the server
}
