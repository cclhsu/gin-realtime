package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GraphQLServerControllerInterface interface {
	GraphQLHandler(ginContext *gin.Context)
}

type GraphQLServerController struct {
	ctx					 context.Context
	logger				 *logrus.Logger
	graphQLServerService *service.GraphQLServerService
}

func NewGraphQLServerController(ctx context.Context, logger *logrus.Logger, graphQLServerService *service.GraphQLServerService) *GraphQLServerController {
	return &GraphQLServerController{
		ctx:				  ctx,
		logger:				  logger,
		graphQLServerService: graphQLServerService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/graphql/handler' | jq
// @Summary graphql handler
// @Description graphql handler
// @Tags graphql
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /graphql/handler [get]
func (wsc *GraphQLServerController) GraphQLHandler(ginContext *gin.Context) {
	wsc.logger.Info("GraphQLHandler")
	// wsc.graphQLServerService.GraphQLHandler(ginContext)
}
