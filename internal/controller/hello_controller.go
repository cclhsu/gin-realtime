package controller

import (
	"context"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HelloControllerInterface interface {
	GetHelloJson(c *gin.Context)
	GetHelloString(c *gin.Context)
}

type HelloController struct {
	ctx			 context.Context
	logger		 *logrus.Logger
	helloService *service.HelloService
}

func NewHelloController(ctx context.Context, logger *logrus.Logger, helloService *service.HelloService) *HelloController {
	return &HelloController{
		ctx:		  ctx,
		logger:		  logger,
		helloService: helloService,
	}
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/hello/json' | jq
// @Summary get hello world json
// @Description get hello world json
// @Tags hello
// @Accept json
// @Produce json
// @Success 200 {object} model.HelloJsonResponse
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /hello/json [get]
func (hc *HelloController) GetHelloJson(c *gin.Context) {
	response, _ := hc.helloService.GetHelloJson()
	c.JSON(http.StatusOK, response)
}

// curl -s -X 'GET' -H 'accept: application/json' 'http://0.0.0.0:3001/hello/string' -w "\n"
// @Summary get hello world string
// @Description get hello world string
// @Tags hello
// @Accept json
// @Produce plain
// @Success 200 {object} model.HelloStringResponse
// @Failure 400 {object} string "Invalid request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal Server Error"
// @Router /hello/string [get]
func (hc *HelloController) GetHelloString(c *gin.Context) {
	response, _ := hc.helloService.GetHelloString()
	c.String(http.StatusOK, string(response))
}
