package repositories

import (
	"database/sql"
	"fmt"
	"inno-accounting/internal/domain"
	custom_errors "inno-accounting/pkg/server_utils/errors"

	"github.com/google/uuid"
)

type LocationRepo struct {
	db *sql.DB
}

func NewLocationRepo(db *sql.DB) *LocationRepo {
	return &LocationRepo{db: db}
}

func (r *LocationRepo) GetByItemID(itemID uuid.UUID) (*domain.ItemLocationDetails, error) {
	query := `
		SELECT 
			item_id,
			status,

			from_location_type,
			from_location_id,

			to_location_type,
			to_location_id,

			location_type,
			location_id
		FROM item_locations
		WHERE item_id = $1
	`

	loc := &domain.ItemLocationDetails{}

	err := r.db.QueryRow(query, itemID).Scan(
		&loc.ItemID,
		&loc.Status,

		&loc.FromLocationType,
		&loc.FromLocationID,

		&loc.ToLocationType,
		&loc.ToLocationID,

		&loc.CurrentLocationType,
		&loc.CurrentLocationID,
	)

	if err != nil {
		fmt.Printf("SQL ERROR Upsert item_locations: %v\n", err)
		if err == sql.ErrNoRows {
			return nil, custom_errors.New(err, 404)
		}
		return nil, custom_errors.New(err, 500)
	}

	return loc, nil
}

func (r *LocationRepo) Upsert(loc *domain.ItemLocation) error {
	query := `
		INSERT INTO item_locations (item_id, location_type, location_id, updated_at)
		VALUES ($1,$2,$3,NOW())
		ON CONFLICT (item_id)
		DO UPDATE SET
			location_type = EXCLUDED.location_type,
			location_id = EXCLUDED.location_id,
			updated_at = NOW()
	`

	_, err := r.db.Exec(
		query,
		loc.ItemID,
		loc.LocationType,
		loc.LocationID,
	)

	if err != nil {
		return custom_errors.New(err, 500)
	}

	return nil
}