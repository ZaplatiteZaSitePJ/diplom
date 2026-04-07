package repositories

import (
	"database/sql"
	"inno-accounting/internal/domain"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	pg_err "inno-accounting/pkg/server_utils/db_errors/postgres"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// SAVE USER IN DATABASE
func (userRepo *UserRepository) Save(newUser *domain.User) (*domain.User, error) {
	var savedUser = &domain.User{}

	query := 
		`INSERT INTO users 
		(id, username, email, hashed_password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, username, email, created_at, updated_at`
		
	err := userRepo.db.QueryRow(
		query, 
		newUser.ID, 
		newUser.Username, 
		newUser.Email, 
		newUser.HashedPassword, 
		newUser.CreatedAt, 
		newUser.UpdatedAt,
	).Scan(
		&savedUser.ID, 
		&savedUser.Username, 
		&savedUser.Email, 
		&savedUser.CreatedAt, 
		&savedUser.UpdatedAt)

	if err != nil {
		logger.Error("db", err)
		if pg_err.IsUniqueViolation(err) {
			return nil, app_errors.AlreadyExists("user already exist", err)
		}
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}
	return savedUser, nil
}

// FIND USER BY ID IN DATABASE
func (userRepo *UserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	findedUser := &domain.User{}
	
	userQuery := "SELECT id, username, email, role FROM users WHERE id = $1"

	err := userRepo.db.QueryRow(userQuery, id).Scan(&findedUser.ID, &findedUser.Username, &findedUser.Email, &findedUser.Role)

	if err != nil {
		logger.Error("db", err)
		if err == sql.ErrNoRows {
			return nil, app_errors.NotFound("user not found", err)
		}
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	return findedUser, nil
}

// // FIND USER BY EMAIL IN DATABASE
// func (userRepo *UserRepository) FindByEmail(email string) (*domain.User, error) {
// 	findedUser := &domain.User{}
	
// 	userQuery := "SELECT id, username, email FROM users WHERE email = $1"

// 	err := userRepo.db.QueryRow(userQuery, email).Scan(&findedUser.ID, &findedUser.Username, &findedUser.Email)

// 	if err != nil {
// 		logger.Error("db", err)
// 		if err == sql.ErrNoRows {
// 			return nil, app_errors.NotFound("user not found", err)
// 		}
// 		return nil, app_errors.Internal("server unavailable now. Try again later", err)
// 	}

// 	return findedUser, nil
// }

// FIND USER BY SEVERAL ID's IN DATABASE
func (userRepo *UserRepository) FindBySeveralIDs(ids []uuid.UUID) ([]*domain.User, error) {
	if len(ids) == 0 {
        return []*domain.User{}, nil
    }

	findedUsers := make([]*domain.User, 0, len(ids))

	query :=
		`
		SELECT id, username, email FROM users
		WHERE id = ANY($1)
		`

	rows, err:= userRepo.db.Query(query, pq.Array(ids))

	if err != nil {
		logger.Error("db", err)
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	defer rows.Close()

	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		
		if err != nil {
			logger.Error("db", err)
			return nil, app_errors.Internal("server unavailable now. Try again later", err)
		}

		findedUsers = append(findedUsers, user)
	}

	return findedUsers, nil
}

func (userRepo *UserRepository) FindAll() ([]*domain.User, error) {
	
	findedUsers := []*domain.User{}

	query := "SELECT id, username, email FROM users"

	rows, err := userRepo.db.Query(query)
	if err != nil {
		logger.Error("db", err)
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	defer rows.Close()

	for rows.Next() {
		user := &domain.User{}
		rows.Scan(&user.ID, &user.Username, &user.Email)
		findedUsers = append(findedUsers, user)
	}

	if err := rows.Err(); err != nil {
		logger.Error("db", err)
        return nil, app_errors.Internal("server unavailable now. Try again later", err)
    }

	return findedUsers, nil
}

func (userRepo *UserRepository) DeleteByID(id int) error {
	return nil
}