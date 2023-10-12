package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type WebpushServerInterface interface {
}

type WebpushServerService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewWebpushServerService(ctx context.Context, logger *logrus.Logger) *WebpushServerService {
	return &WebpushServerService{
		ctx:    ctx,
		logger: logger,
	}
}
