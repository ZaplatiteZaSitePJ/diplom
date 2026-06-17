package domain

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastname"`
	Post      string    `json:"post"`
	Grade     string    `json:"grade"`
	City      string    `json:"city"`
	Role      string    `json:"role"`
	IsActive bool `json:"is_active"`

	HashedPassword string `json:"-"`
	
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (u *User) ValidateUser() error {
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("name required")
	}

	if strings.TrimSpace(u.LastName) == "" {
		return errors.New("lastname required")
	}

	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email")
	}

	validGrades := map[string]bool{
		"intern": true, "junior": true, "middle": true,
		"senior": true, "team lead": true, "manager": true,
	}

	if !validGrades[u.Grade] {
		return errors.New("invalid grade")
	}

	if u.Role != "user" && u.Role != "admin" {
		return errors.New("invalid role")
	}

	return nil
}

