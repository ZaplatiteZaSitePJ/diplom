package auth

import "inno-accounting/internal/domain"

type AuthRepository interface {
	FindByEmail(email string) (*domain.User, error)
}