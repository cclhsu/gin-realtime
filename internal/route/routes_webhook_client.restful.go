package route

import (
	"context"
	"fmt"
	"net/http"

	openapiDocs "github.com/cclhsu/gin-realtime/doc/openapi"
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
func SetupRestfulWebhookClientRoutes(ctx context.Context, r *gin.Engine, host string, port string, logger *logrus.Logger, helloService *service.HelloService, healthService *service.HealthService, webhookClientService *service.WebhookClientService) {

	// Create instances of the controller
	helloController := controller.NewHelloController(ctx, logger, helloService)
	healthController := controller.NewHealthController(ctx, logger, healthService)
	webhookClientController := controller.NewWebhookClientController(ctx, logger, webhookClientService)

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
		openapiDocs.SwaggerInfo.BasePath = "/"
		// openapiDocs.SwaggerInfowebsocket_client_service.BasePath = "/"

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

	// handle webhook response from server
	// webhookClientGroup := r.Group("/api/v1/webhook-client")
	webhookClientGroup := r.Group("/webhook-client")
	{
		// webhookGroup.Use(middleware.WebhookMiddleware())

		// Register a webhook
		webhookClientGroup.POST("/register", webhookClientController.Register)

		// Unregister a webhook
		webhookClientGroup.DELETE("/register/:webhookId", webhookClientController.Unregister)

		// List all webhooks
		webhookClientGroup.GET("/register", webhookClientController.ListRegistrations)

		// Update a webhook
		webhookClientGroup.PUT("/register/:webhookId", webhookClientController.UpdateRegistration)

		// Send a Message to a webhook
		webhookClientGroup.POST("/message/send", webhookClientController.Send)

		// Receive a Message from a webhook
		webhookClientGroup.POST("/message/receive", webhookClientController.Receive)

		// List all Messages
		webhookClientGroup.GET("/message", webhookClientController.ListMessages)

		// Health check
		// webhookClientGroup.GET("/health", webhookClientController.Health)
	}
}
