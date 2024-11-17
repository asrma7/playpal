package auth

import (
	"fmt"

	"github.com/asrma7/playpal/api-gateway/config"
	"github.com/asrma7/playpal/auth-svc/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	fmt.Println("Api Gateway: Connecting to Auth Service")
	conn, err := grpc.NewClient(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Error connecting to Auth Service")
	}

	return pb.NewAuthServiceClient(conn)
}
