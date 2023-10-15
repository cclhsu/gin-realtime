package service

import (
	"context"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WebsocketClientInterface interface {
	NewWebsocketClientService(ctx context.Context, logger *logrus.Logger) *WebsocketClientService
	Initialize()
	Connection() (*websocket.Conn, error)
	Disconnection() (*websocket.Conn, error)
	Trigger(message string) (*websocket.Conn, string, error)
	Echo() (*websocket.Conn, string, error)
	Broadcast() (*websocket.Conn, string, error)
	Health() string
}

type WebsocketClientService struct {
	ctx                       context.Context
	logger                    *logrus.Logger
	websocketServerServiceURL string
	websocketConn             *websocket.Conn
}

func NewWebsocketClientService(ctx context.Context, logger *logrus.Logger) *WebsocketClientService {
	return &WebsocketClientService{
		ctx:    ctx,
		logger: logger,
	}
}

func (wcs *WebsocketClientService) Initialize() {
	wcs.logger.Info("WebsocketClientService Initialize")
	wcs.websocketServerServiceURL = wcs.initializeWebsocketServerServiceURL()
	wcs.logger.Infof("Websocket Server URL: %s\n", wcs.websocketServerServiceURL)

	// connect to server
	go wcs.Connection()

	// defer wcs.Disconnection()
}

func (wcs *WebsocketClientService) initializeWebsocketServerServiceURL() string {
	SERVER_HOST := os.Getenv("SERVER_HOST")
	if SERVER_HOST == "" {
		SERVER_HOST = "0.0.0.0"
	}
	SERVER_PORT := os.Getenv("SERVER_PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "3001"
	}
	return fmt.Sprintf("ws://%s:%s/websocket/ws", SERVER_HOST, SERVER_PORT)
}

func (wcs *WebsocketClientService) Connection() (*websocket.Conn, chan struct{}) {
	wcs.logger.Info("WebsocketClientService Connection")

	// Create a channel to signal when the connection should be closed
	closeSignal := make(chan struct{})

	// Create a goroutine to handle the WebSocket connection
	go func() {
		// Create a WebSocket dialer with any necessary options
		dialer := websocket.DefaultDialer
		// Set dialer options if needed
		// dialer.Proxy = http.ProxyFromEnvironment

		for {
			select {
			case <-closeSignal:
				// Received a signal to close the connection
				wcs.logger.Info("Closing WebSocket connection.")
				return
			default:
				// Attempt to establish the WebSocket connection
				var err error
				wcs.websocketConn, _, err = dialer.Dial(wcs.websocketServerServiceURL, nil)
				if err != nil {
					wcs.logger.Errorf("Error: %v", err)
					// You can choose to retry the connection or handle the error differently.
					// For simplicity, this example doesn't include retries.
					return
				}
				wcs.logger.Infof("Connected to %s", wcs.websocketServerServiceURL)

				err = wcs.websocketConn.WriteMessage(websocket.TextMessage, []byte("hello world"))
				if err != nil {
					wcs.logger.Errorf("Error: %v", err)
					return
				}

				// read message from server
				_, p, err := wcs.websocketConn.ReadMessage()
				if err != nil {
					wcs.logger.Errorf("Error: %v", err)
					return
				}
				wcs.logger.Infof("Received message: %s", p)

				// Keep the connection open until signaled to close
				for {
					select {
					case <-closeSignal:
						// Received a signal to close the connection
						wcs.Disconnection()
						wcs.logger.Info("WebSocket connection closed.")
						return
					}
				}
			}
		}
	}()

	return wcs.websocketConn, closeSignal
}

// func (wcs *WebsocketClientService) Connection() (*websocket.Conn, error) {
// 	wcs.logger.Info("WebsocketClientService Connection")

// 	// connect to server
// 	var err error
// 	wcs.websocketConn, _, err = websocket.DefaultDialer.Dial(wcs.websocketServerServiceURL, nil)
// 	if err != nil {
// 		wcs.logger.Errorf("Error: %v", err)
// 		return nil, err
// 	}

// 	wcs.logger.Infof("Connected to %s", wcs.websocketServerServiceURL)
// 	// defer wcs.Disconnection()
// 	defer wcs.websocketConn.Close()

// 	wcs.websocketConn.SetCloseHandler(func(code int, text string) error {
// 		wcs.logger.Infof("CloseHandler: %v, %v", code, text)
// 		return nil
// 	})

// 	wcs.websocketConn.SetPingHandler(func(appData string) error {
// 		wcs.logger.Infof("PingHandler: %v", appData)
// 		return nil
// 	})

// 	err = wcs.websocketConn.WriteMessage(websocket.TextMessage, []byte("hello world"))
// 	if err != nil {
// 		wcs.logger.Errorf("Error: %v", err)
// 		return nil, err
// 	}

// 	// read message from server
// 	_, p, err := wcs.websocketConn.ReadMessage()
// 	if err != nil {
// 		wcs.logger.Errorf("Error: %v", err)
// 		return nil, err
// 	}
// 	wcs.logger.Infof("Received message: %s", p)

// 	return wcs.websocketConn, nil
// }

func (wcs *WebsocketClientService) Disconnection() (*websocket.Conn, error) {
	wcs.logger.Info("WebsocketClientService Disconnection")

	// close connection
	err := wcs.websocketConn.Close()
	if err != nil {
		wcs.logger.Errorf("Error: %v", err)
		return nil, err
	}

	wcs.websocketConn = nil

	wcs.logger.Infof("Disconnected from %s", wcs.websocketServerServiceURL)
	return wcs.websocketConn, nil
}

func (wcs *WebsocketClientService) Trigger(message string) (*websocket.Conn, string, error) {
	wcs.logger.Info("WebsocketClientService Trigger" + message)

	// send message to server
	err := wcs.websocketConn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		wcs.logger.Errorf("Error: %v", err)
		return nil, "", err
	}

	// read message from server
	_, p, err := wcs.websocketConn.ReadMessage()
	if err != nil {
		wcs.logger.Errorf("Error: %v", err)
		return nil, "", err
	}

	wcs.logger.Infof("Received message: %s", p)
	return wcs.websocketConn, string(p), nil
}

func (wcs *WebsocketClientService) Echo() {
	wcs.logger.Info("WebsocketClientService Echo")
}

func (wcs *WebsocketClientService) Broadcast() {
	wcs.logger.Info("WebsocketClientService Broadcast")
}

func (wcs *WebsocketClientService) Health() string {
	wcs.logger.Info("WebsocketClientService Health")
	return "OK"
}
