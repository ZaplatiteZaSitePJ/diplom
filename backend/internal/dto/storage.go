package dto

import (
	"inno-accounting/internal/domain"

	"github.com/google/uuid"
)

type CreateStorage struct {
	StorageName   string     `db:"storage_name" json:"storageName"`
	Capacity      int        `db:"capacity" json:"capacity"`
	City          string     `db:"city" json:"city"`
}

type PublicStorage struct {
	ID            uuid.UUID  `db:"id" json:"id"`
	StorageName   string     `db:"storage_name" json:"storageName"`
	Capacity      int        `db:"capacity" json:"capacity"`
	OccupiedCells int        `db:"occupied_cells" json:"occupied_cells"`
	City          string     `db:"city" json:"city"`
	ItemsAmount   int        `db:"items_amount" json:"items_amount"`
}

func PublicStorageFromModel(storage *domain.Storage) *PublicStorage {
	return &PublicStorage{
		ID: storage.ID, 
		StorageName: storage.StorageName,
		Capacity: storage.Capacity,
		OccupiedCells: storage.OccupiedCells,
		City: storage.City,
		ItemsAmount: storage.ItemsAmount,
	}
}

func SeveralStoragesToPublic(storages []*domain.Storage) []*PublicStorage {
    publicStorages := make([]*PublicStorage, 0, len(storages))
    for _, s := range storages {
        publicStorages = append(publicStorages, PublicStorageFromModel(s))
    }
    return publicStorages
}