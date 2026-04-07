package domain

import (
	"errors"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID 		`json:"id"`
	Username       string	`json:"username"`
	Email          string	`json:"email"`
	Role		   string   `json:"role"`
	HashedPassword string	`json:"hashed_password,omitempty"`
	CreatedAt      time.Time       `db:"created_at" json:"created_at"`
    UpdatedAt      time.Time       `db:"updated_at" json:"updated_at"`
    DeletedAt      *time.Time      `db:"deleted_at" json:"deleted_at,omitempty"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (u *User) ValidateUser() error {
	trimedUsername := strings.TrimSpace(u.Username)

	if trimedUsername == "" || utf8.RuneCountInString(trimedUsername) < 3 {
		return errors.New("username cannot be empty or less 3 letter")
	}

	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

