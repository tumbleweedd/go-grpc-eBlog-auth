package service

import (
	"context"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/internal/repository"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/pkg/pb"
)

type Authorization interface {
	Register(ctx context.Context, req *pb2.RegisterRequest) (*pb2.RegisterResponse, error)
	Login(ctx context.Context, req *pb2.LoginRequest) (*pb2.LoginResponse, error)
	Validate(ctx context.Context, req *pb2.ValidateRequest) (*pb2.ValidateResponse, error)
}

type Service struct {
	Authorization
	pb2.UnsafeAuthServiceServer
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
	}
}
