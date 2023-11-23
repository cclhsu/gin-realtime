package service

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type WebRTCClientInterface interface {
	NewWebRTCClientService(ctx context.Context, logger *logrus.Logger) *WebRTCClientService
	Initialize() error
	Connection() error
	Disconnection() error
	Send(message string) (string, error)
	Echo() (string, error)
	Broadcast() (string, error)
	Health() string
}

type WebRTCClientService struct {
	ctx                    context.Context
	logger                 *logrus.Logger
	webrtcServerServiceURL string
}

func NewWebRTCClientService(ctx context.Context, logger *logrus.Logger) *WebRTCClientService {
	return &WebRTCClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (wcs *WebRTCClientService) Initialize() error {
	wcs.logger.Info("WebRTCClientService Initialize")
	wcs.webrtcServerServiceURL = wcs.initializeWebRTCServerServiceURL()
	wcs.logger.Infof("WebRTC Server URL: %s\n", wcs.webrtcServerServiceURL)
	return nil
}

func (wcs *WebRTCClientService) initializeWebRTCServerServiceURL() string {
	SERVER_HOST := os.Getenv("SERVER_HOST")
	if SERVER_HOST == "" {
		SERVER_HOST = "0.0.0.0"
	}
	SERVER_PORT := os.Getenv("SERVER_PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "3001"
	}
	return fmt.Sprintf("http://%s:%s/webrtc/handler", SERVER_HOST, SERVER_PORT)
}

func (wcs *WebRTCClientService) Connection() error {
	wcs.logger.Info("WebRTCClientService Connection")
	return nil
}

func (wcs *WebRTCClientService) Disconnection() error {
	wcs.logger.Info("WebRTCClientService Disconnection")
	return nil
}

func (wcs *WebRTCClientService) Send(message string) (string, error) {
	wcs.logger.Info("WebRTCClientService Send")
	return "WebRTCClientService Send", nil
}

func (wcs *WebRTCClientService) Echo() (string, error) {
	wcs.logger.Info("WebRTCClientService Echo")
	return "WebRTCClientService Echo", nil
}

func (wcs *WebRTCClientService) Broadcast() (string, error) {
	wcs.logger.Info("WebRTCClientService Broadcast")
	return "WebRTCClientService Broadcast", nil
}

func (wcs *WebRTCClientService) Health() string {
	wcs.logger.Info("WebRTCClientService Health")
	return "WebRTCClientService Health"
}
