package route

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cclhsu/gin-realtime/internal/controller"
	"github.com/sirupsen/logrus"

	// "github.com/cclhsu/gin-realtime/internal/middleware"
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"

	// "github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes sets up the API routes
func SetupRestfulSocketIOServerRoutes(ctx context.Context, r *gin.Engine, host string, port string, logger *logrus.Logger, helloService *service.HelloService, healthService *service.HealthService, socketIOServerService *service.SocketIOServerService) {

	// Create instances of the controller
	helloController := controller.NewHelloController(ctx, logger, helloService)
	healthController := controller.NewHealthController(ctx, logger, healthService)
	socketIOServerController := controller.NewSocketIOServerController(ctx, logger, socketIOServerService)

	// Enable CORS middleware
	r.Use(func(c *gin.Context) {
		origin := fmt.Sprintf("http://%s:%s", host, port)
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// cors.DefaultConfig()
		// corsConfig := cors.DefaultConfig()
		// corsConfig.AllowAllOrigins = true
		// corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
		// corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
		// corsConfig.AllowCredentials = true
		// corsConfig.AddAllowHeaders("Connection")

		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// helloGroup := r.Group("/api/v1/hello")
	helloGroup := r.Group("/hello")
	{
		// Get hello world json
		helloGroup.GET("/json", helloController.GetHelloJson)

		// Get hello world string
		helloGroup.GET("/string", helloController.GetHelloString)
	}

	// docGroup := r.Group("/api/v1/doc")
	docGroup := r.Group("/doc")
	{
		// openapiDocs.SwaggerInfo.BasePath = "/"
		// openapiDocs.SwaggerInfosocketIO_server_service.BasePath = "/"

		// Serve Swagger documentation
		// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		docGroup.GET("/openapi/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// healthGroup := r.Group("/api/v1/health")
	healthGroup := r.Group("/health")
	{
		// Get health check
		healthGroup.GET("/healthy", healthController.IsHealthy)

		// Get health check
		healthGroup.GET("/live", healthController.IsALive)

		// Get health check
		healthGroup.GET("/ready", healthController.IsReady)
	}

	socketIOGroup := r.Group("/socket-io")
	{
		// Get socketIO handler
		socketIOGroup.GET("/handler", socketIOServerController.SocketIOHandler)
		// socketIOGroup.GET("/echo", socketIOServerController.SocketIOEchoHandler)
		// socketIOGroup.GET("/broadcast", socketIOServerController.SocketIOBroadcastHandler)
		// socketIOGroup.GET("/event", socketIOServerController.SocketIOEventHandler)
		// socketIOGroup.GET("/message", socketIOServerController.SocketIOMessageHandler)
		// socketIOGroup.GET("/health", socketIOServerController.SocketIO HealthHandler)
	}
}
