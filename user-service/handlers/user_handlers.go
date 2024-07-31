package handlers

import (
	"context"
	"log"
	pb "user-service/proto"
	"user-service/services"
)

type UserHandlers struct {
	service *services.UserService
}

func (h *UserHandlers) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Register user: %v", req.Username)
	return h.service.Register(ctx, req)
}

func (h *UserHandlers) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Login user: %v", req.Email)
	return h.service.Login(ctx, req)
}

func (h *UserHandlers) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	log.Printf("Get user profile: %v", req.UserId)
	return h.service.GetUserProfile(ctx, req)
}
