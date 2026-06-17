package domain

import (
	"time"

	"github.com/google/uuid"
)

type ActivationToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
	Used      bool
	CreatedAt time.Time
}