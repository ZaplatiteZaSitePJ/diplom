package tech

import (
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"

	"github.com/google/uuid"
)

type TechRepository interface {
	Save(*domain.Tech) (*domain.Tech, error)
	Change(uuid.UUID, *domain.Tech) (*domain.Tech, error)
	FindByID(id uuid.UUID) (*domain.Tech, error)
	FindAll(filter *dto.TechFilter) ([]*domain.Tech, error)
	DeleteByID(id uuid.UUID) error
	FindCategoryIDByName(string) (int, error)
}