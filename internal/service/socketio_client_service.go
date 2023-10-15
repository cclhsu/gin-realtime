package service

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type SocketIOClientInterface interface {
	NewSocketIOClientService(ctx context.Context, logger *logrus.Logger) *SocketIOClientService
	Initialize()
	Connection() error
	Disconnection() error
	Trigger(message string) (string, error)
	Echo() (string, error)
	Broadcast() (string, error)
	Health() string
}

type SocketIOClientService struct {
	ctx                      context.Context
	logger                   *logrus.Logger
	socketIOServerServiceURL string
}

func NewSocketIOClientService(ctx context.Context, logger *logrus.Logger) *SocketIOClientService {
	return &SocketIOClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (scs *SocketIOClientService) Initialize() {
	scs.logger.Info("SocketIOClientService Initialize")
	scs.socketIOServerServiceURL = scs.initializeSocketIOServerServiceURLL()
	scs.logger.Infof("SOCKET_IO Server URL: %s\n", scs.socketIOServerServiceURL)
}

func (scs *SocketIOClientService) initializeSocketIOServerServiceURLL() string {
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

func (scs *SocketIOClientService) Connection() error {
	scs.logger.Info("SocketIOClientService Connection")
	return nil
}

func (scs *SocketIOClientService) Disconnection() error {
	scs.logger.Info("SocketIOClientService Disconnection")
	return nil
}

func (scs *SocketIOClientService) Trigger(message string) (string, error) {
	scs.logger.Info("SocketIOClientService Trigger")
	return "", nil
}

func (scs *SocketIOClientService) Echo() (string, error) {
	scs.logger.Info("SocketIOClientService Echo")
	return "", nil
}

func (scs *SocketIOClientService) Broadcast() (string, error) {
	scs.logger.Info("SocketIOClientService Broadcast")
	return "", nil
}
