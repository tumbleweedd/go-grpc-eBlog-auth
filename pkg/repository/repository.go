package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/models"
)

type Authorization interface {
	Register(user models.User) error
	Login(username, password string) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
