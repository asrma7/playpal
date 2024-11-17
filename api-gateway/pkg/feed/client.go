package feed

import (
	"fmt"

	"github.com/asrma7/playpal/api-gateway/config"
	"github.com/asrma7/playpal/feed-svc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.FeedServiceClient
}

func InitServiceClient(c *config.Config) pb.FeedServiceClient {
	fmt.Println("Api Gateway: Connecting to Feed Service")
	conn, err := grpc.NewClient(c.FeedSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Error connecting to Feed Service")
	}

	return pb.NewFeedServiceClient(conn)
}
