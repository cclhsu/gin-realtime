package route

// import (
//	"fmt"
//	"net/http"

//	"github.com/99designs/gqlgen/graphql/handler"
//	"github.com/99designs/gqlgen/graphql/playground"
//	openapiDocs "github.com/cclhsu/gin-graphql-gorm/doc/openapi"
//	"github.com/cclhsu/gin-graphql-gorm/graph"

//	// "github.com/cclhsu/gin-graphql-gorm/internal/middleware"
//	"github.com/cclhsu/gin-graphql-gorm/internal/service"
//	"github.com/gin-gonic/gin"

//	// "github.com/gin-contrib/cors"
//	swaggerFiles "github.com/swaggo/files"
//	ginSwagger "github.com/swaggo/gin-swagger"
// )

// // Defining the Graphql handler
// func graphqlHandler(resolver graph.Resolver) gin.HandlerFunc {
//	// NewExecutableSchema and Config are in the generated.go file
//	// Resolver is in the resolver.go file
//	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}))
//	return func(c *gin.Context) {
//		h.ServeHTTP(c.Writer, c.Request)
//	}
// }

// // Defining the Playground handler
// func playgroundHandler() gin.HandlerFunc {
//	h := playground.Handler("GraphQL", "/query")
//	return func(c *gin.Context) {
//		h.ServeHTTP(c.Writer, c.Request)
//	}
// }

// // SetupRoutes sets up the API routes
// func SetupGraphQLRoutes(ctx context.Context, r *gin.Engine, host string, port string, logger *logrus.Logger, authService *service.AuthService, userService *service.UserService, teamService *service.TeamService, helloService *service.HelloService, healthService *service.HealthService) {

//	// Create instances of the resolver
//	resolver := graph.Resolver{
//		ctx:	   ctx,
//		// logger:		 logger,
//		AuthService:		authService,
//		UserService:		userService,
//		TeamService:   userService,
//		HelloService:		helloService,
//		HealthService: healthService,
//	}

//	// Enable CORS middleware
//	r.Use(func(c *gin.Context) {
//		origin := fmt.Sprintf("http://%s:%s", host, port)
//		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
//		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

//		// cors.DefaultConfig()
//		// corsConfig := cors.DefaultConfig()
//		// corsConfig.AllowAllOrigins = true
//		// corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
//		// corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
//		// corsConfig.AllowCredentials = true
//		// corsConfig.AddAllowHeaders("Connection")

//		// Handle preflight OPTIONS request
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusNoContent)
//			return
//		}

//		c.Next()
//	})

// // helloGroup := r.Group("/api/v1/hello")
// helloGroup := r.Group("/hello")
// {
//	// Get hello world json
//	helloGroup.GET("/json", helloController.GetHelloJson)

//	// Get hello world string
//	helloGroup.GET("/string", helloController.GetHelloString)
// }

//	// docGroup := r.Group("/api/v1/doc")
//	docGroup := r.Group("/doc")
//	{
//		openapiDocs.SwaggerInfo.BasePath = "/"

//		// Serve Swagger documentation
//		// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//		docGroup.GET("/openapi/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
//	}

//	// queryGroup := r.Group("/graphql")
//	queryGroup := r.Group("/query")
//	// queryGroup.Use(middleware.AuthMiddleware())
//	{
//		queryGroup.POST("", graphqlHandler(resolver))
//	}

//	playgroundGroup := r.Group("/playground")
//	// playgroundGroup.Use(middleware.AuthMiddleware())
//	{
//		playgroundGroup.GET("", playgroundHandler())
//	}

// }
