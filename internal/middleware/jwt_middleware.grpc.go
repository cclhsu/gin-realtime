package middleware

// import (
//	"context"
//	"os"
//	"strings"

//	"github.com/cclhsu/gin-realtime/internal/model"
//	"github.com/golang-jwt/jwt"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/metadata"
//	"google.golang.org/grpc/status"
// )

// // AuthInterceptor is a gRPC interceptor to verify the JWT token
// func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//	// Retrieve the token from the gRPC metadata
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		return nil, status.Errorf(codes.Unauthenticated, "Missing metadata")
//	}

//	authHeaders := md.Get("authorization")
//	if len(authHeaders) == 0 {
//		return nil, status.Errorf(codes.Unauthenticated, "Missing authorization header")
//	}

//	// Extract the token from the header
//	tokenParts := strings.Split(authHeaders[0], " ")
//	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
//		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
//	}
//	tokenString := tokenParts[1]

//	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

//	// Parse the token
//	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return jwtSecret, nil
//	})

//	// Verify the token
//	if err != nil {
//		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
//	}

//	// Check if the token is valid
//	if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
//		// Set the ID in the context
//		ctx = context.WithValue(ctx, "ID", claims.ID)
//		ctx = context.WithValue(ctx, "UUID", claims.Sub)
//		return handler(ctx, req)
//	}

//	return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
// }
