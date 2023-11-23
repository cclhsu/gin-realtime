package service

import (
	"context"
	"encoding/json"
	"net/http"

	// "net/url"

	"github.com/cclhsu/gin-realtime/internal/model"
	"github.com/cclhsu/gin-realtime/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	// "github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type WebsocketServerInterface interface {
	NewWebsocketServerService(ctx context.Context, logger *logrus.Logger) *WebsocketServerService
	// _interface.ConnectionServiceInterface
	// _interface.MessageServiceInterface
	// _interface.HealthServiceInterface
	// Initialize() error
	WebsocketHandler(ginContext *gin.Context)
	handleConnection(websocketConn *websocket.Conn)
	handleDisconnection(websocketConn *websocket.Conn)
	send(websocketConn *websocket.Conn)
	handleBroadcast(websocketConn *websocket.Conn)
	handleUnicast(websocketConn *websocket.Conn)
	// handleHealth(websocketConn *websocket.Conn)
}

type WebsocketServerService struct {
	ctx             context.Context
	logger          *logrus.Logger
	sessionGroupMap map[string]map[uuid.UUID]*websocket.Conn
	upgrader        *websocket.Upgrader

	// //The client map stores and manages all long connection clients, online is TRUE, and those who are not there are FALSE
	// clients map[*Client]bool
	// //Web side MESSAGE we use Broadcast to receive, and finally distribute it to all clients
	// broadcast chan []byte
	// //Newly created long connection client
	// register chan *Client
	// //Newly canceled long connection client
	// unregister chan *Client
}

// //Client
// type Client struct {
// 	//User ID
// 	id string
// 	//Connected socket
// 	socket *websocket.Conn
// 	//Message
// 	send chan []byte
// }

func NewWebsocketServerService(ctx context.Context, logger *logrus.Logger) *WebsocketServerService {
	// SERVICE_HOST := os.Getenv("SERVICE_HOST")
	// if SERVICE_HOST == "" {
	//	SERVICE_HOST = "localhost"
	// }
	// SERVICE_PORT := os.Getenv("SERVICE_PORT")
	// if SERVICE_PORT == "" {
	//	SERVICE_PORT = "8080"
	// }
	return &WebsocketServerService{
		ctx:             ctx,
		logger:          logger,
		sessionGroupMap: make(map[string]map[uuid.UUID]*websocket.Conn),
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024 * 1024 * 1024,
			WriteBufferSize: 1024 * 1024 * 1024,
			// You can add custom logic here to check the origin
			CheckOrigin: func(r *http.Request) bool {
				// origin := r.Header.Get("Origin")
				// u, err := url.Parse(origin)
				// if err != nil {
				//	return false
				// }
				// return u.Host == SERVICE_HOST+":"+SERVICE_PORT\
				return true
			},
		},

		// broadcast:  make(chan []byte),
		// register:   make(chan *Client),
		// unregister: make(chan *Client),
		// clients:    make(map[*Client]bool),
	}
}

// func (wss *WebsocketServerService) Initialize() error {
// 	wss.logger.Infof("WebsocketServerService Initialize")
// 	return nil
// }

func (wss *WebsocketServerService) WebsocketHandler(ginContext *gin.Context) {
	websocketConn, err := wss.upgrader.Upgrade(ginContext.Writer, ginContext.Request, nil)
	if err != nil {
		wss.logger.Error("Failed to upgrade connection:", err)
		return
	}
	defer websocketConn.Close()

	// Handle connection
	wss.handleConnection(websocketConn)

	// Handle disconnection
	defer wss.handleDisconnection(websocketConn)

	// Handle Send messages
	wss.Send(websocketConn)

	// Handle health
	// wss.handleHealth(websocketConn)
}

func (wss *WebsocketServerService) handleConnection(websocketConn *websocket.Conn) {
	// Generate a unique session ID for the connection
	sessionID := uuid.New()

	// Log the new connection
	wss.logger.Infof("New connection from %s, sessionID: %s", websocketConn.RemoteAddr().String(), sessionID.String())

	// Initialize the session and add the connection to the session group
	wss.sessionGroupMap[sessionID.String()] = map[uuid.UUID]*websocket.Conn{uuid.New(): websocketConn}

	// Log the updated sessionGroupMap
	wss.logger.Infof("New connection from %s, sessionID: %s, sessionGroupMap: %v", websocketConn.RemoteAddr().String(), sessionID.String(), wss.sessionGroupMap)
}

func (wss *WebsocketServerService) handleDisconnection(websocketConn *websocket.Conn) {
	// Log the disconnection
	wss.logger.Infof("Disconnection from %s closed.", websocketConn.RemoteAddr().String())

	// Iterate through sessionGroupMap to find and remove the disconnected connection
	for sessionID, sessionGroup := range wss.sessionGroupMap {
		for groupID, conn := range sessionGroup {
			if conn == websocketConn {
				delete(sessionGroup, groupID)
				wss.logger.Infof("Disconnection from %s closed, sessionID: %s, groupID: %s", websocketConn.RemoteAddr().String(), sessionID, groupID)
			}
		}

		// If the session is empty after disconnection, remove it from sessionGroupMap
		if len(sessionGroup) == 0 {
			delete(wss.sessionGroupMap, sessionID)
		}
	}

	// Log the updated sessionGroupMap
	wss.logger.Infof("Disconnection from %s closed, sessionGroupMap: %v", websocketConn.RemoteAddr().String(), wss.sessionGroupMap)
}

func (wss *WebsocketServerService) Send(websocketConn *websocket.Conn) {
	wss.logger.Info("Send")

	defer func() {
		if err := websocketConn.Close(); err != nil {
			wss.logger.Error("Failed to close WebSocket connection:", err)
		}
		wss.logger.Info("WebSocket connection closed successfully: %s", websocketConn.RemoteAddr().String())
	}()

	for {
		select {
		case <-wss.ctx.Done():
			return
		default:
			messageType, msg, err := websocketConn.ReadMessage()
			if err != nil {
				wss.logger.Error("Failed to read message:", err)
				return
			}

			wss.logger.Infof("Received Message: %s from %s", msg, websocketConn.RemoteAddr().String())

			var message model.MessageDTO
			if err := json.Unmarshal(msg, &message); err != nil {
				wss.logger.Error("Failed to unmarshal message:", err)
				return
			}

			switch message.Type {
			case types.MESSAGE_TYPES_ECHO:
				wss.logger.Infof("[echo] MessageType: %d", message.Type)
				wss.echo(msg)
			case types.MESSAGE_TYPES_BROADCAST:
				wss.logger.Infof("[broadcast] MessageType: %d", message.Type)
				wss.broadcast(msg)
			case types.MESSAGE_TYPES_UNICAST:
				wss.logger.Infof("[unicast] MessageType: %d", message.Type)
				wss.unicast(msg)
			default:
				wss.logger.Infof("Received message: %s, messageType: %d", msg, messageType)

				// // Echo the message back to the client
				// if err := websocketConn.WriteMessage(messageType, msg); err != nil {
				// 	wss.logger.Error("Failed to write message:", err)
				// 	return
				// }
			}

			// // Echo the message back to the client
			// if err := websocketConn.WriteMessage(messageType, msg); err != nil {
			// 	wss.logger.Error("Failed to write message:", err)
			// 	return
			// }
		}
	}
}

func (wss *WebsocketServerService) broadcast(message []byte) {
	// wss.logger.Infof("broadcast message %s to sessionGroupMap %v", message, wss.sessionGroupMap)
	for _, sessionGroup := range wss.sessionGroupMap {
		// wss.logger.Infof("broadcast message %s to sessionGroup %v", message, sessionGroup)
		for _, conn := range sessionGroup {
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				wss.logger.Error("Failed to write message:", err)
				return
			}
			wss.logger.Infof("Broadcast message: %s to %s", message, conn.RemoteAddr().String())
		}
	}
}

func (wss *WebsocketServerService) echo(message []byte) {
	// wss.logger.Infof("echo message %s to sessionGroupMap %v", message, wss.sessionGroupMap)
	for _, sessionGroup := range wss.sessionGroupMap {
		// wss.logger.Infof("echo message %s to sessionGroup %v", message, sessionGroup)
		for _, conn := range sessionGroup {
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				wss.logger.Error("Failed to write message:", err)
				return
			}
			wss.logger.Infof("Echo message: %s to %s", message, conn.RemoteAddr().String())
		}
	}
}

func (wss *WebsocketServerService) unicast(message []byte) {
	// wss.logger.Infof("unicast message %s to sessionGroupMap %v", message, wss.sessionGroupMap)
	for _, sessionGroup := range wss.sessionGroupMap {
		// wss.logger.Infof("unicast message %s to sessionGroup %v", message, sessionGroup)
		for _, conn := range sessionGroup {
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				wss.logger.Error("Failed to write message:", err)
				return
			}
			wss.logger.Infof("Unicast message: %s to %s", message, conn.RemoteAddr().String())
		}
	}
}

func (wss *WebsocketServerService) HealthHandler(ginContext *gin.Context) {
	ginContext.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}
