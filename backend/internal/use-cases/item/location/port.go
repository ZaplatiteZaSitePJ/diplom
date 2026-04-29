package location

import (
	"inno-accounting/internal/domain"

	"github.com/google/uuid"
)

type LocationRepository interface {
	GetByItemID(itemID uuid.UUID) (*domain.ItemLocationDetails, error)
	Upsert(loc *domain.ItemLocation) error
}