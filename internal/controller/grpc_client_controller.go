package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type GrpcClientControllerInterface interface {
}

type GrpcClientController struct {
	ctx               context.Context
	logger            *logrus.Logger
	grpcClientService *service.GrpcClientService
}

func NewGrpcClientController(ctx context.Context, logger *logrus.Logger, grpcClientService *service.GrpcClientService) *GrpcClientController {
	return &GrpcClientController{
		ctx:               ctx,
		logger:            logger,
		grpcClientService: grpcClientService,
	}
}
