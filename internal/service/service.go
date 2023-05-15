package service

import (
	"context"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/internal/repository"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/pb"
)

type Authorization interface {
	Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error)
}

type Service struct {
	Authorization
	pb.UnsafeAuthServiceServer
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
	}
}
