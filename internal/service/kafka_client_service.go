package service

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type KafkaClientInterface interface {
	NewKafkaClientService(ctx context.Context, logger *logrus.Logger) *KafkaClientService
	Initialize()
	Connection() error
	Disconnection() error
	Trigger(message string) (string, error)
	Echo() (string, error)
	Broadcast() (string, error)
	Health() string
}

type KafkaClientService struct {
	ctx                   context.Context
	logger                *logrus.Logger
	kafkaServerServiceURL string
}

func NewKafkaClientService(ctx context.Context, logger *logrus.Logger) *KafkaClientService {
	return &KafkaClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (kcs *KafkaClientService) Initialize() {
	kcs.logger.Info("KafkaClientService Initialize")
	kcs.kafkaServerServiceURL = kcs.initializeKafkaServerServiceURL()
	kcs.logger.Infof("Kafka Server URL: %s\n", kcs.kafkaServerServiceURL)
}

func (kcs *KafkaClientService) initializeKafkaServerServiceURL() string {
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

func (kcs *KafkaClientService) Connection() error {
	kcs.logger.Info("KafkaClientService Connection")
	return nil
}

func (kcs *KafkaClientService) Disconnection() error {
	kcs.logger.Info("KafkaClientService Disconnection")
	return nil
}

func (kcs *KafkaClientService) Trigger(message string) (string, error) {
	kcs.logger.Info("KafkaClientService Trigger")
	return "KafkaClientService Trigger", nil
}

func (kcs *KafkaClientService) Echo() (string, error) {
	kcs.logger.Info("KafkaClientService Echo")
	return "KafkaClientService Echo", nil
}

func (kcs *KafkaClientService) Broadcast() (string, error) {
	kcs.logger.Info("KafkaClientService Broadcast")
	return "KafkaClientService Broadcast", nil
}

func (kcs *KafkaClientService) Health() string {
	kcs.logger.Info("KafkaClientService Health")
	return "KafkaClientService Health"
}
