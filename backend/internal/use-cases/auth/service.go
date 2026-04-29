package auth

import (
	"inno-accounting/internal/use-cases/user"
	"inno-accounting/pkg/crypt_password"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	"time"

	"github.com/google/uuid"
)

type TokenManager interface {
	GenerateTokens(userID uuid.UUID, role string) (string, string, error)
	ValidateRefresh(token string) (uuid.UUID, error)
}

func New(repo AuthRepository, userService *user.UserService, tokens TokenManager) *AuthService {
	return &AuthService{
		repo:   repo,
		user: userService,
		tokens: tokens,
	}
}

type AuthService struct {
	repo AuthRepository
	user *user.UserService
	tokens TokenManager
}

func (a *AuthService) Login(email, password string) (string, string, error) {
	user, err := a.repo.FindByEmail(email)
	if err != nil {
		return "", "", app_errors.Unprocessable("No user match this email", err)
	}

	err = crypt_password.CompareWithHash(user.HashedPassword, password)
	if err != nil {
		logger.Info("service", err)
		return "", "", app_errors.Unprocessable("Wrong password", err)
	}

	access, refresh, err := a.tokens.GenerateTokens(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	expireAt := time.Now().Add(7 * 24 * time.Hour)

	err = a.repo.SaveRefreshToken(user.ID, refresh, expireAt)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (a *AuthService) Refresh(refreshToken string) (string, string, error) {
	// проверяем токен в БД
	userID, err := a.repo.FindRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	// валидируем JWT
	parsedUserID, err := a.tokens.ValidateRefresh(refreshToken)
	if err != nil {
		return "", "", app_errors.Unprocessable("invalid refresh token", err)
	}

	if userID != parsedUserID {
		return "", "", app_errors.Unprocessable("token mismatch", nil)
	}

	// инвалидируем старый refresh
	err = a.repo.DeactivateRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	user, err := a.user.FindUserByID(userID)
	if err != nil {
		return "", "", err
	}

	// генерим новые токены
	access, newRefresh, err := a.tokens.GenerateTokens(userID, user.Role)
	if err != nil {
		return "", "", err
	}

	expireAt := time.Now().Add(7 * 24 * time.Hour)

	// сохраняем новый refresh
	err = a.repo.SaveRefreshToken(userID, newRefresh, expireAt)
	if err != nil {
		return "", "", err
	}

	return access, newRefresh, nil
}

func (a *AuthService) Logout(refreshToken string) error {
	err := a.repo.DeactivateRefreshToken(refreshToken)
	if err != nil {
		return err
	}

	return nil
}