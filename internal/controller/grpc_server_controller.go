package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type GrpcServerControllerInterface interface {
}

type GrpcServerController struct {
	ctx               context.Context
	logger            *logrus.Logger
	grpcServerService *service.GrpcServerService
}

func NewGrpcServerController(ctx context.Context, logger *logrus.Logger, grpcServerService *service.GrpcServerService) *GrpcServerController {
	return &GrpcServerController{
		ctx:               ctx,
		logger:            logger,
		grpcServerService: grpcServerService,
	}
}
