package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	// _interface "github.com/cclhsu/gin-realtime/internal/interface"

	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/types"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WebsocketClientInterface interface {
	NewWebsocketClientService(ctx context.Context, logger *logrus.Logger) *WebsocketClientService
	// _interface.ConnectionServiceInterface
	// _interface.MessageServiceInterface
	// _interface.HealthServiceInterface
	Initialize() error
	Connection() error
	Disconnection() error
	Send(message model.MessageDTO) error
	Receive() error
	ListMessages() error
	Health() string
}

type WebsocketClientService struct {
	ctx                       context.Context
	logger                    *logrus.Logger
	websocketServerServiceURL string
	clientID                  string
	websocketConn             *websocket.Conn
	closeSignal               chan struct{}
	errorChannel              chan error
}

func NewWebsocketClientService(ctx context.Context, logger *logrus.Logger) *WebsocketClientService {
	return &WebsocketClientService{
		ctx:          ctx,
		logger:       logger,
		closeSignal:  make(chan struct{}),
		errorChannel: make(chan error),
	}
}

func (wcs *WebsocketClientService) Initialize() error {
	wcs.logger.Info("WebsocketClientService Initialize")
	wcs.websocketServerServiceURL = wcs.initializeWebsocketServerServiceURL()
	wcs.logger.Infof("Websocket Server URL: %s\n", wcs.websocketServerServiceURL)

	wcs.clientID = uuid.New().String()

	// Connect to the server
	go wcs.Connection()

	// Wait until the connection is established
	wcs.logger.Info("Waiting for connection to be established...")
	for wcs.websocketConn == nil {
		time.Sleep(1 * time.Second)
	}
	wcs.logger.Info("Connection established.")

	// Read messages from the server
	go wcs.Receive()

	// Send a message to the server (optional)
	message := model.MessageDTO{
		UUID:          "00000000-0000-0000-0000-000000000000",
		Type:          types.MESSAGE_TYPES_ECHO,
		Action:        types.ACTION_TYPES_UNSPECIFIED,
		Stage:         types.STAGE_TYPES_UNSPECIFIED,
		Environment:   types.ENVIRONMENT_TYPES_UNSPECIFIED,
		Sender:        wcs.clientID,
		Recipient:     "00000000-0000-0000-0000-000000000000",
		Recipients:    []string{},
		RecipientType: types.RECIPIENT_TYPES_UNSPECIFIED,
		Data:          map[string]interface{}{"message": "hello world"},
		Metadata:      map[string]interface{}{"additionalProp1": map[string]interface{}{}},
	}
	err := wcs.Send(message)
	if err != nil {
		wcs.logger.Errorf("Error: %v", err)
		return err
	}

	// defer wcs.Disconnection()
	return nil
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

func (wcs *WebsocketClientService) Connection() error {
	wcs.logger.Info("WebSocketClientService Connection")

	// Create a WebSocket dialer with any necessary options
	dialer := websocket.DefaultDialer

	for {
		select {
		case <-wcs.closeSignal:
			// Received a signal to close the connection
			wcs.logger.Info("Closing WebSocket connection: received a signal to close the connection.")
			return wcs.Disconnection()

		case <-wcs.ctx.Done():
			// The application context was canceled
			wcs.logger.Info("Closing WebSocket connection: application context canceled.")
			return wcs.Disconnection()

		default:
			// Attempt to establish the WebSocket connection
			var err error
			wcs.websocketConn, _, err = dialer.Dial(wcs.websocketServerServiceURL, nil)
			if err != nil {
				wcs.logger.Errorf("Error: %v", err)
				wcs.errorChannel <- fmt.Errorf("error establishing WebSocket connection: %w", err)
				// You may choose to retry the connection or handle the error differently.
				// For simplicity, this example doesn't include retries.
				time.Sleep(5 * time.Second) // Wait before retrying
				continue
			}
			wcs.logger.Infof("Connected to %s", wcs.websocketServerServiceURL)
			return nil
		}
	}
}

func (wcs *WebsocketClientService) Disconnection() error {
	wcs.logger.Info("WebsocketClientService Disconnection")

	if wcs.websocketConn != nil {
		// Close connection
		err := wcs.websocketConn.Close()
		if err != nil {
			wcs.logger.Errorf("Error: %v", err)
			return err
		}

		wcs.websocketConn = nil
		wcs.logger.Infof("Disconnected from %s", wcs.websocketServerServiceURL)
	}

	return nil
}

func (wcs *WebsocketClientService) Send(message model.MessageDTO) error {
	wcs.logger.Infof("WebsocketClientService Send: %v", message)

	if wcs.websocketConn == nil {
		return fmt.Errorf("WebSocket connection not established")
	}

	// Send message to server
	err := wcs.websocketConn.WriteJSON(message)
	if err != nil {
		wcs.logger.Errorf("Error: %v", err)
		return err
	}

	wcs.logger.Infof("Sent message: %v", message)
	return nil
}

func (wcs *WebsocketClientService) Receive() error {
	wcs.logger.Info("WebsocketClientService Receive")

	defer func() {
		if err := wcs.websocketConn.Close(); err != nil {
			wcs.logger.Error("Failed to close WebSocket connection:", err)
		}
		wcs.logger.Info("WebSocket connection closed successfully: %s", wcs.websocketConn.RemoteAddr().String())
	}()

	for {
		select {
		case <-wcs.closeSignal:
			// Received a signal to close the connection
			wcs.logger.Info("Closing WebSocket connection: received a signal to close the connection.")
			return wcs.Disconnection()

		case <-wcs.ctx.Done():
			// The application context was canceled
			wcs.logger.Info("Closing WebSocket connection: application context canceled.")
			return wcs.Disconnection()

		case err := <-wcs.errorChannel:
			// Received an error from another goroutine
			wcs.logger.Errorf("Closing WebSocket connection: received an error from another goroutine: %v", err)
			return wcs.Disconnection()

		case <-time.After(10 * time.Second):
			// No incoming message within the deadline
			wcs.logger.Info("Closing WebSocket connection: no incoming message within the deadline.")
			return wcs.Disconnection()

		default:
			// Read message from WebSocket
			var message model.MessageDTO
			err := wcs.websocketConn.ReadJSON(&message)
			if err != nil {
				// Check for timeout error (no incoming message)
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					wcs.logger.Errorf("Error reading WebSocket message: %v", err)
					wcs.errorChannel <- fmt.Errorf("error reading WebSocket message: %w", err)
					return err
				} else if strings.Contains(err.Error(), "i/o timeout") {
					// No incoming message within the deadline
					wcs.logger.Infof("No incoming message within the deadline: %v", err)
					continue
				}
				wcs.logger.Errorf("Error reading WebSocket message: %v", err)
				wcs.errorChannel <- fmt.Errorf("error reading WebSocket message: %w", err)
				return err
			}
			wcs.logger.Infof("Received message: %v", message)

			// Log the raw JSON message received from the server
			rawMessage, err := json.Marshal(message)
			if err != nil {
				wcs.logger.Errorf("Error marshaling message: %v", err)
				wcs.errorChannel <- fmt.Errorf("error marshaling message: %w", err)
				return err
			}
			wcs.logger.Infof("Raw JSON message: %s", rawMessage)
		}
	}
}

func (wcs *WebsocketClientService) ListMessages() error {
	wcs.logger.Info("WebsocketClientService ListMessages")
	return fmt.Errorf("Unimplemented method")
}

func (wcs *WebsocketClientService) Health() string {
	wcs.logger.Info("WebsocketClientService Health")
	return "OK"
}
