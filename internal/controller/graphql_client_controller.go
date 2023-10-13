package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type GraphQLClientControllerInterface interface {
}

type GraphQLClientController struct {
	ctx                  context.Context
	logger               *logrus.Logger
	graphqlClientService *service.GraphQLClientService
}

func NewGraphQLClientController(ctx context.Context, logger *logrus.Logger, graphqlClientService *service.GraphQLClientService) *GraphQLClientController {
	return &GraphQLClientController{
		ctx:                  ctx,
		logger:               logger,
		graphqlClientService: graphqlClientService,
	}
}
