package repositories

import (
	"database/sql"
	"inno-accounting/internal/domain"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

// FIND USER BY EMAIL IN DATABASE
func (authRepo *AuthRepository) FindByEmail(email string) (*domain.User, error) {
	findedUser := &domain.User{}
	
	userQuery := "SELECT id, email, hashed_password, role FROM users WHERE email = $1"

	err := authRepo.db.QueryRow(userQuery, email).Scan(&findedUser.ID, &findedUser.Email, &findedUser.HashedPassword, &findedUser.Role)

	if err != nil {
		logger.Error("db", err)
		if err == sql.ErrNoRows {
			return nil, app_errors.NotFound("user not found", err)
		}
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	return findedUser, nil
}