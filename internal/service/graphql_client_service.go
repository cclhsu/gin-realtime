package service

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type GraphQLClientInterface interface {
	NewGraphQLClientService(ctx context.Context, logger *logrus.Logger) *GraphQLClientService
	Initialize()
	Connection() error
	Disconnection() error
	Trigger(message string) (string, error)
	Echo() (string, error)
	Broadcast() (string, error)
	Health() string
}

type GraphQLClientService struct {
	ctx                     context.Context
	logger                  *logrus.Logger
	graphQLServerServiceURL string
}

func NewGraphQLClientService(ctx context.Context, logger *logrus.Logger) *GraphQLClientService {
	return &GraphQLClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (gcs *GraphQLClientService) Initialize() {
	gcs.logger.Info("graphQLClientService Initialize")
	gcs.graphQLServerServiceURL = gcs.initializeGraphQLServerURL()
	gcs.logger.Infof("GraphQL Server URL: %s\n", gcs.graphQLServerServiceURL)
}

func (gcs *GraphQLClientService) initializeGraphQLServerURL() string {
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

func (gcs *GraphQLClientService) Connection() error {
	gcs.logger.Info("graphQLClientService Connection")
	return nil
}

func (gcs *GraphQLClientService) Disconnection() error {
	gcs.logger.Info("graphQLClientService Disconnection")
	return nil
}

func (gcs *GraphQLClientService) Trigger(message string) (string, error) {
	gcs.logger.Info("graphQLClientService Trigger")
	return "", nil
}

func (gcs *GraphQLClientService) Echo() (string, error) {
	gcs.logger.Info("graphQLClientService Echo")
	return "", nil
}

func (gcs *GraphQLClientService) Broadcast() (string, error) {
	gcs.logger.Info("graphQLClientService Broadcast")
	return "", nil
}

func (gcs *GraphQLClientService) Health() string {
	gcs.logger.Info("graphQLClientService Health")
	return "OK"
}
