package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ElasticsearchServerControllerInterface interface {
	ElasticsearchHandler(ginContext *gin.Context)
}

type ElasticsearchServerController struct {
	ctx						   context.Context
	logger					   *logrus.Logger
	elasticsearchServerService *service.ElasticsearchServerService
}

func NewElasticsearchServerController(ctx context.Context, logger *logrus.Logger, elasticsearchServerService *service.ElasticsearchServerService) *ElasticsearchServerController {
	return &ElasticsearchServerController{
		ctx:						ctx,
		logger:						logger,
		elasticsearchServerService: elasticsearchServerService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/elasticsearch/handler' | jq
// @Summary elasticsearch handler
// @Description elasticsearch handler
// @Tags elasticsearch
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /elasticsearch/handler [get]
func (wsc *ElasticsearchServerController) ElasticsearchHandler(ginContext *gin.Context) {
	wsc.logger.Info("ElasticsearchHandler")
	// wsc.elasticsearchServerService.ElasticsearchHandler(ginContext)
}
