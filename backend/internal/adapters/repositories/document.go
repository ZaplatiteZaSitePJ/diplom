package repositories

import (
	"database/sql"
	"fmt"
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"
	"inno-accounting/pkg/logger"
	custom_errors "inno-accounting/pkg/server_utils/errors"
	"strings"

	"github.com/google/uuid"
)

type DocumentRepository struct {
	db *sql.DB
}

func NewDocumentRepository(db *sql.DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

// ===================== CREATE =====================

func (r *DocumentRepository) Save(d *domain.Document) (*domain.Document, error) {
	queryItem := `
		INSERT INTO items 
			(id, universal_name, type_id, category_id, last_storage_id, last_worker_id, transfer_status, quality_status)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`

	queryDoc := `
		INSERT INTO docs
		(item_id, responsible_worker_id, full_signed_at, responsible_worker_email, needed_signs, received_signs, doc_number)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
	`

	tx, err := r.db.Begin()
	if err != nil {
		return nil, custom_errors.New(err, 500)
	}

	_, err = tx.Exec(
		queryItem,
		d.ID,
		d.UniversalName,
		d.TypeID,
		d.CategoryID,
		d.LastStorageID,
		d.LastWorkerID,
		d.TransferStatus,
		d.QualityStatus,
	)
	if err != nil {
		tx.Rollback()
		logger.Error("db", err)
		return nil, custom_errors.New(err, 500)
	}

	_, err = tx.Exec(
		queryDoc,
		d.ItemID,
		d.ResponsibleWorkerID,
		d.FullSignedAt,
		d.ResponsibleWorkerEmail,
		d.NeededSigns,
		d.ReceivedSigns,
		d.DocNumber,
	)
	if err != nil {
		tx.Rollback()
		logger.Error("db", err)
		return nil, custom_errors.New(err, 500)
	}

	if err := tx.Commit(); err != nil {
		return nil, custom_errors.New(err, 500)
	}

	return d, nil
}

// ===================== UPDATE =====================

func (r *DocumentRepository) Change(id uuid.UUID, d *domain.Document) (*domain.Document, error) {
	queryItem := `
		UPDATE items
		SET universal_name=$1,
			category_id=$2,
			last_storage_id=$3,
			last_worker_id=$4,
			transfer_status=$5
		WHERE id=$6
	`

	queryDoc := `
		UPDATE docs
		SET responsible_worker_id=$1,
			full_signed_at=$2,
			responsible_worker_email=$3,
			needed_signs=$4,
			received_signs=$5,
			doc_number=$6
		WHERE item_id=$7
	`

	tx, err := r.db.Begin()
	if err != nil {
		return nil, custom_errors.New(err, 500)
	}

	_, err = tx.Exec(
		queryItem,
		d.UniversalName,
		d.CategoryID,
		d.LastStorageID,
		d.LastWorkerID,
		d.TransferStatus,
		id,
	)
	if err != nil {
		tx.Rollback()
		return nil, custom_errors.New(err, 500)
	}

	_, err = tx.Exec(
		queryDoc,
		d.ResponsibleWorkerID,
		d.FullSignedAt,
		d.ResponsibleWorkerEmail,
		d.NeededSigns,
		d.ReceivedSigns,
		d.DocNumber,
		id,
	)
	if err != nil {
		tx.Rollback()
		return nil, custom_errors.New(err, 500)
	}

	if err := tx.Commit(); err != nil {
		return nil, custom_errors.New(err, 500)
	}

	return d, nil
}

// ===================== GET BY ID =====================

func (r *DocumentRepository) FindByID(id uuid.UUID) (*domain.Document, error) {
	query := `
		SELECT 
			i.id,
			i.universal_name,
			i.type_id,
			i.category_id,
			i.last_storage_id,
			i.last_worker_id,
			i.transfer_status,
			d.responsible_worker_id,
			d.full_signed_at,
			d.responsible_worker_email,
			d.needed_signs,
			d.received_signs,
			d.doc_number
		FROM items i
		JOIN docs d ON d.item_id = i.id
		WHERE i.id = $1
	`

	doc := &domain.Document{}

	err := r.db.QueryRow(query, id).Scan(
		&doc.ID,
		&doc.UniversalName,
		&doc.TypeID,
		&doc.CategoryID,
		&doc.LastStorageID,
		&doc.LastWorkerID,
		&doc.TransferStatus,
		&doc.ResponsibleWorkerID,
		&doc.FullSignedAt,
		&doc.ResponsibleWorkerEmail,
		&doc.NeededSigns,
		&doc.ReceivedSigns,
		&doc.DocNumber,
	)

	if err != nil {
		logger.Error("db", err)

		if err == sql.ErrNoRows {
			return nil, custom_errors.New(err, 404)
		}

		return nil, custom_errors.New(err, 500)
	}

	doc.ItemID = doc.ID
	return doc, nil
}

// ===================== GET ALL (FIXED) =====================

func (r *DocumentRepository) FindAll(filter *dto.DocsFilter) ([]*dto.DocsItemPublic, error) {
	baseQuery := `
		SELECT 
			i.id,
			i.universal_name,
			c.name,
			s.storage_name,
			u.email,
			i.transfer_status,
			d.responsible_worker_email,
			d.full_signed_at,
			d.needed_signs,
			d.received_signs,
			d.doc_number
		FROM items i
		JOIN docs d ON d.item_id = i.id
		LEFT JOIN storages s ON s.id = i.last_storage_id
		LEFT JOIN users u ON u.id = i.last_worker_id
		LEFT JOIN categories c ON c.id = i.category_id
		WHERE i.type_id = 1
	`

	args := []interface{}{}
	conditions := []string{}
	argPos := 1

	if filter != nil {
		if filter.ID != nil {
			conditions = append(conditions, fmt.Sprintf("i.id = $%d", argPos))
			args = append(args, *filter.ID)
			argPos++
		}

		if filter.UserID != nil { // 👈 ВОТ ЭТО
			conditions = append(conditions, fmt.Sprintf("i.last_worker_id = $%d", argPos))
			args = append(args, *filter.UserID)
			argPos++
		}

		if filter.DocNumber != nil {
			conditions = append(conditions, fmt.Sprintf("d.doc_number ILIKE $%d", argPos))
			args = append(args, "%"+*filter.DocNumber+"%")
			argPos++
		}

		if filter.LastWorkerEmail != nil {
			conditions = append(conditions, fmt.Sprintf("u.email ILIKE $%d", argPos))
			args = append(args, "%"+*filter.LastWorkerEmail+"%")
			argPos++
		}

		if filter.ResponsibleWorkerEmail != nil {
			conditions = append(conditions, fmt.Sprintf("d.responsible_worker_email ILIKE $%d", argPos))
			args = append(args, "%"+*filter.ResponsibleWorkerEmail+"%")
			argPos++
		}
		
		if filter.LastStorage != nil {
			conditions = append(conditions, fmt.Sprintf("s.storage_name ILIKE $%d", argPos))
			args = append(args, "%"+*filter.LastStorage+"%")
			argPos++
		}

		if filter.Category != nil {
			conditions = append(conditions, fmt.Sprintf("c.name ILIKE $%d", argPos))
			args = append(args, "%"+*filter.Category+"%")
			argPos++
		}

		// статус обычно лучше оставлять точным
		if filter.TransferStatus != nil {
			conditions = append(conditions, fmt.Sprintf("i.transfer_status = $%d", argPos))
			args = append(args, *filter.TransferStatus)
			argPos++
		}
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		logger.Error("db query error:", err)
		return nil, custom_errors.New(err, 500)
	}
	defer rows.Close()

	var items []*dto.DocsItemPublic

	for rows.Next() {
		item := &dto.DocsItemPublic{}

		err := rows.Scan(
			&item.ID,
			&item.UniversalName, // ✅ FIX
			&item.Category,
			&item.LastStorage,
			&item.LastWorkerEmail,
			&item.TransferStatus,
			&item.ResponsibleWorkerEmail,
			&item.FullSignedAt,
			&item.NeededSigns,
			&item.ReceivedSigns,
			&item.DocNumber,
		)
		if err != nil {
			logger.Error("row scan error:", err)
			return nil, custom_errors.New(err, 500)
		}

		items = append(items, item)
	}

	return items, nil
}

// ===================== CATEGORY =====================

func (r *DocumentRepository) FindCategoryIDByName(name string) (int, error) {
	var id int
	err := r.db.QueryRow(`SELECT id FROM categories WHERE name = $1`, name).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *DocumentRepository) FindCategoryNameByID(id int) (*string, error) {
	var name string
	err := r.db.QueryRow(`SELECT name FROM categories WHERE id = $1`, id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_errors.New(err, 404)
		}
		return nil, custom_errors.New(err, 500)
	}
	return &name, nil
}