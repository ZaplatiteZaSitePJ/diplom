package domain

import (
	"time"

	"github.com/google/uuid"
)

type Storage struct {
	ID            uuid.UUID  `db:"id" json:"id"`
	StorageName   string     `db:"storage_name" json:"storageName"`
	Capacity      int        `db:"capacity" json:"capacity"`
	OccupiedCells int        `db:"occupied_cells" json:"occupied_cells"`
	City          string     `db:"city" json:"city"`
	ItemsAmount   int        `db:"items_amount" json:"items_amount"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt     *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}