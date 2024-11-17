package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/asrma7/playpal/feed-svc/pkg/pb"
	"github.com/gin-gonic/gin"
)

func FindOne(ctx *gin.Context, conn pb.FeedServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := conn.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
