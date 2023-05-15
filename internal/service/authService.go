package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/internal/models"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/internal/repository"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/pkg/pb"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/pkg/utils"
	"net/http"
	"time"
)

const (
	salt       = "qelwnjgo23ijqpk1"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (authService *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	user.Name = req.Name
	user.Lastname = req.Username
	user.Username = req.Username
	user.Email = req.Email
	user.Password = generateHashPassword(req.Password)

	err := authService.repo.Register(user)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusOK,
	}, nil
}

func (authService *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := authService.GenerateToken(req.Username, req.Password)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (authService *AuthService) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	userId, err := ParseToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: int64(userId),
	}, nil
}

func (authService *AuthService) GenerateToken(username, password string) (string, error) {

	user, err := authService.repo.Login(username, generateHashPassword(password))
	fmt.Println(user.Id, ":", user.Email)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &utils.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: int(user.Id),
	})

	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &utils.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
		}
		return []byte("qrkjk#4#%35FSFJlja#4353KSFjH"), nil
	})

	if claims, ok := token.Claims.(*utils.TokenClaims); ok && token.Valid {
		return claims.UserId, nil
	}

	return 0, err
}

func generateHashPassword(password string) string {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(salt))
	return fmt.Sprintf("%x", pwd.Sum(nil))
}
