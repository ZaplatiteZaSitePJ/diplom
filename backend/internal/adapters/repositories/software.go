package repositories

import (
	"database/sql"
	"fmt"
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"
	"strings"

	"github.com/google/uuid"
)

type SoftwareRepository struct {
	db *sql.DB
}

func NewSoftwareRepository(db *sql.DB) *SoftwareRepository {
	return &SoftwareRepository{db: db}
}

func (r *SoftwareRepository) Save(s *domain.Software) (*domain.Software, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`
		INSERT INTO items 
		(id, universal_name, type_id, category_id, last_worker_id, purchase_price, transfer_status, quality_status)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`,
		s.ID,
		s.UniversalName,
		s.TypeID,
		s.CategoryID,
		s.LastWorkerID,
		s.PurchasePrice,
		s.TransferStatus,
		"new",
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec(`
		INSERT INTO software 
		(item_id, vendor, license_key, title, started_at, expired_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,NOW())
	`,
		s.ID,
		s.Vendor,
		s.LicenseKey,
		s.Title,
		s.StartedAt,
		s.ExpiredAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return s, tx.Commit()
}

func (r *SoftwareRepository) FindByID(id uuid.UUID) (*domain.Software, error) {
	query := `
	SELECT 
		i.id,
		i.universal_name,
		i.category_id,
		i.last_worker_id,
		i.purchase_price,
		s.vendor,
		s.license_key,
		s.title,
		s.started_at,
		s.expired_at,
		s.updated_at
	FROM items i
	JOIN software s ON s.item_id = i.id
	WHERE i.id = $1
	`

	item := &domain.Software{}

	err := r.db.QueryRow(query, id).Scan(
		&item.ID,
		&item.UniversalName,
		&item.CategoryID,
		&item.LastWorkerID,
		&item.PurchasePrice,
		&item.Vendor,
		&item.LicenseKey,
		&item.Title,
		&item.StartedAt,
		&item.ExpiredAt,
		&item.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *SoftwareRepository) FindAll(filter *dto.SoftwareFilter) ([]*dto.SoftwareItemPublic, error) {
	baseQuery := `
	SELECT 
		i.id,
		i.universal_name,
		c.name,
		s.storage_name,
		u.email,
		i.transfer_status,
		i.purchase_price,
		sw.vendor,
		sw.license_key,
		sw.title,
		sw.started_at,
		sw.expired_at,
		sw.updated_at
	FROM items i
	JOIN software sw ON sw.item_id = i.id
	LEFT JOIN categories c ON c.id = i.category_id
	LEFT JOIN users u ON u.id = i.last_worker_id
	LEFT JOIN storages s ON s.id = i.last_storage_id
	WHERE i.type_id = 1
	`

	args := []interface{}{}
	conditions := []string{}
	argPos := 1

	if filter != nil {

		if filter.ID != nil && *filter.ID != "" {
			conditions = append(conditions, fmt.Sprintf("i.id = $%d", argPos))
			args = append(args, *filter.ID)
			argPos++
		}

		if filter.UserID != nil {
			conditions = append(conditions, fmt.Sprintf("i.last_worker_id = $%d", argPos))
			args = append(args, *filter.UserID)
			argPos++
		}

		if filter.Category != nil && *filter.Category != "" {
			conditions = append(conditions, fmt.Sprintf("c.name = $%d", argPos))
			args = append(args, *filter.Category)
			argPos++
		}

		if filter.LastWorkerEmail != nil && *filter.LastWorkerEmail != "" {
			conditions = append(conditions, fmt.Sprintf("u.email = $%d", argPos))
			args = append(args, *filter.LastWorkerEmail)
			argPos++
		}

		if filter.LastStorage != nil && *filter.LastStorage != "" {
			conditions = append(conditions, fmt.Sprintf("s.storage_name = $%d", argPos))
			args = append(args, *filter.LastStorage)
			argPos++
		}

		if filter.TransferStatus != nil && *filter.TransferStatus != "" {
			conditions = append(conditions, fmt.Sprintf("i.transfer_status = $%d", argPos))
			args = append(args, *filter.TransferStatus)
			argPos++
		}

		if filter.Vendor != nil && *filter.Vendor != "" {
			conditions = append(conditions, fmt.Sprintf("sw.vendor ILIKE $%d", argPos))
			args = append(args, "%"+*filter.Vendor+"%")
			argPos++
		}

		if filter.Title != nil && *filter.Title != "" {
			conditions = append(conditions, fmt.Sprintf("sw.title ILIKE $%d", argPos))
			args = append(args, "%"+*filter.Title+"%")
			argPos++
		}

		if filter.LicenseKey != nil && *filter.LicenseKey != "" {
			conditions = append(conditions, fmt.Sprintf("sw.license_key ILIKE $%d", argPos))
			args = append(args, "%"+*filter.LicenseKey+"%")
			argPos++
		}

		// 🔥 FIXED expired_at (СТРОКА → DATE SAFE)
		if filter.ExpiredAt != nil && *filter.ExpiredAt != "" {
			conditions = append(conditions,
				fmt.Sprintf("sw.expired_at < ($%d::date + INTERVAL '1 day')", argPos),
			)
			args = append(args, *filter.ExpiredAt)
			argPos++
		}
	}

	if len(conditions) > 0 {
		baseQuery += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*dto.SoftwareItemPublic

	for rows.Next() {
		item := &dto.SoftwareItemPublic{}

		var category sql.NullString
		var storage sql.NullString
		var email sql.NullString
		var transferStatus sql.NullString
		var purchasePrice sql.NullFloat64
		var startedAt sql.NullTime
		var expiredAt sql.NullTime
		var updatedAt sql.NullTime

		err := rows.Scan(
			&item.ID,
			&item.UniversalName,
			&category,
			&storage,
			&email,
			&transferStatus,
			&purchasePrice,
			&item.Vendor,
			&item.LicenseKey,
			&item.Title,
			&startedAt,
			&expiredAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		if category.Valid {
			item.Category = category.String
		}

		if email.Valid {
			item.LastWorkerEmail = &email.String
		}

		if purchasePrice.Valid {
			item.PurchasePrice = purchasePrice.Float64
		}

		if startedAt.Valid {
			item.StartedAt = &startedAt.Time
		}

		if expiredAt.Valid {
			item.ExpiredAt = &expiredAt.Time
		}

		if updatedAt.Valid {
			item.UpdatedAt = &updatedAt.Time
		}

		result = append(result, item)
	}

	return result, nil
}

func (r *SoftwareRepository) Change(id uuid.UUID, s *domain.Software) (*domain.Software, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`
		UPDATE items
		SET universal_name=$1,
			category_id=$2,
			last_worker_id=$3,
			purchase_price=$4
		WHERE id=$5
	`,
		s.UniversalName,
		s.CategoryID,
		s.LastWorkerID,
		s.PurchasePrice,
		id,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec(`
		UPDATE software
		SET vendor=$1,
			license_key=$2,
			title=$3,
			started_at=$4,
			expired_at=$5,
			updated_at=NOW()
		WHERE item_id=$6
	`,
		s.Vendor,
		s.LicenseKey,
		s.Title,
		s.StartedAt,
		s.ExpiredAt,
		id,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return s, tx.Commit()
}

func (r *SoftwareRepository) FindCategoryIDByName(name string) (int, error) {
	var id int

	query := `SELECT id FROM categories WHERE name = $1`

	err := r.db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *SoftwareRepository) FindCategoryNameByID(id int) (*string, error) {
	var name string

	query := `SELECT name FROM categories WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(&name)
	if err != nil {
		return nil, err
	}

	return &name, nil
}

func (r *SoftwareRepository) GetCategoriesByTypeID(typeID int) ([]string, error) {
	query := `
		SELECT name 
		FROM categories 
		WHERE type_id = $1
	`

	rows, err := r.db.Query(query, typeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		result = append(result, name)
	}

	return result, nil
}