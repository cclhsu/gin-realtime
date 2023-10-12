package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type WebsocketServerInterface interface {
}

type WebsocketServerService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewWebsocketServerService(ctx context.Context, logger *logrus.Logger) *WebsocketServerService {
	return &WebsocketServerService{
		ctx:    ctx,
		logger: logger,
	}
}
