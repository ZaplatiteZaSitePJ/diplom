package domain

import "github.com/google/uuid"

type LocationType string

const (
	LocationStorage LocationType = "storage"
	LocationUser    LocationType = "user"
	LocationTransit LocationType = "transit"
)

type ItemLocation struct {
	ItemID       uuid.UUID
	LocationType LocationType
	LocationID   *uuid.UUID
}

type ItemLocationDetails struct {
	ItemID uuid.UUID

	Status string

	FromLocationType LocationType
	FromLocationID   *uuid.UUID

	ToLocationType LocationType
	ToLocationID   *uuid.UUID

	CurrentLocationType LocationType
	CurrentLocationID   *uuid.UUID
}

// type MovementStatus string

// const (
// 	MovementInProgress MovementStatus = "in_progress"
// 	MovementCompleted  MovementStatus = "completed"
// )

// type ItemMovement struct {
// 	ID     uuid.UUID
// 	ItemID uuid.UUID

// 	FromType LocationType
// 	FromID   *uuid.UUID

// 	ToType LocationType
// 	ToID   *uuid.UUID

// 	Status MovementStatus
// }
