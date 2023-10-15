package service

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type GrpcClientInterface interface {
	NewGrpcClientService(ctx context.Context, logger *logrus.Logger) *GrpcClientService
	Initialize()
	Connection() error
	Disconnection() error
	Trigger(message string) (string, error)
	Echo() (string, error)
	Broadcast() (string, error)
	Health() string
}

type GrpcClientService struct {
	ctx                  context.Context
	logger               *logrus.Logger
	grpcServerServiceURL string
}

func NewGrpcClientService(ctx context.Context, logger *logrus.Logger) *GrpcClientService {
	return &GrpcClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (gcs *GrpcClientService) Initialize() {
	gcs.logger.Info("GrpcClientService Initialize")
	gcs.grpcServerServiceURL = gcs.initializeGrpcServerServiceURL()
	gcs.logger.Infof("Grpc Server URL: %s\n", gcs.grpcServerServiceURL)
}

func (gcs *GrpcClientService) initializeGrpcServerServiceURL() string {
	SERVER_HOST := os.Getenv("SERVER_HOST")
	if SERVER_HOST == "" {
		SERVER_HOST = "0.0.0.0"
	}
	SERVER_PORT := os.Getenv("SERVER_PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "3001"
	}
	return fmt.Sprintf("http://%s:%s/ws", SERVER_HOST, SERVER_PORT)
}

func (gcs *GrpcClientService) Connection() error {
	gcs.logger.Info("GrpcClientService Connection")
	return nil
}

func (gcs *GrpcClientService) Disconnection() error {
	gcs.logger.Info("GrpcClientService Disconnection")
	return nil
}

func (gcs *GrpcClientService) Trigger(message string) (string, error) {
	gcs.logger.Info("GrpcClientService Trigger")
	return "", nil
}

func (gcs *GrpcClientService) Echo() (string, error) {
	gcs.logger.Info("GrpcClientService Echo")
	return "", nil
}

func (gcs *GrpcClientService) Broadcast() (string, error) {
	gcs.logger.Info("GrpcClientService Broadcast")
	return "", nil
}

func (gcs *GrpcClientService) Health() string {
	gcs.logger.Info("GrpcClientService Health")
	return "OK"
}
