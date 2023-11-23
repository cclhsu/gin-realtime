package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type ElasticsearchServerInterface interface {
}

type ElasticsearchServerService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewElasticsearchServerService(ctx context.Context, logger *logrus.Logger) *ElasticsearchServerService {
	return &ElasticsearchServerService{
		ctx:    ctx,
		logger: logger,
	}
}
