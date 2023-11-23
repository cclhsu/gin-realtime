package controller

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GrpcServerControllerInterface interface {
	GrpcHandler(ginContext *gin.Context)
}

type GrpcServerController struct {
	ctx				  context.Context
	logger			  *logrus.Logger
	grpcServerService *service.GrpcServerService
}

func NewGrpcServerController(ctx context.Context, logger *logrus.Logger, grpcServerService *service.GrpcServerService) *GrpcServerController {
	return &GrpcServerController{
		ctx:			   ctx,
		logger:			   logger,
		grpcServerService: grpcServerService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/grpc/handler' | jq
// @Summary grpc handler
// @Description grpc handler
// @Tags grpc
// @Accept json
// @Produce json
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /grpc/handler [get]
func (wsc *GrpcServerController) GrpcHandler(ginContext *gin.Context) {
	wsc.logger.Info("GrpcHandler")
	// wsc.grpcServerService.GrpcHandler(ginContext)
}
