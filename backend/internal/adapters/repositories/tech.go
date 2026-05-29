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

type TechRepository struct {
	db *sql.DB
}

func NewTechRepository(db *sql.DB) *TechRepository {
	return &TechRepository{db: db}
}

func (r *TechRepository) Save(t *domain.Tech) (*domain.Tech, error) {
	queryItem := `
		INSERT INTO items 
		(id, universal_name, type_id, category_id, last_storage_id, last_worker_id, transfer_status, quality_status, purchase_price, occupied_cells)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`

	queryTech := `
		INSERT INTO tech
		(item_id, brand, model, warranty_started_at, warranty_end_at)
		VALUES ($1,$2,$3,$4,$5)
	`

	tx, err := r.db.Begin()
	if err != nil {
		return nil, custom_errors.New(err, 500)
	}

	_, err = tx.Exec(
		queryItem,
		t.ID,
		t.UniversalName,
		t.TypeID,
		t.CategoryID,
		t.LastStorageID,
		t.LastWorkerID,
		t.TransferStatus,
		t.QualityStatus,
		t.PurchasePrice,
		t.OccupiedCells,
	)

	if err != nil {
		tx.Rollback()
		logger.Error("db", err)
		return nil, custom_errors.New(err, 500)
	}

	_, err = tx.Exec(
		queryTech,
		t.ID,
		t.Brand,
		t.Model,
		t.WarrantyStartedAt,
		t.WarrantyEndAt,
	)

	if err != nil {
		tx.Rollback()
		logger.Error("db", err)
		return nil, custom_errors.New(err, 500)
	}

	if err := tx.Commit(); err != nil {
		return nil, custom_errors.New(err, 500)
	}

	return t, nil
}

func (r *TechRepository) Change(id uuid.UUID, t *domain.Tech) (*domain.Tech, error) {
	queryItem := `
		UPDATE items
		SET universal_name=$1,
			type_id=$2,
			category_id=$3,
			last_storage_id=$4,
			last_worker_id=$5,
			transfer_status=$6,
			quality_status=$7,
			purchase_price=$8,
			occupied_cells=$9
		WHERE id=$10
	`

	queryTech := `
		UPDATE tech
		SET brand=$1,
			model=$2,
			warranty_started_at=$3,
			warranty_end_at=$4
		WHERE item_id=$5
	`

	tx, err := r.db.Begin()
	if err != nil {
		return nil, custom_errors.New(err, 500)
	}

	_, err = tx.Exec(
		queryItem,
		t.UniversalName,
		t.TypeID,
		t.CategoryID,
		t.LastStorageID,
		t.LastWorkerID,
		t.TransferStatus,
		t.QualityStatus,
		t.PurchasePrice,
		t.OccupiedCells,
		id,
	)

	if err != nil {
		tx.Rollback()
		return nil, custom_errors.New(err, 500)
	}

	_, err = tx.Exec(
		queryTech,
		t.Brand,
		t.Model,
		t.WarrantyStartedAt,
		t.WarrantyEndAt,
		id,
	)

	if err != nil {
		tx.Rollback()
		return nil, custom_errors.New(err, 500)
	}

	if err := tx.Commit(); err != nil {
		return nil, custom_errors.New(err, 500)
	}

	return t, nil
}

func (r *TechRepository) FindByID(id uuid.UUID) (*domain.Tech, error) {
	query := `
		SELECT 
			i.id,
			i.universal_name,
			i.type_id,
			i.category_id,
			i.last_storage_id,
			i.last_worker_id,
			i.transfer_status,
			i.quality_status,
			i.purchase_price,
			i.occupied_cells,
			t.brand,
			t.model,
			t.warranty_started_at,
			t.warranty_end_at
		FROM items i
		JOIN tech t ON t.item_id = i.id
		WHERE i.id = $1
	`

	item := &domain.Tech{}

	err := r.db.QueryRow(query, id).Scan(
		&item.ID,
		&item.UniversalName,
		&item.TypeID,
		&item.CategoryID,
		&item.LastStorageID,
		&item.LastWorkerID,
		&item.TransferStatus,
		&item.QualityStatus,
		&item.PurchasePrice,
		&item.OccupiedCells,
		&item.Brand,
		&item.Model,
		&item.WarrantyStartedAt,
		&item.WarrantyEndAt,
	)

	if err != nil {
		logger.Error("db", err)

		if err == sql.ErrNoRows {
			return nil, custom_errors.New(err, 404)
		}

		return nil, custom_errors.New(err, 500)
	}

	return item, nil
}

func (r *TechRepository) FindAll(filter *dto.TechFilter) ([]*dto.TechItemPublic, error) {
	baseQuery := `
		SELECT 
			i.id,
			i.type_id,
			c.name AS category,
			s.storage_name AS last_storage,
			u.email AS last_worker_email,
			i.transfer_status,
			i.quality_status,
			i.purchase_price,
			i.occupied_cells,
			t.brand,
			t.model,
			t.warranty_started_at,
			t.warranty_end_at,
			i.universal_name
		FROM items i
		JOIN tech t ON t.item_id = i.id
		LEFT JOIN storages s ON s.id = i.last_storage_id
		LEFT JOIN users u ON u.id = i.last_worker_id
		LEFT JOIN categories c ON c.id = i.category_id
		WHERE i.type_id = 0
	`

	args := []interface{}{}
	conditions := []string{}

	if filter != nil {
		argPos := 1

		if filter.ID != nil {
			conditions = append(conditions, fmt.Sprintf("i.id = $%d", argPos))
			args = append(args, *filter.ID)
			argPos++
		}

		if filter.UserID != nil {
			conditions = append(conditions, fmt.Sprintf("i.last_worker_id = $%d", argPos))
			args = append(args, *filter.UserID)
			argPos++
		}

		if filter.Brand != nil {
			conditions = append(conditions, fmt.Sprintf("t.brand ILIKE $%d", argPos))
			args = append(args, "%"+*filter.Brand+"%")
			argPos++
		}

		if filter.Model != nil {
			conditions = append(conditions, fmt.Sprintf("t.model ILIKE $%d", argPos))
			args = append(args, "%"+*filter.Model+"%")
			argPos++
		}

		if filter.LastWorker != nil {
			conditions = append(conditions, fmt.Sprintf("u.email = $%d", argPos))
			args = append(args, *filter.LastWorker)
			argPos++
		}

		if filter.LastStorage != nil {
			conditions = append(conditions, fmt.Sprintf("s.storage_name = $%d", argPos))
			args = append(args, *filter.LastStorage)
			argPos++
		}

		if filter.Category != nil {
			conditions = append(conditions, fmt.Sprintf("c.name = $%d", argPos))
			args = append(args, *filter.Category)
			argPos++
		}

		if filter.QualityStatus != nil {
			conditions = append(conditions, fmt.Sprintf("i.quality_status = $%d", argPos))
			args = append(args, *filter.QualityStatus)
			argPos++
		}

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

	var items []*dto.TechItemPublic

	for rows.Next() {
		item := &dto.TechItemPublic{}

		err := rows.Scan(
			&item.ID,
			&item.Type_ID,
			&item.Category,
			&item.LastStorage,
			&item.LastWorkerEmail,
			&item.TransferStatus,
			&item.QualityStatus,
			&item.PurchasePrice,
			&item.OccupiedCells,
			&item.Brand,
			&item.Model,
			&item.WarrantyStartedAt,
			&item.WarrantyEndAt,
			&item.UniversalName,
		)
		if err != nil {
			logger.Error("row scan error:", err)
			return nil, custom_errors.New(err, 500)
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		logger.Error("rows error:", err)
		return nil, custom_errors.New(err, 500)
	}

	return items, nil
}

func (r *TechRepository) DeleteByID(id uuid.UUID) error {
	query := `DELETE FROM items WHERE id = $1`

	tx, err := r.db.Begin()
	if err != nil {
		return custom_errors.New(err, 500)
	}

	result, err := tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return custom_errors.New(err, 500)
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		tx.Rollback()
		return custom_errors.New(sql.ErrNoRows, 404)
	}

	return tx.Commit()
}

func (r *TechRepository) FindCategoryIDByName(name string) (int, error) {

	var id int

	query := `SELECT id FROM categories WHERE name = $1`

	err := r.db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil

}

func (r *TechRepository) FindCategoryNameByID(id int) (*string, error) {
	query := `SELECT name FROM categories WHERE id = $1`

	var name string

	err := r.db.QueryRow(query, id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, custom_errors.New(err, 404)
		}
		return nil, custom_errors.New(err, 500)
	}

	return &name, nil
}

func (r *TechRepository) GetCategoriesByTypeID(typeID int) ([]string, error) {
	query := `SELECT name 
			FROM categories 
			WHERE type_id = $1 
			ORDER BY name`

	rows, err := r.db.Query(query, typeID)
	if err != nil {
		return nil, custom_errors.New(err, 500)
	}
	defer rows.Close()

	var categories []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, custom_errors.New(err, 500)
		}
		categories = append(categories, name)
	}

	if err := rows.Err(); err != nil {
		return nil, custom_errors.New(err, 500)
	}

	return categories, nil
}

