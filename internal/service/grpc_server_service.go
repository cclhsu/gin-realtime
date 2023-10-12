package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type GrpcServerInterface interface {
}

type GrpcServerService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewGrpcServerService(ctx context.Context, logger *logrus.Logger) *GrpcServerService {
	return &GrpcServerService{
		ctx:    ctx,
		logger: logger,
	}
}
