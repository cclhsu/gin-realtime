package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type ServerSentEventClientInterface interface {
}

type ServerSentEventClientService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewServerSentEventClientService(ctx context.Context, logger *logrus.Logger) *ServerSentEventClientService {
	return &ServerSentEventClientService{
		ctx:    ctx,
		logger: logger,
	}
}
