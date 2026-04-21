package repositories

import (
	"database/sql"
	"fmt"
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	pg_err "inno-accounting/pkg/server_utils/db_errors/postgres"
	custom_errors "inno-accounting/pkg/server_utils/errors"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

//
// 🔹 SAVE USER
//
func (r *UserRepository) Save(user *domain.User) (*domain.User, error) {
	saved := &domain.User{}

	query := `
		INSERT INTO users 
		(id, email, name, lastname, post, grade, city, role, hashed_password, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING id, email, name, lastname, post, grade, city, role, created_at, updated_at
	`

	err := r.db.QueryRow(
		query,
		user.ID,
		user.Email,
		user.Name,
		user.LastName,
		user.Post,
		user.Grade,
		user.City,
		user.Role,
		user.HashedPassword,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(
		&saved.ID,
		&saved.Email,
		&saved.Name,
		&saved.LastName,
		&saved.Post,
		&saved.Grade,
		&saved.City,
		&saved.Role,
		&saved.CreatedAt,
		&saved.UpdatedAt,
	)

	if err != nil {
		logger.Error("db", err)
		if pg_err.IsUniqueViolation(err) {
			return nil, app_errors.AlreadyExists("user already exists", err)
		}
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	return saved, nil
}

//
// 🔹 FIND BY ID
//
func (r *UserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	user := &domain.User{}

	query := `
		SELECT id, email, name, lastname, post, grade, city, role
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.LastName,
		&user.Post,
		&user.Grade,
		&user.City,
		&user.Role,
	)

	if err != nil {
		logger.Error("db", err)
		if err == sql.ErrNoRows {
			return nil, app_errors.NotFound("user not found", err)
		}
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	return user, nil
}

//
// 🔹 FIND ALL
//
func (r *UserRepository) FindAll(filter *dto.UserFilter) ([]*domain.User, error) {
	users := []*domain.User{}

	query := `
		SELECT id, email, name, lastname, post, grade, city
		FROM users
		WHERE deleted_at IS NULL
	`

	args := []interface{}{}
	i := 1

	if filter != nil {
		if filter.ID != nil {
			query += fmt.Sprintf(" AND id::text ILIKE $%d", i)
			args = append(args, "%"+*filter.ID+"%")
			i++
		}

		if filter.Email != nil {
			query += fmt.Sprintf(" AND email ILIKE $%d", i)
			args = append(args, "%"+*filter.Email+"%")
			i++
		}

		if filter.Name != nil {
			query += fmt.Sprintf(" AND name ILIKE $%d", i)
			args = append(args, "%"+*filter.Name+"%")
			i++
		}

		if filter.LastName != nil {
			query += fmt.Sprintf(" AND lastname ILIKE $%d", i)
			args = append(args, "%"+*filter.LastName+"%")
			i++
		}

		if filter.Post != nil {
			query += fmt.Sprintf(" AND post ILIKE $%d", i)
			args = append(args, "%"+*filter.Post+"%")
			i++
		}

		if filter.Grade != nil {
			query += fmt.Sprintf(" AND grade ILIKE $%d", i)
			args = append(args, "%"+*filter.Grade+"%")
			i++
		}

		if filter.City != nil {
			query += fmt.Sprintf(" AND city ILIKE $%d", i)
			args = append(args, "%"+*filter.City+"%")
			i++
		}
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		logger.Error("db", err)
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := &domain.User{}

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.LastName,
			&user.Post,
			&user.Grade,
			&user.City,
		)
		if err != nil {
			logger.Error("db", err)
			return nil, app_errors.Internal("server unavailable now. Try again later", err)
		}

		users = append(users, user)
	}

	return users, nil
}

//
// 🔹 UPDATE USER
//
func (r *UserRepository) Update(user *domain.User) (*domain.User, error) {
	query := `
		UPDATE users
		SET email=$1, name=$2, lastname=$3, post=$4, grade=$5, city=$6, updated_at=$7
		WHERE id=$8 AND deleted_at IS NULL
		RETURNING id, email, name, lastname, post, grade, city, role, created_at, updated_at
	`

	updated := &domain.User{}

	err := r.db.QueryRow(
		query,
		user.Email,
		user.Name,
		user.LastName,
		user.Post,
		user.Grade,
		user.City,
		user.UpdatedAt,
		user.ID,
	).Scan(
		&updated.ID,
		&updated.Email,
		&updated.Name,
		&updated.LastName,
		&updated.Post,
		&updated.Grade,
		&updated.City,
		&updated.Role,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)

	if err != nil {
		logger.Error("db", err)
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	return updated, nil
}

//
// 🔹 DELETE (SOFT)
//
func (r *UserRepository) Delete(id uuid.UUID) error {
	query := `
		UPDATE users
		SET deleted_at = NOW()
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		logger.Error("db", err)
		return app_errors.Internal("server unavailable now. Try again later", err)
	}

	return nil
}

//
// 🔹 FIND BY EMAIL
//
func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	user := &domain.User{}

	query := `
		SELECT id, email, name, lastname, post, grade, city, role, hashed_password
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.LastName,
		&user.Post,
		&user.Grade,
		&user.City,
		&user.Role,
		&user.HashedPassword,
	)

	if err != nil {
		logger.Error("db", err)
		if err == sql.ErrNoRows {
			return nil, app_errors.NotFound(fmt.Sprintf("user with email '%s' not found", email), err)
		}
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	return user, nil
}

func (r *TechRepository) FindByUserID(userID uuid.UUID) ([]*dto.TechItemPublic, error) {
	query := `
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
		WHERE i.last_worker_id = $1
	`

	rows, err := r.db.Query(query, userID)
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