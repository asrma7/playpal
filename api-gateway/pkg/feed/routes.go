package feed

import (
	"fmt"

	"github.com/asrma7/playpal/api-gateway/config"
	"github.com/asrma7/playpal/api-gateway/pkg/auth"
	"github.com/asrma7/playpal/api-gateway/pkg/feed/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authsvc *auth.ServiceClient) {
	fmt.Println("Api Gateway: Registering routes")
	auth := auth.InitAuthMiddleware(authsvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	feed := r.Group("/feed")
	feed.Use(auth.UserAuth)
	feed.POST("/", svc.Create)
	feed.GET("/:id", svc.FindOne)
	feed.GET("/", svc.FindAll)

}

func (svc *ServiceClient) Create(ctx *gin.Context) {
	routes.CreateFeed(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) FindAll(ctx *gin.Context) {
	routes.FindAll(ctx, svc.Client)
}
