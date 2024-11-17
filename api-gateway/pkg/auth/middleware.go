package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/asrma7/playpal/auth-svc/pkg/pb"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	fmt.Println("Api Gateway: Initializing Auth Middleware")
	return AuthMiddlewareConfig{svc: svc}
}

func (c *AuthMiddlewareConfig) UserAuth(ctx *gin.Context) {
	fmt.Println("Api Gateway: User Auth Middleware")
	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.ValidateToken(context.Background(), &pb.ValidateTokenRequest{
		Token: token[1],
		Role:  "user",
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}
