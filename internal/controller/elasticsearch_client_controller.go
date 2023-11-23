package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ElasticsearchClientControllerInterface interface {
	Connect(c *gin.Context)
	Disconnect(c *gin.Context)
	Send(c *gin.Context)
	// Echo(c *gin.Context)
	// Broadcast(c *gin.Context)
	Health(c *gin.Context)
}

type ElasticsearchClientController struct {
	ctx                        context.Context
	logger                     *logrus.Logger
	elasticsearchClientService *service.ElasticsearchClientService
}

func NewElasticsearchClientController(ctx context.Context, logger *logrus.Logger, elasticsearchClientService *service.ElasticsearchClientService) *ElasticsearchClientController {
	return &ElasticsearchClientController{
		ctx:                        ctx,
		logger:                     logger,
		elasticsearchClientService: elasticsearchClientService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/elasticsearch-client/health' | jq
// @Summary elasticsearch client health
// @Description elasticsearch client health
// @Tags elasticsearch-client
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /elasticsearch-client/health [get]
func (wcc *ElasticsearchClientController) Health(ginContext *gin.Context) {
	wcc.logger.Info("ElasticsearchClientController HealthHandler")

	ginContext.JSON(http.StatusOK, gin.H{
		"message": wcc.elasticsearchClientService.Health(),
	})
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/elasticsearch-client/send?message=hello' | jq
// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3002/elasticsearch-client/send' -d '{"message":"hello"}' | jq
// @Summary elasticsearch client send message
// @Description elasticsearch client send message
// @Tags elasticsearch-client
// @Accept json
// @Produce json
// @Param message query string true "message"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /elasticsearch-client/send [get]
func (wcc *ElasticsearchClientController) Send(ginContext *gin.Context) {
	wcc.logger.Info("ElasticsearchClientController SendHandler")

	message := ginContext.Query("message")
	message, err := wcc.elasticsearchClientService.Send(message)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// var data model.ElasticsearchMessageDTO
	// if err := ginContext.ShouldBindJSON(&webhookData); err != nil {
	//	ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	// }

	// message, err	 := wcc.elasticsearchClientService.Send(data)
	// if err != nil {
	//	ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	// }

	ginContext.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
