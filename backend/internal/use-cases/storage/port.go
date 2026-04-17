package storages

import (
	"inno-accounting/internal/domain"

	"github.com/google/uuid"
)

type StorageRepository interface {
	Save(*domain.Storage) (*domain.Storage, error)
	Change(uuid.UUID, *domain.Storage) (*domain.Storage, error)
	FindByID(id uuid.UUID) (*domain.Storage, error)
	FindByExactName(string) (*domain.Storage, error)
	FindAll() ([]*domain.Storage, error)
	GetStorageStats(storageID uuid.UUID) (int, int, error)
	TransferAndDelete(oldStorageID, newStorageID uuid.UUID) error
	DeleteWithItems(storageID uuid.UUID) error
}