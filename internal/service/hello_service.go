package service

import (
	"context"

	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/sirupsen/logrus"
)

type HelloServiceInterface interface {
	GetHelloString() (model.HelloStringResponse, error)
	GetHelloJson() (model.HelloJsonResponse, error)
}

type HelloService struct {
	ctx    context.Context
	logger *logrus.Logger
}

func NewHelloService(ctx context.Context, logger *logrus.Logger) *HelloService {
	return &HelloService{
		ctx:    ctx,
		logger: logger,
	}
}

func (hs *HelloService) GetHelloString() (model.HelloStringResponse, error) {
	return model.HelloStringResponse("Hello, World!"), nil
}

func (hs *HelloService) GetHelloJson() (model.HelloJsonResponse, error) {
	response := model.HelloJsonResponse{
		Data: model.Data{
			Message: "Hello, World!",
		},
	}

	return response, nil
}
