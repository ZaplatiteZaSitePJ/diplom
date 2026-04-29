package location

import (
	"errors"
	"inno-accounting/internal/domain"
	"inno-accounting/pkg/server_utils/app_errors"

	"github.com/google/uuid"
)

type LocationService struct {
	repo LocationRepository
}

func NewLocationService(repo LocationRepository) *LocationService {
	return &LocationService{repo: repo}
}


func (s *LocationService) GetLocation(itemID uuid.UUID) (*domain.ItemLocationDetails, error) {
	loc, err := s.repo.GetByItemID(itemID)
	if err != nil {
		return nil, err
	}
	return loc, nil
}

func (s *LocationService) MoveItem(
	itemID uuid.UUID,
	toType domain.LocationType,
	toID *uuid.UUID,
) error {

	switch toType {

	case domain.LocationStorage, domain.LocationUser:
		if toID == nil {
			err := errors.New("location_id required")
			return app_errors.InvalidInput("location_id required", err)
		}

	case domain.LocationTransit:
		toID = nil

	default:
		err := errors.New("invalid location type")
		return app_errors.InvalidInput("invalid location type", err)
	}

	loc := &domain.ItemLocation{
		ItemID:       itemID,
		LocationType: toType,
		LocationID:   toID,
	}

	return s.repo.Upsert(loc)
}

func (s *LocationService) MoveToUser(itemID, userID uuid.UUID) error {
	return s.MoveItem(itemID, domain.LocationUser, &userID)
}

func (s *LocationService) MoveToStorage(itemID, storageID uuid.UUID) error {
	return s.MoveItem(itemID, domain.LocationStorage, &storageID)
}

func (s *LocationService) MoveToTransit(itemID uuid.UUID) error {
	return s.MoveItem(itemID, domain.LocationTransit, nil)
}