package auth

import (
	"inno-accounting/pkg/crypt_password"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"

	"github.com/google/uuid"
)

type TokenManager interface {
	GenerateTokens(userID uuid.UUID, role string) (string, string, error)
}

func New(repo AuthRepository, tokens TokenManager) *AuthService {
	return &AuthService{
		repo:   repo,
		tokens: tokens,
	}
}

type AuthService struct {
	repo AuthRepository
	tokens TokenManager
}

func (a *AuthService) Login(email, password string) (string, string, error) {
	user, err := a.repo.FindByEmail(email)
	if err != nil {
		return "", "", app_errors.Unprocessable("No user match this email", err)
	}

	// проверка пароля
	err = crypt_password.CompareWithHash(user.HashedPassword, password)
	if err != nil {
		logger.Info("service", err)
		return "", "", app_errors.Unprocessable("Wrong password", err)
	}
	

	return a.tokens.GenerateTokens(user.ID, user.Role)
}