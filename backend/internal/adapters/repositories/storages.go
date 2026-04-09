package repositories

import (
	"database/sql"
	"fmt"
	"inno-accounting/internal/domain"
	"inno-accounting/pkg/logger"
	pg_err "inno-accounting/pkg/server_utils/db_errors/postgres"
	custom_errors "inno-accounting/pkg/server_utils/errors"

	"github.com/google/uuid"
)

type StorageRepository struct {
	db *sql.DB
}

func NewStorageRepository(db *sql.DB) *StorageRepository {
	return &StorageRepository{
		db: db,
	}
}

// SAVE STORAGE IN DATABASE
func (storageRepo *StorageRepository) Save(newStorage *domain.Storage) (*domain.Storage, error) {
	var savedStorage = &domain.Storage{}

	query := `
		INSERT INTO storages 
		(id, storage_name, capacity, occupied_cells, city, items_amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, storage_name, capacity, occupied_cells, city, items_amount, created_at, updated_at, deleted_at
	`

	err := storageRepo.db.QueryRow(
		query,
		newStorage.ID,
		newStorage.StorageName,
		newStorage.Capacity,
		newStorage.OccupiedCells,
		newStorage.City,
		newStorage.ItemsAmount,
		newStorage.CreatedAt,
		newStorage.UpdatedAt,
	).Scan(
		&savedStorage.ID,
		&savedStorage.StorageName,
		&savedStorage.Capacity,
		&savedStorage.OccupiedCells,
		&savedStorage.City,
		&savedStorage.ItemsAmount,
		&savedStorage.CreatedAt,
		&savedStorage.UpdatedAt,
		&savedStorage.DeletedAt,
	)

	if err != nil {
		logger.Error("db", err)

		if pg_err.IsUniqueViolation(err) {
			return nil, custom_errors.New(err, 409)
		}

		return nil, custom_errors.New(err, 500)
	}

	return savedStorage, nil
}

// CHANGE STORAGE IN DATABASE
func (storageRepo *StorageRepository) Change(id uuid.UUID, updatedStorage *domain.Storage) (*domain.Storage, error) {
	query := `
		UPDATE storages
		SET storage_name = $1,
			capacity = $2,
			city = $3,
			occupied_cells = $4,
			items_amount = $5,
			updated_at = $6
		WHERE id = $7 AND deleted_at IS NULL
		RETURNING id, storage_name, capacity, occupied_cells, city, items_amount, created_at, updated_at, deleted_at
	`

	storage := &domain.Storage{}

	err := storageRepo.db.QueryRow(
		query,
		updatedStorage.StorageName,
		updatedStorage.Capacity,
		updatedStorage.City,
		updatedStorage.OccupiedCells,
		updatedStorage.ItemsAmount,
		updatedStorage.UpdatedAt,
		id,
	).Scan(
		&storage.ID,
		&storage.StorageName,
		&storage.Capacity,
		&storage.OccupiedCells,
		&storage.City,
		&storage.ItemsAmount,
		&storage.CreatedAt,
		&storage.UpdatedAt,
		&storage.DeletedAt,
	)

	if err != nil {
		logger.Error("db", err)

		if err == sql.ErrNoRows {
			return nil, custom_errors.New(err, 404)
		}

		return nil, custom_errors.New(err, 500)
	}

	return storage, nil
}

// FIND USER BY ID IN DATABASE
func (storageRepo *StorageRepository) FindByID(id uuid.UUID) (*domain.Storage, error) {
	storage := &domain.Storage{}

	query := `
		SELECT id, storage_name, capacity, occupied_cells, city, items_amount, created_at, updated_at, deleted_at
		FROM storages
		WHERE id = $1 AND deleted_at IS NULL
	`

	err := storageRepo.db.QueryRow(query, id).Scan(
		&storage.ID,
		&storage.StorageName,
		&storage.Capacity,
		&storage.OccupiedCells,
		&storage.City,
		&storage.ItemsAmount,
		&storage.CreatedAt,
		&storage.UpdatedAt,
		&storage.DeletedAt,
	)

	if err != nil {
		logger.Error("db", err)

		if err == sql.ErrNoRows {
			return nil, custom_errors.New(err, 404)
		}

		return nil, custom_errors.New(err, 500)
	}

	return storage, nil
}

func (storageRepo *StorageRepository) FindAll() ([]*domain.Storage, error) {
	storages := []*domain.Storage{}

	query := `
		SELECT id, storage_name, capacity, occupied_cells, city, items_amount, created_at, updated_at, deleted_at
		FROM storages
		WHERE deleted_at IS NULL
	`

	rows, err := storageRepo.db.Query(query)
	if err != nil {
		logger.Error("db", err)
		return nil, custom_errors.New(err, 500)
	}
	defer rows.Close()

	for rows.Next() {
		storage := &domain.Storage{}

		err := rows.Scan(
			&storage.ID,
			&storage.StorageName,
			&storage.Capacity,
			&storage.OccupiedCells,
			&storage.City,
			&storage.ItemsAmount,
			&storage.CreatedAt,
			&storage.UpdatedAt,
			&storage.DeletedAt,
		)
		if err != nil {
			logger.Error("db", err)
			return nil, custom_errors.New(err, 500)
		}

		storages = append(storages, storage)
	}

	if err := rows.Err(); err != nil {
		logger.Error("db", err)
		return nil, custom_errors.New(err, 500)
	}

	return storages, nil
}

func (storageRepo *StorageRepository) DeleteByID(id uuid.UUID) error {
	query := `
		UPDATE storages
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := storageRepo.db.Exec(query, id)
	if err != nil {
		logger.Error("db", err)
		return custom_errors.New(err, 500)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("db", err)
		return custom_errors.New(err, 500)
	}

	if rowsAffected == 0 {
		return custom_errors.New(sql.ErrNoRows, 404)
	}

	return nil
}

// FindByName ищет склады, где имя содержит переданную строку (без учета регистра)
func (storageRepo *StorageRepository) FindByName(name string) ([]*domain.Storage, error) {
	storages := []*domain.Storage{}

	query := `
		SELECT id, storage_name, capacity, occupied_cells, city, items_amount, created_at, updated_at, deleted_at
		FROM storages
		WHERE deleted_at IS NULL AND storage_name ILIKE $1
	`

	rows, err := storageRepo.db.Query(query, "%"+name+"%")
	if err != nil {
		logger.Error("db", err)
		return nil, custom_errors.New(err, 500)
	}
	defer rows.Close()

	for rows.Next() {
		storage := &domain.Storage{}
		err := rows.Scan(
			&storage.ID,
			&storage.StorageName,
			&storage.Capacity,
			&storage.OccupiedCells,
			&storage.City,
			&storage.ItemsAmount,
			&storage.CreatedAt,
			&storage.UpdatedAt,
			&storage.DeletedAt,
		)
		if err != nil {
			logger.Error("db", err)
			return nil, custom_errors.New(err, 500)
		}
		storages = append(storages, storage)
	}

	if err := rows.Err(); err != nil {
		logger.Error("db", err)
		return nil, custom_errors.New(err, 500)
	}

	logger.Info(fmt.Sprintf("Found %d storages matching name '%s'", len(storages), name))
	return storages, nil
}

func (storageRepo *StorageRepository) FindByExactName(name string) (*domain.Storage, error) {
	storage := &domain.Storage{}

	query := `
		SELECT id, storage_name, capacity, occupied_cells, city, items_amount, created_at, updated_at, deleted_at
		FROM storages
		WHERE deleted_at IS NULL AND storage_name = $1
		LIMIT 1
	`

	err := storageRepo.db.QueryRow(query, name).Scan(
		&storage.ID,
		&storage.StorageName,
		&storage.Capacity,
		&storage.OccupiedCells,
		&storage.City,
		&storage.ItemsAmount,
		&storage.CreatedAt,
		&storage.UpdatedAt,
		&storage.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, custom_errors.New(err, 500)
	}

	return storage, nil
}