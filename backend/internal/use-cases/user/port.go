package user

import (
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	FindAll(filter *dto.UserFilter) ([]*domain.User, error)
	FindByID(id uuid.UUID) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(id uuid.UUID) error
	FindByEmail(email string) (*domain.User, error)
}