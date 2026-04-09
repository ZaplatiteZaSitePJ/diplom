package tech

import (
	"fmt"
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"
	storages "inno-accounting/internal/use-cases/storage"
	"inno-accounting/internal/use-cases/user"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	custom_errors "inno-accounting/pkg/server_utils/errors"

	"github.com/google/uuid"
)

type TechService struct {
	repo           TechRepository
	storageService *storages.StorageService
	userService *user.UserService
}

func New(repo TechRepository, storageService *storages.StorageService, userService *user.UserService) *TechService {
	return &TechService{
		repo:           repo,
		storageService: storageService,
		userService: userService,
	}
}
// CreateTech создаёт новый тех. объект
func (t *TechService) CreateTech(input *dto.TechItemPublic) (*domain.Tech, error) {
	logger.Info("Creating new tech: ", input)

	// --- Storage (optional)
	var storageID *uuid.UUID
	if input.LastStorage != nil {
		storage, err := t.storageService.FindStorageByExactName(*input.LastStorage)
		if err != nil {
			return nil, app_errors.Unprocessable("Storage does not exist", err)
		}
		storageID = &storage.ID
	}

	// --- User (optional)
	var userID *uuid.UUID
	if input.LastWorkerEmail != nil {
		user, err := t.userService.FindUserByEmail(*input.LastWorkerEmail)
		if err != nil {
			return nil, app_errors.Unprocessable("Worker does not exist", err)
		}
		userID = &user.ID
	}

	// --- Category (optional)
	var categoryID *int
	if input.Category != nil {
		id, err := t.repo.FindCategoryIDByName(*input.Category)
		if err != nil {
			return nil, app_errors.Unprocessable("Category does not exist", err)
		}
		categoryID = &id
	}

	// --- optional числовые поля
	var occupiedCells *int
	if input.OccupiedCells != 0 {
		occupiedCells = &input.OccupiedCells
	}

	var purchasePrice *float64
	if input.PurchasePrice != 0 {
		purchasePrice = &input.PurchasePrice
	}

	// --- quality status (дефолт)
	quality := "new"
	if input.QualityStatus != "" {
		quality = input.QualityStatus
	}

	// --- создаём объект
	id := uuid.New()

	newTech := domain.Tech{
		Item: domain.Item{
			ID:             id,
			CategoryID:     categoryID,
			UniversalName:  input.Brand + " " + input.Model,
			OccupiedCells:  occupiedCells,
			LastStorageID:  storageID,
			LastWorkerID:   userID,
			TransferStatus: input.TransferStatus,
			QualityStatus:  &quality,
			PurchasePrice:  purchasePrice,
			TypeID:         0,
		},
		Brand:             input.Brand,
		Model:             input.Model,
		WarrantyStartedAt: input.WarrantyStartedAt,
		WarrantyEndAt:     input.WarrantyEndAt,
		ItemID:            id,
	}

	// --- сохраняем
	tech, err := t.repo.Save(&newTech)
	if err != nil {
		logger.Error("Failed to create tech: ", err)
		return nil, err
	}

	logger.Info(fmt.Sprintf("Tech created successfully: %+v", tech))
	return tech, nil
}

// ChangeTechByID обновляет существующий тех. объект по ID
func (t *TechService) ChangeTechByID(techID uuid.UUID, input *dto.TechItemPublic) (*domain.Tech, error) {
	logger.Info("Trying to change tech: ", techID)

	existingTech, err := t.repo.FindByID(techID)
	if err != nil {
		return nil, err
	}

	if input.Brand != "" {
		existingTech.Brand = input.Brand
	}
	if input.Model != "" {
		existingTech.Model = input.Model
	}
	if !input.WarrantyStartedAt.IsZero() {
		existingTech.WarrantyStartedAt = input.WarrantyStartedAt
	}
	if !input.WarrantyEndAt.IsZero() {
		existingTech.WarrantyEndAt = input.WarrantyEndAt
	}

	updatedTech, err := t.repo.Change(techID, existingTech)
	if err != nil {
		logger.Error("Failed to update tech: ", err)
		return nil, err
	}

	logger.Info(fmt.Sprintf("Tech updated successfully: %+v", updatedTech))
	return updatedTech, nil
}

// FindTechByID ищет тех. объект по UUID
func (t *TechService) FindTechByID(techID uuid.UUID) (*domain.Tech, error) {
	logger.Info("Trying to find tech: ", techID)

	tech, err := t.repo.FindByID(techID)
	if err != nil {
		return nil, err
	}

	return tech, nil
}

// FindAllTechs возвращает все тех. объекты с возможностью фильтрации
func (t *TechService) FindAllTechs(filter *dto.TechFilter) ([]*domain.Tech, error) {
	logger.Info("Trying to find all techs with filter: ", filter)

	techs, err := t.repo.FindAll(filter)
	if err != nil {
		wErr := custom_errors.New(err, 500)
		wErr.AddResponseData("Internal server error")
		wErr.AddLogData(err.Error())
		return nil, wErr
	}

	// даже если пустой список — это не ошибка
	if techs == nil {
		techs = []*domain.Tech{}
	}

	return techs, nil
}

// DeleteTechByID мягко удаляет тех. объект по UUID
func (t *TechService) DeleteTechByID(techID uuid.UUID) error {
	logger.Info("Trying to delete tech: ", techID)

	err := t.repo.DeleteByID(techID)
	if err != nil {
		logger.Error("Failed to delete tech: ", err)
		return err
	}

	logger.Info(fmt.Sprintf("Tech deleted successfully: %v", techID))
	return nil
}