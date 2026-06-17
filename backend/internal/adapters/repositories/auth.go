package repositories

import (
	"database/sql"
	"inno-accounting/internal/domain"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	"time"

	"github.com/google/uuid"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

// FIND USER BY EMAIL IN DATABASE
func (authRepo *AuthRepository) FindByEmail(email string) (*domain.User, error) {
	findedUser := &domain.User{}
	
	userQuery := `
		SELECT
			id,
			email,
			hashed_password,
			role,
			is_active
		FROM users
		WHERE email = $1
	`

	err := authRepo.db.QueryRow(userQuery, email).Scan(
		&findedUser.ID,
		&findedUser.Email,
		&findedUser.HashedPassword,
		&findedUser.Role,
		&findedUser.IsActive,
	)

	if err != nil {
		logger.Error("db", err)
		if err == sql.ErrNoRows {
			return nil, app_errors.NotFound("user not found", err)
		}
		return nil, app_errors.Internal("server unavailable now. Try again later", err)
	}

	return findedUser, nil
}

// SAVE REFRESH TOKEN
func (authRepo *AuthRepository) SaveRefreshToken(userID uuid.UUID, refresh string, expireAt time.Time) error {
	query := `
		INSERT INTO refresh_tokens (user_id, refresh, expire_at, is_active)
		VALUES ($1, $2, $3, true)
	`

	_, err := authRepo.db.Exec(query, userID, refresh, expireAt)
	if err != nil {
		logger.Error("db", err)
		return app_errors.Internal("failed to save refresh token", err)
	}

	return nil
}

// FIND REFRESH TOKEN
func (authRepo *AuthRepository) FindRefreshToken(refresh string) (uuid.UUID, error) {
	var userID uuid.UUID
	var isActive bool

	query := `
		SELECT user_id, is_active
		FROM refresh_tokens
		WHERE refresh = $1
	`

	err := authRepo.db.QueryRow(query, refresh).Scan(&userID, &isActive)
	if err != nil {
		logger.Error("db", err)
		if err == sql.ErrNoRows {
			return uuid.Nil, app_errors.NotFound("refresh token not found", err)
		}
		return uuid.Nil, app_errors.Internal("server unavailable", err)
	}

	if !isActive {
		return uuid.Nil, app_errors.Unprocessable("token inactive", nil)
	}

	return userID, nil
}

// DEACTIVATE REFRESH TOKEN
func (authRepo *AuthRepository) DeactivateRefreshToken(refresh string) error {
	query := `
		UPDATE refresh_tokens
		SET is_active = false
		WHERE refresh = $1
	`

	_, err := authRepo.db.Exec(query, refresh)
	if err != nil {
		logger.Error("db", err)
		return app_errors.Internal("failed to deactivate token", err)
	}

	return nil
}

func (authRepo *AuthRepository) CreateActivationToken(
	userID uuid.UUID,
	token string,
	expiresAt time.Time,
) error {

	query := `
		INSERT INTO activation_tokens (
			id,
			user_id,
			token,
			expires_at
		)
		VALUES ($1, $2, $3, $4)
	`

	_, err := authRepo.db.Exec(
		query,
		uuid.New(),
		userID,
		token,
		expiresAt,
	)

	return err
}

func (authRepo *AuthRepository) FindActivationToken(
	token string,
) (*domain.ActivationToken, error) {

	act := &domain.ActivationToken{}

	query := `
		SELECT
			id,
			user_id,
			token,
			expires_at,
			used
		FROM activation_tokens
		WHERE token = $1
	`

	err := authRepo.db.QueryRow(
		query,
		token,
	).Scan(
		&act.ID,
		&act.UserID,
		&act.Token,
		&act.ExpiresAt,
		&act.Used,
	)

	if err != nil {
		return nil, err
	}

	return act, nil
}

func (authRepo *AuthRepository) ActivateUser(
	userID uuid.UUID,
) error {

	query := `
		UPDATE users
		SET is_active = true
		WHERE id = $1
	`

	_, err := authRepo.db.Exec(query, userID)

	return err
}

func (authRepo *AuthRepository) MarkActivationTokenUsed(
	token string,
) error {

	query := `
		UPDATE activation_tokens
		SET used = true
		WHERE token = $1
	`

	_, err := authRepo.db.Exec(query, token)

	return err
}