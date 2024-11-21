package main

import (
	"fmt"
	"log"
	"net"

	"github.com/asrma7/playpal/feed-svc/config"
	"github.com/asrma7/playpal/feed-svc/internal/db"
	"github.com/asrma7/playpal/feed-svc/internal/services"
	"github.com/asrma7/playpal/feed-svc/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config, %s", err)
	}

	db := db.Init(cfg)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf("Feed Service is running on port: %s\n", cfg.Port)

	s := services.Server{
		DBHandler: db,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterFeedServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
