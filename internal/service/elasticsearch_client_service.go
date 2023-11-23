package service

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type ElasticsearchClientInterface interface {
	NewElasticsearchClientService(ctx context.Context, logger *logrus.Logger) *ElasticsearchClientService
	Initialize() error
	Connection() error
	Disconnection() error
	Send(message string) (string, error)
	Echo() (string, error)
	Broadcast() (string, error)
	Health() string
}

type ElasticsearchClientService struct {
	ctx                           context.Context
	logger                        *logrus.Logger
	elasticsearchServerServiceURL string
}

func NewElasticsearchClientService(ctx context.Context, logger *logrus.Logger) *ElasticsearchClientService {
	return &ElasticsearchClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (gcs *ElasticsearchClientService) Initialize() error {
	gcs.logger.Info("ElasticsearchClientService Initialize")
	gcs.elasticsearchServerServiceURL = gcs.initializeElasticsearchServerServiceURL()
	gcs.logger.Infof("Elasticsearch Server URL: %s\n", gcs.elasticsearchServerServiceURL)
	return nil
}

func (gcs *ElasticsearchClientService) initializeElasticsearchServerServiceURL() string {
	SERVER_HOST := os.Getenv("SERVER_HOST")
	if SERVER_HOST == "" {
		SERVER_HOST = "0.0.0.0"
	}
	SERVER_PORT := os.Getenv("SERVER_PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "3001"
	}
	return fmt.Sprintf("http://%s:%s/elasticsearch/handler", SERVER_HOST, SERVER_PORT)
}

func (gcs *ElasticsearchClientService) Connection() error {
	gcs.logger.Info("ElasticsearchClientService Connection")
	return nil
}

func (gcs *ElasticsearchClientService) Disconnection() error {
	gcs.logger.Info("ElasticsearchClientService Disconnection")
	return nil
}

func (gcs *ElasticsearchClientService) Send(message string) (string, error) {
	gcs.logger.Info("ElasticsearchClientService Send")
	return "", nil
}

func (gcs *ElasticsearchClientService) Echo() (string, error) {
	gcs.logger.Info("ElasticsearchClientService Echo")
	return "", nil
}

func (gcs *ElasticsearchClientService) Broadcast() (string, error) {
	gcs.logger.Info("ElasticsearchClientService Broadcast")
	return "", nil
}

func (gcs *ElasticsearchClientService) Health() string {
	gcs.logger.Info("ElasticsearchClientService Health")
	return "OK"
}
