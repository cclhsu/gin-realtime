package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type WebsocketClientInterface interface {
}

type WebsocketClientService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewWebsocketClientService(ctx context.Context, logger *logrus.Logger) *WebsocketClientService {
	return &WebsocketClientService{
		ctx:    ctx,
		logger: logger,
	}
}
