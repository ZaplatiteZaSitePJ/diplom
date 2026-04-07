package storages

import (
	"inno-accounting/internal/domain"

	"github.com/google/uuid"
)

type StorageRepository interface {
	Save(*domain.Storage) (*domain.Storage, error)
	FindByID(id uuid.UUID) (*domain.Storage, error)
	FindAll() ([]*domain.Storage, error)
	DeleteByID(id uuid.UUID) error
}