package services

import (
    "context"
    "user-service/models"
    "user-service/repository"
    pb "user-service/proto"
    "user-service/utils"

    "github.com/go-redis/redis/v8"
    "github.com/streadway/amqp"
)

type UserService struct {
    repo  *repository.UserRepository
    redis *redis.Client
    mq    *amqp.Connection
    pb.UnimplementedUserServiceServer
}

func NewUserService(repo *repository.UserRepository, redis *redis.Client, mq *amqp.Connection) *UserService {
    return &UserService{
        repo:  repo,
        redis: redis,
        mq:    mq,
    }
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
    hashedPassword, err := utils.HashPassword(req.Password)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Username: req.Username,
        Email:    req.Email,
        Password: hashedPassword,
    }

    err = s.repo.CreateUser(ctx, user)
    if err != nil {
        return nil, err
    }

    return &pb.RegisterResponse{UserId: user.UserID, Message: "User registered successfully"}, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
    user, err := s.repo.GetUserByEmail(ctx, req.Email)
    if err != nil {
        return nil, err
    }

    if !utils.CheckPasswordHash(req.Password, user.Password) {
        return nil, err
    }

    return &pb.LoginResponse{UserId: user.UserID, Message: "User logged in successfully"}, nil
}

func (s *UserService) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
    user, err := s.repo.GetUserByID(ctx, req.UserId)
    if err != nil {
        return nil, err
    }

    return &pb.GetUserProfileResponse{UserId: user.UserID, Username: user.Username, Email: user.Email}, nil
}
