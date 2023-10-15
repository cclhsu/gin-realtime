package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type SocketIOServerInterface interface {
}

type SocketIOServerService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewSocketIOServerService(ctx context.Context, logger *logrus.Logger) *SocketIOServerService {
	return &SocketIOServerService{
		ctx:    ctx,
		logger: logger,
	}
}
