package service

import (
	"context"
	"net/http"

	// "net/url"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	// "github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type WebsocketServerInterface interface {
	WebsocketHandler(ginContext *gin.Context)
	HealthHandler(ginContext *gin.Context)
	handleConnection(websocketConn *websocket.Conn)
	handleDisconnection(websocketConn *websocket.Conn)
	Echo(websocketConn *websocket.Conn)
	Broadcast(websocketConn *websocket.Conn)
	handleMessage(websocketConn *websocket.Conn, messageType int, message []byte)
	handleEvent(websocketConn *websocket.Conn, event string, message []byte)
}

type WebsocketServerService struct {
	ctx             context.Context
	logger          *logrus.Logger
	sessionGroupMap map[string]map[uuid.UUID]*websocket.Conn
	upgrader        *websocket.Upgrader
}

func NewWebsocketServerService(ctx context.Context, logger *logrus.Logger) *WebsocketServerService {
	// SERVICE_HOST := os.Getenv("SERVICE_HOST")
	// if SERVICE_HOST == "" {
	// 	SERVICE_HOST = "localhost"
	// }
	// SERVICE_PORT := os.Getenv("SERVICE_PORT")
	// if SERVICE_PORT == "" {
	// 	SERVICE_PORT = "8080"
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
				// 	return false
				// }
				// return u.Host == SERVICE_HOST+":"+SERVICE_PORT\
				return true
			},
		},
	}
}

func (s *WebsocketServerService) Initialize() {
	s.logger.Infof("WebsocketServerService Initialize")
}

func (s *WebsocketServerService) WebsocketHandler(ginContext *gin.Context) {
	websocketConn, err := s.upgrader.Upgrade(ginContext.Writer, ginContext.Request, nil)
	if err != nil {
		s.logger.Error("Failed to upgrade connection:", err)
		return
	}
	defer websocketConn.Close()

	// Handle connection
	s.handleConnection(websocketConn)
	// s.logger.Infof("New connection from %s", websocketConn.RemoteAddr().String())

	// Handle disconnection
	defer s.handleDisconnection(websocketConn)
	// defer func() {
	// 	s.logger.Infof("Connection from %s closed.", websocketConn.RemoteAddr().String())
	// }()

	// Echo messages
	s.Echo(websocketConn)

	// Broadcast messages
	s.Broadcast(websocketConn)

	// Handle messages
	// s.handleMessage(websocketConn)

	// Handle events
	// s.handleEvent(websocketConn, "event", []byte("message"))
}

func (s *WebsocketServerService) HealthHandler(ginContext *gin.Context) {
	ginContext.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func (s *WebsocketServerService) handleConnection(websocketConn *websocket.Conn) {
	// Implement your connection logic here
	// s.logger.Infof("New connection from %s", websocketConn.RemoteAddr().String())
	sessionID := uuid.New()
	s.logger.Infof("New connection from %s, sessionID: %s", websocketConn.RemoteAddr().String(), sessionID.String())
	s.sessionGroupMap[sessionID.String()] = make(map[uuid.UUID]*websocket.Conn)
	s.sessionGroupMap[sessionID.String()][uuid.New()] = websocketConn
	s.logger.Infof("New connection from %s, sessionID: %s, sessionGroupMap: %v", websocketConn.RemoteAddr().String(), sessionID.String(), s.sessionGroupMap)
}

func (s *WebsocketServerService) handleDisconnection(websocketConn *websocket.Conn) {
	// Implement your disconnection logic here
	s.logger.Infof("Connection from %s closed.", websocketConn.RemoteAddr().String())

	for sessionID, sessionGroup := range s.sessionGroupMap {
		for groupID, conn := range sessionGroup {
			if conn == websocketConn {
				delete(sessionGroup, groupID)
				s.logger.Infof("Connection from %s closed, sessionID: %s, groupID: %s", websocketConn.RemoteAddr().String(), sessionID, groupID)
			}
		}
	}
	s.logger.Infof("Connection from %s closed, sessionGroupMap: %v", websocketConn.RemoteAddr().String(), s.sessionGroupMap)
}

func (s *WebsocketServerService) Echo(websocketConn *websocket.Conn) {
	s.logger.Infof("Echo")
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			messageType, message, err := websocketConn.ReadMessage()
			if err != nil {
				s.logger.Error("Failed to read message:", err)
				return
			}

			s.logger.Infof("Received message: %s, messageType: %d", message, messageType)

			// Echo the message back to the client
			if err := websocketConn.WriteMessage(messageType, message); err != nil {
				s.logger.Error("Failed to write message:", err)
				return
			}
			s.logger.Infof("Echo message: %s to %s", message, websocketConn.RemoteAddr().String())
		}
	}
}

func (s *WebsocketServerService) Broadcast(websocketConn *websocket.Conn) {
	s.logger.Infof("Broadcast")
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			messageType, message, err := websocketConn.ReadMessage()
			if err != nil {
				s.logger.Error("Failed to read message:", err)
				return
			}

			s.logger.Infof("Received message: %s, messageType: %d", message, messageType)

			// Broadcast the message to all the clients
			for _, sessionGroup := range s.sessionGroupMap {
				for _, conn := range sessionGroup {
					if err := conn.WriteMessage(messageType, message); err != nil {
						s.logger.Error("Failed to write message:", err)
						return
					}
					s.logger.Infof("Broadcast message: %s to %s", message, conn.RemoteAddr().String())
				}
			}
		}
	}
}

func (s *WebsocketServerService) handleMessage(websocketConn *websocket.Conn) {
	s.logger.Infof("handleMessage")
	for {
		select {
		case <-s.ctx.Done():
			return
		// case <-websocketConn.CloseHandler():
		// 	return
		// case <-websocketConn.PingHandler():
		// 	return
		default:
			messageType, message, err := websocketConn.ReadMessage()
			if err != nil {
				s.logger.Error("Failed to read message:", err)
				return
			}

			// Handle the message based on the messageType (text, binary, etginContext.)
			switch messageType {
			case websocket.TextMessage:
				s.logger.Infof("Received message: %s, messageType: %d", message, messageType)
				s.Echo(websocketConn)
			}
		}
	}
}

func (s *WebsocketServerService) handleEvent(websocketConn *websocket.Conn, event string, message []byte) {
	s.logger.Infof("Received event: %s, message: %s", event, message)
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			messageType, message, err := websocketConn.ReadMessage()
			if err != nil {
				s.logger.Error("Failed to read message:", err)
				return
			}

			// Handle the message based on the messageType (text, binary, etginContext.)
			switch messageType {
			case websocket.TextMessage:
				s.logger.Infof("Received message: %s, messageType: %d", message, messageType)
				s.Echo(websocketConn)
			}
		}
	}
}
