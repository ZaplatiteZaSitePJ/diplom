package auth

import (
	"inno-accounting/internal/domain"
	"time"

	"github.com/google/uuid"
)

type AuthRepository interface {
	FindByEmail(email string) (*domain.User, error)
	SaveRefreshToken(userID uuid.UUID, refresh string, expireAt time.Time) error
	FindRefreshToken(refresh string) (uuid.UUID, error)
	DeactivateRefreshToken(refresh string) error
}