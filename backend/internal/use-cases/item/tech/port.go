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
	FindAll(filter *dto.TechFilter) ([]*dto.TechItemPublic, error)
	DeleteByID(id uuid.UUID) error
	FindCategoryIDByName(string) (int, error)
	FindCategoryNameByID(int) (*string, error)
	GetCategoriesByTypeID(typeID int) ([]string, error)
}