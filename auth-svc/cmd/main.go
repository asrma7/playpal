package main

import (
	"fmt"
	"log"
	"net"

	"github.com/asrma7/playpal/auth-svc/config"
	"github.com/asrma7/playpal/auth-svc/internal/db"
	"github.com/asrma7/playpal/auth-svc/internal/services"
	"github.com/asrma7/playpal/auth-svc/pkg/pb"
	"github.com/asrma7/playpal/auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config, %s", err)
	}

	db := db.Init(cfg)

	jwt := utils.JWTWrapper{
		SecretKey:       cfg.JWTSecretKey,
		Issuer:          "playpal-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf("Auth Service is running on port: %s\n", cfg.Port)

	s := services.Server{
		DBHandler: db,
		JWT:       jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
