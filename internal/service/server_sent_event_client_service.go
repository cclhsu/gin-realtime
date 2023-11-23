package service

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type ServerSentEventClientInterface interface {
	NewServerSentEventClientService(ctx context.Context, logger *logrus.Logger) *ServerSentEventClientService
	Initialize() error
	Connection() error
	Disconnection() error
	Send(message string) (string, error)
	Echo() (string, error)
	Broadcast() (string, error)
	Health() string
}

type ServerSentEventClientService struct {
	ctx                             context.Context
	logger                          *logrus.Logger
	serverSentEventServerServiceURL string
}

func NewServerSentEventClientService(ctx context.Context, logger *logrus.Logger) *ServerSentEventClientService {
	return &ServerSentEventClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (ssecs *ServerSentEventClientService) Initialize() error {
	ssecs.logger.Info("ServerSentEventClientService Initialize")
	ssecs.serverSentEventServerServiceURL = ssecs.initializeServerSentEventServerServiceURL()
	ssecs.logger.Infof("ServerSentEvent Server URL: %s\n", ssecs.serverSentEventServerServiceURL)
	return nil
}

func (ssecs *ServerSentEventClientService) initializeServerSentEventServerServiceURL() string {
	SERVER_HOST := os.Getenv("SERVER_HOST")
	if SERVER_HOST == "" {
		SERVER_HOST = "0.0.0.0"
	}
	SERVER_PORT := os.Getenv("SERVER_PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "3001"
	}
	return fmt.Sprintf("http://%s:%s/server-sent-event/handler", SERVER_HOST, SERVER_PORT)
}

func (ssecs *ServerSentEventClientService) Connection() error {
	ssecs.logger.Info("ServerSentEventClientService Connection")
	return nil
}

func (ssecs *ServerSentEventClientService) Disconnection() error {
	ssecs.logger.Info("ServerSentEventClientService Disconnection")
	return nil
}

func (ssecs *ServerSentEventClientService) Send(message string) (string, error) {
	ssecs.logger.Info("ServerSentEventClientService Send")
	return "", nil
}

func (ssecs *ServerSentEventClientService) Echo() (string, error) {
	ssecs.logger.Info("ServerSentEventClientService Echo")
	return "ServerSentEventClientService Echo", nil
}

func (ssecs *ServerSentEventClientService) Broadcast() (string, error) {
	ssecs.logger.Info("ServerSentEventClientService Broadcast")
	return "ServerSentEventClientService Broadcast", nil
}

func (ssecs *ServerSentEventClientService) Health() string {
	ssecs.logger.Info("ServerSentEventClientService Health")
	return "ServerSentEventClientService Health"
}
