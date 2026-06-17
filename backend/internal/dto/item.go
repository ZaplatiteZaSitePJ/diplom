package dto

import (
	"time"

	"github.com/google/uuid"
)

type TechItemPublic struct {
	ID				  uuid.UUID `json:"id"`
	UniversalName     string  `json:"universal_name"`
	Type_ID           int `json:"type_id"`
	Category          *string    `json:"category"`
	LastStorage       *string    `json:"last_storage"`
	LastWorkerEmail   *string    `json:"last_worker_email"`
	TransferStatus    string    `json:"transfer_status"`
	QualityStatus     string    `json:"quality_status"`
	PurchasePrice     float64   `json:"purchase_price"`
	OccupiedCells     int       `json:"occupied_cells"`
	LastStorageID *uuid.UUID `json:"last_storage_id"`
	PostNumber *string `json:"post_number,omitempty"`

	Brand             string    `json:"brand"`
	Model             string    `json:"model"`
	WarrantyStartedAt *time.Time `json:"warranty_started_at"`
	WarrantyEndAt     *time.Time `json:"warranty_end_at"`
}

type TechFilter struct {
	ID            *string `json:"id"`             // uuid как string
	Brand         *string `json:"brand"`
	Model         *string `json:"model"`
	LastWorker    *string `json:"last_worker"`    // email или имя
	LastStorage   *string `json:"last_storage"`
	LastStorageID *uuid.UUID `json:"last_storage_id"`
	Category      *string `json:"category"`
	QualityStatus *string `json:"quality_status"`
	TransferStatus *string `json:"transfer_status"`
	UserID         *uuid.UUID 
}

type SoftwareItemPublic struct {
	ID                uuid.UUID `json:"id"`
	UniversalName     string    `json:"universal_name"`
	Type              string    `json:"type"`
	Category          string    `json:"category"`
	LastWorkerEmail        *string    `json:"last_worker_email"`
	TransferStatus    *string    `json:"transfer_status"`
	PurchasePrice     float64   `json:"purchase_price"`

	Vendor            string     `json:"vendor"`
	LicenseKey        string     `json:"license_key"`
	Title             string     `json:"title"`
	StartedAt         *time.Time  `json:"started_at"`
	ExpiredAt         *time.Time  `json:"expired_at"`
	UpdatedAt         *time.Time  `json:"updated_at"`
}

type SoftwareFilter struct {
	ID                *string  `json:"id"`
	Category          *string  `json:"category"`
	LastWorkerEmail   *string  `json:"last_worker_email"`
	Vendor            *string  `json:"vendor"`
	LicenseKey        *string  `json:"license_key"`
	Title             *string  `json:"title"`
	PurchasePrice     *float64 `json:"purchase_price"`

	TransferStatus    *string  `json:"transfer_status"`   
	LastStorage       *string  `json:"last_storage"`
	
	ExpiredAt         *string `json:"expired_at"`
	UserID         *uuid.UUID 
}

type DocsItemPublic struct {
	ID                    uuid.UUID `json:"id"`
	UniversalName         string    `json:"universal_name"`
	Type                  string    `json:"type"`
	Category              string    `json:"category"`
	LastStorage           *string    `json:"last_storage"`
	LastWorkerEmail            *string    `json:"last_worker_email"`
	TransferStatus        *string    `json:"transfer_status"`

	FullSignedAt           *time.Time `json:"full_signed_at"`
	ResponsibleWorkerEmail string    `json:"responsible_worker_email"`
	NeededSigns            int      `json:"needed_signs"`
	ReceivedSigns          int      `json:"received_signs"`
	DocNumber              string    `json:"doc_number"`
}

type DocsFilter struct {
	ID             *uuid.UUID `json:"id,omitempty"`
	DocNumber      *string    `json:"doc_number,omitempty"`
	LastWorkerEmail     *string    `json:"last_worker_email,omitempty"`
	LastStorage    *string    `json:"last_storage,omitempty"`
	Category       *string    `json:"category,omitempty"`
	TransferStatus *string    `json:"transfer_status,omitempty"`
	ResponsibleWorkerEmail *string `json:"responsible_worker_email,omitempty"`
	UserID         *uuid.UUID 
}

