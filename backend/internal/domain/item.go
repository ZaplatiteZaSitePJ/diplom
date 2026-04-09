package domain

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID             uuid.UUID `db:"id" json:"id"`
	UniversalName  string    `db:"universal_name" json:"universal_name"`
	TypeID         int       `db:"type_id" json:"type_id"`
	CategoryID     *int       `db:"category_id" json:"category_id"`
	LastStorageID  *uuid.UUID `db:"last_storage_id" json:"last_storage_id"`
	LastWorkerID   *uuid.UUID `db:"last_worker_id" json:"last_worker_id"`
	TransferStatus string    `db:"transfer_status" json:"transfer_status"`
	QualityStatus  *string    `db:"quality_status" json:"quality_status"`
	PurchasePrice  *float64   `db:"purchase_price" json:"purchase_price"`
	OccupiedCells  *int       `db:"occupied_cells" json:"occupied_cells"`
}

type Tech struct {
	Item
	ItemID            uuid.UUID `db:"item_id" json:"item_id"`
	Brand             string    `db:"brand" json:"brand"`
	Model             string    `db:"model" json:"model"`
	WarrantyStartedAt *time.Time `db:"warranty_started_at" json:"warranty_started_at"`
	WarrantyEndAt     *time.Time `db:"warranty_end_at" json:"warranty_end_at"`
}

type Document struct {
	Item
	ItemID                uuid.UUID  `db:"item_id" json:"item_id"`
	ResponsibleWorkerID   uuid.UUID  `db:"responsible_worker_id" json:"responsible_worker_id"`
	FullSignedAt          time.Time  `db:"full_signed_at" json:"full_signed_at"`
	ResponsibleWorkerEmail string    `db:"responsible_worker_email" json:"responsible_worker_email"`
	NeededSigns           bool       `db:"needed_signs" json:"needed_signs"`
	ReceivedSigns         bool       `db:"received_signs" json:"received_signs"`
	DocNumber             string     `db:"doc_number" json:"doc_number"`
	DocType               string     `db:"doc_type" json:"doc_type"`
}

type Merch struct {
	Item
	ItemID uuid.UUID       `db:"item_id" json:"item_id"`
	Title  string          `db:"title" json:"title"`
	Size   string          `db:"size" json:"size"`
	Price  int `db:"price" json:"price"`
	Color  string          `db:"color" json:"color"`
}

type Software struct {
	Item
	ItemID            uuid.UUID  `db:"item_id" json:"item_id"`
	Vendor            string     `db:"vendor" json:"vendor"`
	LicenseKey        string     `db:"license_key" json:"license_key"`
	Title             string     `db:"title" json:"title"`
	ResponsibleWorker string     `db:"responsible_worker" json:"responsible_worker"`
	StartedAt         time.Time  `db:"started_at" json:"started_at"`
	ExpiredAt         time.Time  `db:"expired_at" json:"expired_at"`
	UpdatedAt         time.Time  `db:"updated_at" json:"updated_at"`
}

type Consumable struct {
	Item
	ItemID   uuid.UUID `db:"item_id" json:"item_id"`
	Title    string    `db:"title" json:"title"`
	Quantity int       `db:"quantity" json:"quantity"`
	Unit     string    `db:"unit" json:"unit"`
}