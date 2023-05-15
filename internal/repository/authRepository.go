package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-auth/internal/models"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (authRepo *AuthRepository) Register(user models.User) error {
	query := fmt.Sprintf(`
								INSERT INTO %s (email, is_account_non_locked, lastname, name, password, role, username) values ($1, $2, $3, $4, $5, $6, $7)
                                RETURNING user_id`, usersTable)
	_, err := authRepo.db.Exec(query, user.Email, models.IsAccountNonLocked, user.Lastname, user.Name, user.Password, models.USER, user.Username)

	if err != nil {
		return err
	}
	return nil
}

func (authRepo *AuthRepository) Login(username, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf(`SELECT user_id FROM %s u WHERE u.username=$1 AND u.password=$2`, usersTable)
	err := authRepo.db.Get(&user, query, username, password)

	return user, err
}
