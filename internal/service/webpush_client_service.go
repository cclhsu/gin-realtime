package service

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type WebpushClientInterface interface {
	NewWebpushClientService(ctx context.Context, logger *logrus.Logger) *WebpushClientService
	Initialize() error
	Connection() error
	Disconnection() error
	Send(message string) (string, error)
	Echo() (string, error)
	Broadcast() (string, error)
	Health() string
}

type WebpushClientService struct {
	ctx                     context.Context
	logger                  *logrus.Logger
	webpushServerServiceURL string
}

func NewWebpushClientService(ctx context.Context, logger *logrus.Logger) *WebpushClientService {
	return &WebpushClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (wcs *WebpushClientService) Initialize() error {
	wcs.logger.Info("WebpushClientService Initialize")
	wcs.webpushServerServiceURL = wcs.initializeWebpushServerServiceURL()
	wcs.logger.Infof("Webpush Server URL: %s\n", wcs.webpushServerServiceURL)
	return nil
}

func (wcs *WebpushClientService) initializeWebpushServerServiceURL() string {
	SERVER_HOST := os.Getenv("SERVER_HOST")
	if SERVER_HOST == "" {
		SERVER_HOST = "0.0.0.0"
	}
	SERVER_PORT := os.Getenv("SERVER_PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "3001"
	}
	return fmt.Sprintf("http://%s:%s/webpush/handler", SERVER_HOST, SERVER_PORT)
}

func (wcs *WebpushClientService) Connection() error {
	wcs.logger.Info("WebpushClientService Connection")
	return nil
}

func (wcs *WebpushClientService) Disconnection() error {
	wcs.logger.Info("WebpushClientService Disconnection")
	return nil
}

func (wcs *WebpushClientService) Send(message string) (string, error) {
	wcs.logger.Info("WebpushClientService Send")
	return "WebpushClientService Send", nil
}

func (wcs *WebpushClientService) Echo() (string, error) {
	wcs.logger.Info("WebpushClientService Echo")
	return "WebpushClientService Echo", nil
}

func (wcs *WebpushClientService) Broadcast() (string, error) {
	wcs.logger.Info("WebpushClientService Broadcast")
	return "WebpushClientService Broadcast", nil
}

func (wcs *WebpushClientService) Health() string {
	wcs.logger.Info("WebpushClientService Health")
	return "WebpushClientService Health"
}
