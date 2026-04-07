package user

import (
	"inno-accounting/internal/domain"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(*domain.User) (*domain.User, error)
	FindByID(id uuid.UUID) (*domain.User, error)
	FindAll() ([]*domain.User, error)
	DeleteByID(id int) error
	FindBySeveralIDs(ids []uuid.UUID) ([]*domain.User, error)
}