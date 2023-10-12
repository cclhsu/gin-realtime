package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type ServerSentEventServerInterface interface {
}

type ServerSentEventServerService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewServerSentEventServerService(ctx context.Context, logger *logrus.Logger) *ServerSentEventServerService {
	return &ServerSentEventServerService{
		ctx:    ctx,
		logger: logger,
	}
}
