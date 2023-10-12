package main

// go run ./cmd/websocket-client-service/main.go
// swag init -g cmd/websocket-client-service/main.go -o doc/openapi
// go build -o ./bin/websocket-client-service ./cmd/websocket-client-service
// ./bin/websocket-client-service

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/cclhsu/gin-realtime/internal/config"
	"github.com/cclhsu/gin-realtime/internal/route"
	"github.com/cclhsu/gin-realtime/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	ctx    context.Context
	logger *logrus.Logger

	host     string
	port     string
	endpoint string
	router   *gin.Engine

	helloService  *service.HelloService
	healthService *service.HealthService
)

// CallerPrettyfier is a function that formats the caller information.
func CallerPrettyfier(f *runtime.Frame) (string, string) {
	// Split the file path by slashes
	parts := strings.Split(f.File, "/")
	// Get the last part of the split, which is the file name
	fileName := parts[len(parts)-1]
	// Format the caller information as "filename:line"
	return fmt.Sprintf("%5d ", f.Line), fmt.Sprintf("%20v ", fileName)
}

func setupLogger() {
	// Create a new logrus logger
	logger = logrus.New()
	logger.SetOutput(os.Stdout)

	// // Create a new loggly hook
	// hook := logrusly.NewLogglyHook("https://logs-01.loggly.com/inputs/0b0f0f1e-0b0b-4b0b-8b0b-0b0b0b0b0b0b/tag/http/", "gin-restful-gorm", logrus.InfoLevel)

	// // Add the hook to the logger
	// logger.Hooks.Add(hook)

	// Set the logger formatter
	logger.SetReportCaller(true)

	// // Set the logger formatter
	// logger.SetFormatter(&logrus.JSONFormatter{})
	// logrus.SetFormatter(&logrus.JSONFormatter{
	//	DisableTimestamp: false,
	//	TimestampFormat:  "2006-01-02 15:04:05",
	//	FieldMap: logrus.FieldMap{
	//		logrus.FieldKeyTime:  "@timestamp",
	//		logrus.FieldKeyLevel: "@level",
	//		logrus.FieldKeyMsg:	  "@message",
	//	}})
	// logger.SetFormatter(&logrus.TextFormatter{
	//	FullTimestamp:			true,
	//	TimestampFormat:		"2006-01-02 15:04:05",
	//	DisableLevelTruncation: true,
	//	CallerPrettyfier:		CallerPrettyfier,
	//	// You can customize other formatting options here
	// })
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:    false, // Disable colored output
		FullTimestamp:    true,  // Include the timestamp
		TimestampFormat:  time.RFC3339,
		CallerPrettyfier: CallerPrettyfier,
	})

	// Set the logger level: info, debug, warn, error, fatal
	logger.SetLevel(logrus.TraceLevel)

	// Log a message
	logger.Info("Hello world!")
}

func startGinServer() {
	host = os.Getenv("SERVICE_HOST")
	port = os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8080"
	}
	endpoint = host + ":" + port

	// Set Gin to production mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Set up Gin server
	router := gin.Default()

	route.SetupRestfulWebsocketClientRoutes(router, host, port, logger, helloService, healthService)
	// route.SetupGraphQLRoutes(router, host, port, logger, authService, userService, teamService, helloService, healthService)

	// // Add redis client to gin context
	// router.Use(func(c *gin.Context) {
	// 	c.Set("redis", redisClient)
	// 	c.Next()
	// })

	fmt.Printf("Starting up on http://%s/\n", endpoint)
	fmt.Printf("Starting up on http://%s/doc/openapi/index.html\n", endpoint)
	// Start the Gin server
	err := router.Run(endpoint)
	if err != nil {
		logger.Fatalf("Failed to start the server: %v", err)
	}
}
func main() {
	// Create a context
	ctx = context.Background()

	// Set up logger
	setupLogger()

	// Load environment variables
	err := config.LoadEnv()
	if err != nil {
		logger.Fatal("Failed to load environment variables")
	}

	// Create the hello service
	helloService = service.NewHelloService(ctx, logger)

	// Create the health check service
	healthService = service.NewHealthService(ctx, logger)

	// Start Gin server in a goroutine
	go startGinServer()

	// Block the main goroutine to keep the servers running
	select {}
}
