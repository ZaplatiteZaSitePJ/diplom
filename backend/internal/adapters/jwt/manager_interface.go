package jwt

import (
	jwt_config "inno-accounting/pkg/jwt"

	"github.com/google/uuid"
)

type AuthData struct {
	UserID uuid.UUID
	Role   string
}

type TokenManager interface {
	GenerateTokens(userID uuid.UUID, role string) (access, refresh string, err error)
	ValidateAccess(token string) (AuthData, error)
	ValidateRefresh(token string) (uuid.UUID, error)
}

type JWTTokenManager struct {
	manager *jwt_config.JWTConfig
}

func NewJWTTokenManager(manager *jwt_config.JWTConfig) *JWTTokenManager {
	return &JWTTokenManager{
		manager: manager,
	}
}

func (m *JWTTokenManager) GenerateTokens(userID uuid.UUID, role string) (access, refresh string, err error) {
	return m.manager.GenerateTokens(userID, role)
}

func (m *JWTTokenManager) ValidateAccess(token string) (AuthData, error) {
	claims, err := m.manager.ValidateAccess(token)
	if err != nil {
		return AuthData{}, err
	}

	id, err := uuid.Parse(claims.UserID)
	if err != nil {
		return AuthData{}, err
	}

	return AuthData{
		UserID: id,
		Role:   claims.Role,
	}, nil
}

func (m *JWTTokenManager) ValidateRefresh(token string) (uuid.UUID, error) {
	claims, err := m.manager.ValidateRefresh(token)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(claims.UserID)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}