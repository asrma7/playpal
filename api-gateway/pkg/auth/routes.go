package auth

import (
	"fmt"

	"github.com/asrma7/playpal/api-gateway/config"
	"github.com/asrma7/playpal/api-gateway/pkg/auth/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	fmt.Println("Api Gateway: Registering routes")
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	user := r.Group("/auth")
	user.POST("/register", svc.Register)
	user.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
