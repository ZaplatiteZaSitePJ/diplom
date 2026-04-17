package repositories

import (
	"database/sql"
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"

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
	query := `
	SELECT 
		i.id,
		i.universal_name,
		c.name,
		u.email,
		i.purchase_price,
		s.vendor,
		s.license_key,
		s.title,
		s.started_at,
		s.expired_at,
		s.updated_at
	FROM items i
	JOIN software s ON s.item_id = i.id
	LEFT JOIN categories c ON c.id = i.category_id
	LEFT JOIN users u ON u.id = i.last_worker_id
	WHERE i.type_id = 1
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*dto.SoftwareItemPublic

	for rows.Next() {
		item := &dto.SoftwareItemPublic{}

		var category sql.NullString
		var email sql.NullString
		var purchasePrice sql.NullFloat64
		var startedAt sql.NullTime
		var expiredAt sql.NullTime
		var updatedAt sql.NullTime

		err := rows.Scan(
			&item.ID,
			&item.UniversalName,
			&category,
			&email,
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

		// category
		if category.Valid {
			item.Category = category.String
		} else {
			item.Category = ""
		}

		// last worker
		if email.Valid {
			item.LastWorkerEmail = &email.String
		} else {
			item.LastWorkerEmail = nil
		}

		// purchase price
		if purchasePrice.Valid {
			item.PurchasePrice = purchasePrice.Float64
		} else {
			item.PurchasePrice = 0
		}

		// started at
		if startedAt.Valid {
			item.StartedAt = &startedAt.Time
		} else {
			item.StartedAt = nil
		}

		// expired at
		if expiredAt.Valid {
			item.ExpiredAt = &expiredAt.Time
		} else {
			item.ExpiredAt = nil
		}

		// updated at
		if updatedAt.Valid {
			item.UpdatedAt = &updatedAt.Time
		} else {
			item.UpdatedAt = nil
		}

		result = append(result, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
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