package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type WebpushClientInterface interface {
}

type WebpushClientService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewWebpushClientService(ctx context.Context, logger *logrus.Logger) *WebpushClientService {
	return &WebpushClientService{
		ctx:    ctx,
		logger: logger,
	}
}
