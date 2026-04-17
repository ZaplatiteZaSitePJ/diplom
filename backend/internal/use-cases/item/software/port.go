package software

import (
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"

	"github.com/google/uuid"
)

type SoftwareRepository interface {
	Save(*domain.Software) (*domain.Software, error)
	Change(uuid.UUID, *domain.Software) (*domain.Software, error)
	FindByID(id uuid.UUID) (*domain.Software, error)
	FindAll(filter *dto.SoftwareFilter) ([]*dto.SoftwareItemPublic, error)

	FindCategoryIDByName(string) (int, error)
	FindCategoryNameByID(int) (*string, error)
	GetCategoriesByTypeID(typeID int) ([]string, error)
}