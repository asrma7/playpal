package routes

import (
	"context"
	"net/http"

	"github.com/asrma7/playpal/feed-svc/pkg/pb"
	"github.com/gin-gonic/gin"
)

type CreateFeedRequest struct {
	Title    string `json:"content"`
	Author   string `json:"author"`
	Datetime string `json:"datetime"`
	Players  int64  `json:"players"`
	Location string `json:"location"`
}

func CreateFeed(ctx *gin.Context, conn pb.FeedServiceClient) {
	body := CreateFeedRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := conn.CreateFeed(context.Background(), &pb.CreateFeedRequest{
		Title:    body.Title,
		Author:   body.Author,
		Datetime: body.Datetime,
		Players:  body.Players,
		Location: body.Location,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
