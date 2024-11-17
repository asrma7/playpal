package routes

import (
	"context"
	"net/http"

	"github.com/asrma7/playpal/feed-svc/pkg/pb"
	"github.com/gin-gonic/gin"
)

func FindAll(ctx *gin.Context, conn pb.FeedServiceClient) {
	res, err := conn.FindAll(context.Background(), &pb.FindAllRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
