package services

import (
	"context"
	"fmt"
	"github.com/asrma7/playpal/auth-svc/internal/db"
	"github.com/asrma7/playpal/auth-svc/internal/models"
	"github.com/asrma7/playpal/auth-svc/pkg/pb"
	"github.com/asrma7/playpal/auth-svc/pkg/utils"
	"net/http"
)

type Server struct {
	DBHandler db.DB
	JWT       utils.JWTWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User
	if result := s.DBHandler.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "Email already exists",
		}, nil
	}

	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)
	s.DBHandler.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User
	if result := s.DBHandler.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "Invalid credentials",
		}, nil
	}

	token, _ := s.JWT.GenerateToken(user)
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	claims, err := s.JWT.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateTokenResponse{
			Status: http.StatusBadRequest,
			Error:  fmt.Sprintf("Invalid token: %s", err),
		}, nil
	}

	var user models.User
	if result := s.DBHandler.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateTokenResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateTokenResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}
