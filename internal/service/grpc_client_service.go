package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type GrpcClientInterface interface {
}

type GrpcClientService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewGrpcClientService(ctx context.Context, logger *logrus.Logger) *GrpcClientService {
	return &GrpcClientService{
		ctx:    ctx,
		logger: logger,
	}
}
