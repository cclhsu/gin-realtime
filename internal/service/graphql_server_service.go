package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

type GraphQLServerInterface interface {
}

type GraphQLServerService struct {
	ctx	   context.Context
	logger *logrus.Logger
}

func NewGraphQLServerService(ctx context.Context, logger *logrus.Logger) *GraphQLServerService {
	return &GraphQLServerService{
		ctx:	ctx,
		logger: logger,
	}
}
