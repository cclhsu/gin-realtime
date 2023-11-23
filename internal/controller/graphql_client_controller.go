package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GraphQLClientControllerInterface interface {
	Connect(c *gin.Context)
	Disconnect(c *gin.Context)
	Send(c *gin.Context)
	// Echo(c *gin.Context)
	// Broadcast(c *gin.Context)
	Health(c *gin.Context)
}

type GraphQLClientController struct {
	ctx                  context.Context
	logger               *logrus.Logger
	graphQLClientService *service.GraphQLClientService
}

func NewGraphQLClientController(ctx context.Context, logger *logrus.Logger, graphQLClientService *service.GraphQLClientService) *GraphQLClientController {
	return &GraphQLClientController{
		ctx:                  ctx,
		logger:               logger,
		graphQLClientService: graphQLClientService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/graphql-client/health' | jq
// @Summary graphql client health
// @Description graphql client health
// @Tags graphql-client
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /graphql-client/health [get]
func (wcc *GraphQLClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("GraphQLClientController HealthHandler")

	ginContext.JSON(http.StatusOK, gin.H{
		"message": wcc.graphQLClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/graphql-client/send?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/graphql-client/send' -d '{"message":"hello"}' | jq
// @Summary graphql client send message
// @Description graphql client send message
// @Tags graphql-client
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /graphql-client/send [get]
func (wcc *GraphQLClientController) Send(ginContext *gin.Context) {
	wcc.logger.Info("GraphQLClientController SendHandler")

	message := ginContext.Query("message")
	message, err := wcc.graphQLClientService.Send(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.GraphQLMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	//	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	// }

	// message, err	 := wcc.graphQLClientService.Send(data)
	// if err != nil {
	//	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	// }

	ginContext.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
