package dto

import (
	"time"

	"github.com/google/uuid"
)

type TechItemPublic struct {
	Type_ID           int `json:"type_id"`
	Category          *string    `json:"category"`
	LastStorage       *string    `json:"last_storage"`
	LastWorkerEmail   *string    `json:"last_worker_email"`
	TransferStatus    string    `json:"transfer_status"`
	QualityStatus     string    `json:"quality_status"`
	PurchasePrice     float64   `json:"purchase_price"`
	OccupiedCells     int       `json:"occupied_cells"`

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
	Category      *string `json:"category"`
	QualityStatus *string `json:"quality_status"`
}

type SoftwareItemPublic struct {
	ID                uuid.UUID `json:"id"`
	UniversalName     string    `json:"universal_name"`
	Type              string    `json:"type"`
	Category          string    `json:"category"`
	LastWorker        string    `json:"last_worker"`
	TransferStatus    string    `json:"transfer_status"`
	PurchasePrice     float64   `json:"purchase_price"`

	Vendor            string     `json:"vendor"`
	LicenseKey        string     `json:"license_key"`
	Title             string     `json:"title"`
	StartedAt         time.Time  `json:"started_at"`
	ExpiredAt         time.Time  `json:"expired_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type MerchItemPublic struct {
	ID                uuid.UUID `json:"id"`
	UniversalName     string    `json:"universal_name"`
	Type              string    `json:"type"`
	Category          string    `json:"category"`
	LastStorage       string    `json:"last_storage"`
	LastWorker        string    `json:"last_worker"`
	TransferStatus    string    `json:"transfer_status"`
	QualityStatus     string    `json:"quality_status"`
	PurchasePrice     float64   `json:"purchase_price"`
	OccupiedCells     int       `json:"occupied_cells"`

	Title  string          `json:"title"`
	Size   string          `json:"size"`
	Price  int             `json:"price"`
	Color  string          `json:"color"`
}

type DocsItemPublic struct {
	ID                uuid.UUID `json:"id"`
	UniversalName     string    `json:"universal_name"`
	Type              string    `json:"type"`
	Category          string    `json:"category"`
	LastStorage       string    `json:"last_storage"`
	LastWorker        string    `json:"last_worker"`
	TransferStatus    string    `json:"transfer_status"`

	ResponsibleWorker        string  `json:"responsible_worker"`
	FullSignedAt             time.Time  `json:"full_signed_at"`
	ResponsibleWorkerEmail   string    `json:"responsible_worker_email"`
	NeededSigns              bool       `json:"needed_signs"`
	ReceivedSigns            bool       `json:"received_signs"`
	DocNumber                string     `json:"doc_number"`
	DocType                  string     `json:"doc_type"`
}

