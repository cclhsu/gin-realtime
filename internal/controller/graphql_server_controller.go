package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/sirupsen/logrus"
)

type GraphQLServerControllerInterface interface {
}

type GraphQLServerController struct {
	ctx					 context.Context
	logger				 *logrus.Logger
	graphqlServerService *service.GraphQLServerService
}

func NewGraphQLServerController(ctx context.Context, logger *logrus.Logger, graphqlServerService *service.GraphQLServerService) *GraphQLServerController {
	return &GraphQLServerController{
		ctx:				  ctx,
		logger:				  logger,
		graphqlServerService: graphqlServerService,
	}
}
