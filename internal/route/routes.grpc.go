package route

// import (
//	"github.com/cclhsu/gin-realtime/internal/grpc_service_server"
//	"github.com/cclhsu/gin-realtime/internal/service"
//	"google.golang.org/grpc"

//	authGrpcService "github.com/cclhsu/gin-realtime/generated/grpc/pb/auth"
//	healthGrpcService "github.com/cclhsu/gin-realtime/generated/grpc/pb/health"
//	helloGrpcService "github.com/cclhsu/gin-realtime/generated/grpc/pb/hello"
//	teamGrpcService "github.com/cclhsu/gin-realtime/generated/grpc/pb/team"
//	userGrpcService "github.com/cclhsu/gin-realtime/generated/grpc/pb/user"
// )

// func SetupGrpcRoutes(grpcServer *grpc.Server, authService *service.AuthService, userService *service.UserService, teamService *service.TeamService, helloService *service.HelloService, healthService *service.HealthService) {

//	// Register Auth service server
//	authGrpcServiceServer := grpc_service_server.NewAuthServiceServer(authService)
//	authGrpcService.RegisterAuthServiceServer(grpcServer, authGrpcServiceServer)

//	// Register User service server
//	userGrpcServiceServer := grpc_service_server.NewUserServiceServer(userService)
//	userGrpcService.RegisterUserServiceServer(grpcServer, userGrpcServiceServer)

//	// Register team service server
//	teamGrpcServiceServer := grpc_service_server.NewTeamServiceServer(teamService)
//	teamGrpcService.RegisterTeamServiceServer(grpcServer, teamGrpcServiceServer)

//	// Register Hello service server
//	helloGrpcServiceServer := grpc_service_server.NewHelloServiceServer(helloService)
//	helloGrpcService.RegisterHelloServiceServer(grpcServer, helloGrpcServiceServer)

//	// Register Health service server
//	healthGrpcServiceServer := grpc_service_server.NewHealthServiceServer(healthService)
//	healthGrpcService.RegisterHealthServiceServer(grpcServer, healthGrpcServiceServer)
// }
