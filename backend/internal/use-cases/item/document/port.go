package document

import (
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"

	"github.com/google/uuid"
)

type DocumentRepository interface {
	Save(*domain.Document) (*domain.Document, error)
	Change(uuid.UUID, *domain.Document) (*domain.Document, error)
	FindByID(id uuid.UUID) (*domain.Document, error)
	FindAll(filter *dto.DocsFilter) ([]*dto.DocsItemPublic, error)
	FindCategoryIDByName(string) (int, error)
	FindCategoryNameByID(int) (*string, error)
}