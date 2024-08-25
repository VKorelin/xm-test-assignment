package authorization

import (
	"context"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var jwtSecret = []byte("qwertyuiopasdfghjklzxcvbnm123456") //hardcoded just for test assignment

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, logger *zap.Logger) (resp interface{}, err error) {
	adminRole, ok := accessabilityRoles[info.FullMethod]
	if !ok {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	headers := md["authorization"]
	if len(headers) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	tokenString := strings.TrimPrefix(headers[0], "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, status.Errorf(codes.Unauthenticated, "expired token")
	}

	role, ok := claims["role"].(string)
	if !ok || role != adminRole {
		return nil, status.Errorf(codes.PermissionDenied, "insufficient permissions")
	} else {
		return handler(ctx, req)
	}
}
