package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type GraphQLClientInterface interface {
}

type GraphQLClientService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewGraphQLClientService(ctx context.Context, logger *logrus.Logger) *GraphQLClientService {
	return &GraphQLClientService{
		ctx:    ctx,
		logger: logger,
	}
}
