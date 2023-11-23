package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type WebRTCServerInterface interface {
}

type WebRTCServerService struct {
	ctx	   context.Context
	logger *logrus.Logger
}

func NewWebRTCServerService(ctx context.Context, logger *logrus.Logger) *WebRTCServerService {
	return &WebRTCServerService{
		ctx:	ctx,
		logger: logger,
	}
}
