package software

import (
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"
	"inno-accounting/internal/use-cases/user"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	custom_errors "inno-accounting/pkg/server_utils/errors"

	"github.com/google/uuid"
)

type SoftwareService struct {
	repo        SoftwareRepository
	userService *user.UserService
}

func New(repo SoftwareRepository, userService *user.UserService) *SoftwareService {
	return &SoftwareService{
		repo:        repo,
		userService: userService,
	}
}

// CreateSoftware
func (s *SoftwareService) CreateSoftware(input *dto.SoftwareItemPublic) (*domain.Software, error) {
	logger.Info("Creating software: ", input)

	var categoryID *int
	if input.Category != "" {
		id, err := s.repo.FindCategoryIDByName(input.Category)
		if err != nil {
			return nil, app_errors.Unprocessable("Category does not exist", err)
		}
		categoryID = &id
	}

	var userID *uuid.UUID
	if input.LastWorkerEmail != nil {
		user, err := s.userService.FindUserByEmail(*input.LastWorkerEmail)
		if err != nil {
			return nil, app_errors.Unprocessable("Worker does not exist", err)
		}
		userID = &user.ID
	}

	var purchasePrice *float64
	if input.PurchasePrice != 0 {
		purchasePrice = &input.PurchasePrice
	}

	id := uuid.New()

	newSoftware := domain.Software{
	Item: domain.Item{
		ID:            id,
		UniversalName: input.Vendor + " " + input.Title,
		TypeID:        1,
		CategoryID:    categoryID,
		LastWorkerID:  userID,
		TransferStatus: "worker",
		PurchasePrice: purchasePrice,
	},
	ItemID:     id,
	Vendor:     input.Vendor,
	LicenseKey: input.LicenseKey,
	Title:      input.Title,
	StartedAt:  *input.StartedAt,
	ExpiredAt:  *input.ExpiredAt,
}

	return s.repo.Save(&newSoftware)
}

// FindSoftwareByID
func (s *SoftwareService) FindSoftwareByID(id uuid.UUID) (*dto.SoftwareItemPublic, error) {
	software, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	var category string
	if software.CategoryID != nil {
		cat, _ := s.repo.FindCategoryNameByID(*software.CategoryID)
		if cat != nil {
			category = *cat
		}
	}

	var lastWorkerEmail *string
	if software.LastWorkerID != nil {
		user, err := s.userService.FindUserByID(*software.LastWorkerID)
		if err == nil && user != nil {
			lastWorkerEmail = &user.Email
		}
	}

	var purchasePrice float64
	if software.PurchasePrice != nil {
		purchasePrice = *software.PurchasePrice
	}

	return &dto.SoftwareItemPublic{
		ID:              software.ID,
		UniversalName:   software.UniversalName,
		Category:        category,
		LastWorkerEmail: lastWorkerEmail,
		PurchasePrice:   purchasePrice,

		Vendor:     software.Vendor,
		LicenseKey: software.LicenseKey,
		Title:      software.Title,
		StartedAt:  &software.StartedAt,
		ExpiredAt:  &software.ExpiredAt,
		UpdatedAt:  &software.UpdatedAt,
	}, nil
}

// FindAll
func (s *SoftwareService) FindAllSoftware(filter *dto.SoftwareFilter) ([]*dto.SoftwareItemPublic, error) {
	items, err := s.repo.FindAll(filter)
	if err != nil {
		wErr := custom_errors.New(err, 500)
		return nil, wErr
	}

	if items == nil {
		return []*dto.SoftwareItemPublic{}, nil
	}

	return items, nil
}

// Patch
func (s *SoftwareService) ChangeSoftwareByID(id uuid.UUID, input *dto.SoftwareItemPublic) (*domain.Software, error) {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if input.Title != "" {
		existing.Title = input.Title
	}

	if input.Vendor != "" {
		existing.Vendor = input.Vendor
	}

	if input.LicenseKey != "" {
		existing.LicenseKey = input.LicenseKey
	}

	if input.StartedAt != nil {
		existing.StartedAt = *input.StartedAt
	}

	if input.ExpiredAt != nil {
		existing.ExpiredAt = *input.ExpiredAt
	}

	if input.Category != "" {
		id, err := s.repo.FindCategoryIDByName(input.Category)
		if err != nil {
			return nil, app_errors.Unprocessable("Category does not exist", err)
		}
		existing.CategoryID = &id
	}

	if input.LastWorkerEmail != nil {
		user, err := s.userService.FindUserByEmail(*input.LastWorkerEmail)
		if err != nil {
			return nil, app_errors.Unprocessable("Worker does not exist", err)
		}
		existing.LastWorkerID = &user.ID
	}

	if input.PurchasePrice != 0 {
		existing.PurchasePrice = &input.PurchasePrice
	}

	existing.UniversalName = existing.Vendor + " " + existing.Title

	return s.repo.Change(id, existing)
}