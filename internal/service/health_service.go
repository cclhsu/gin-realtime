package service

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/sirupsen/logrus"
)

type HealthServiceInterface interface {
	IsHealthy(request model.HealthRequest) (model.HealthResponse, error)
	IsALive(request model.HealthRequest) (model.HealthResponse, error)
	IsReady(request model.HealthRequest) (model.HealthResponse, error)
}

type HealthService struct {
	ctx	   context.Context
	logger *logrus.Logger
}

func NewHealthService(ctx context.Context, logger *logrus.Logger) *HealthService {
	return &HealthService{
		ctx:	ctx,
		logger: logger,
	}
}

func (hs *HealthService) IsHealthy(_request model.HealthRequest) (model.HealthResponse, error) {
	response := model.HealthResponse{
		Status: model.ServingStatus_SERVING,
	}
	return response, nil
}

func (hs *HealthService) IsALive(_request model.HealthRequest) (model.HealthResponse, error) {
	response := model.HealthResponse{
		Status: model.ServingStatus_SERVING,
	}
	return response, nil
}

func (hs *HealthService) IsReady(_request model.HealthRequest) (model.HealthResponse, error) {
	response := model.HealthResponse{
		Status: model.ServingStatus_SERVING,
	}
	return response, nil
}
