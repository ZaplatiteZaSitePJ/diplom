package tech

import (
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
	userService    *user.UserService
}

func New(repo TechRepository, storageService *storages.StorageService, userService *user.UserService) *TechService {
	return &TechService{
		repo:           repo,
		storageService: storageService,
		userService:    userService,
	}
}

// CreateTech создаёт новый тех. объект
func (t *TechService) CreateTech(input *dto.TechItemPublic) (*domain.Tech, error) {
	logger.Info("Creating new tech: ", input)

	var storageID *uuid.UUID
	if input.LastStorage != nil {
		storage, err := t.storageService.FindStorageByExactName(*input.LastStorage)
		if err != nil {
			return nil, app_errors.Unprocessable("Storage does not exist", err)
		}
		storageID = &storage.ID
	}

	var userID *uuid.UUID
	if input.LastWorkerEmail != nil {
		user, err := t.userService.FindUserByEmail(*input.LastWorkerEmail)
		if err != nil {
			return nil, app_errors.Unprocessable("Worker does not exist", err)
		}
		userID = &user.ID
	}

	var categoryID *int
	if input.Category != nil {
		id, err := t.repo.FindCategoryIDByName(*input.Category)
		if err != nil {
			return nil, app_errors.Unprocessable("Category does not exist", err)
		}
		categoryID = &id
	}

	var occupiedCells *int
	if input.OccupiedCells != 0 {
		occupiedCells = &input.OccupiedCells
	}

	var purchasePrice *float64
	if input.PurchasePrice != 0 {
		purchasePrice = &input.PurchasePrice
	}

	quality := "new"
	if input.QualityStatus != "" {
		quality = input.QualityStatus
	}

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
			PostNumber: input.PostNumber,
			MovementFrom: input.MovementFrom,
			MovementTo:   input.MovementTo,
			SendedAt:     input.SendedAt,
			ArrivedAt:    input.ArrivedAt,
			IsActual:     input.IsActual,
		},
		Brand:             input.Brand,
		Model:             input.Model,
		WarrantyStartedAt: input.WarrantyStartedAt,
		WarrantyEndAt:     input.WarrantyEndAt,
		ItemID:            id,
	}

	return t.repo.Save(&newTech)
}

// FindTechByID возвращает DTO с join данными
func (t *TechService) FindTechByID(techID uuid.UUID) (*dto.TechItemPublic, error) {
	logger.Info("Trying to find tech: ", techID)

	tech, err := t.repo.FindByID(techID)
	if err != nil {
		return nil, err
	}

	// --- category (nullable)
	var category *string
	if tech.CategoryID != nil {
		cat, err := t.repo.FindCategoryNameByID(*tech.CategoryID)
		if err == nil {
			category = cat
		}
	}

	// --- storage (nullable)
	var lastStorage *string
	if tech.LastStorageID != nil {
		storage, err := t.storageService.FindStorageByID(*tech.LastStorageID)
		if err == nil && storage != nil {
			lastStorage = &storage.StorageName
		}
	}

	// --- user (nullable)
	var lastWorkerEmail *string
	if tech.LastWorkerID != nil {
		user, err := t.userService.FindUserByID(*tech.LastWorkerID)
		if err == nil && user != nil {
			lastWorkerEmail = &user.Email
		}
	}

	// --- safe values
	var qualityStatus string
	if tech.QualityStatus != nil {
		qualityStatus = *tech.QualityStatus
	}

	var purchasePrice float64
	if tech.PurchasePrice != nil {
		purchasePrice = *tech.PurchasePrice
	}

	var occupiedCells int
	if tech.OccupiedCells != nil {
		occupiedCells = *tech.OccupiedCells
	}

	return &dto.TechItemPublic{
		ID:                tech.ID,
		Type_ID:           tech.TypeID,
		Category:          category,
		LastStorage:       lastStorage,
		LastWorkerEmail:   lastWorkerEmail,
		TransferStatus:    tech.TransferStatus,
		QualityStatus:     qualityStatus,
		PurchasePrice:     purchasePrice,
		OccupiedCells:     occupiedCells,
		Brand:             tech.Brand,
		Model:             tech.Model,
		WarrantyStartedAt: tech.WarrantyStartedAt,
		WarrantyEndAt:     tech.WarrantyEndAt,
		UniversalName:     tech.UniversalName,
		LastStorageID: tech.LastStorageID,
		PostNumber: tech.PostNumber,
		SendedAt: tech.SendedAt,
		ArrivedAt: tech.ArrivedAt,
		MovementFrom: tech.MovementFrom,
		MovementTo: tech.MovementTo,
	}, nil
}

// FindAllTechs
func (t *TechService) FindAllTechs(filter *dto.TechFilter) ([]*dto.TechItemPublic, error) {
	logger.Info("Trying to find all techs with filter: ", filter)

	techs, err := t.repo.FindAll(filter)
	if err != nil {
		wErr := custom_errors.New(err, 500)
		wErr.AddResponseData("Internal server error")
		wErr.AddLogData(err.Error())
		return nil, wErr
	}

	if techs == nil {
		return []*dto.TechItemPublic{}, nil
	}

	return techs, nil
}

// DeleteTechByID
func (t *TechService) DeleteTechByID(techID uuid.UUID) error {
	logger.Info("Trying to delete tech: ", techID)

	err := t.repo.DeleteByID(techID)
	if err != nil {
		logger.Error("Failed to delete tech: ", err)
		return err
	}

	return nil
}

func (t *TechService) ChangeTechByID(techID uuid.UUID, input *dto.TechItemPublic) (*domain.Tech, error) {
	logger.Info("Trying to patch tech: ", techID)

	existingTech, err := t.repo.FindByID(techID)
	if err != nil {
		return nil, err
	}

	// --- simple fields
	if input.Brand != "" {
		existingTech.Brand = input.Brand
	}

	if input.Model != "" {
		existingTech.Model = input.Model
	}

	if input.WarrantyStartedAt != nil && !input.WarrantyStartedAt.IsZero() {
		existingTech.WarrantyStartedAt = input.WarrantyStartedAt
	}

	if input.WarrantyEndAt != nil && !input.WarrantyEndAt.IsZero() {
		existingTech.WarrantyEndAt = input.WarrantyEndAt
	}

	// --- category
	if input.Category != nil {
		id, err := t.repo.FindCategoryIDByName(*input.Category)
		if err != nil {
			return nil, app_errors.Unprocessable("Category does not exist", err)
		}
		existingTech.CategoryID = &id
	}

	// --- storage
	if input.LastStorage != nil {
		storage, err := t.storageService.FindStorageByExactName(*input.LastStorage)
		if err != nil {
			return nil, app_errors.Unprocessable("Storage does not exist", err)
		}
		existingTech.LastStorageID = &storage.ID
	}

	// --- user
	if input.LastWorkerEmail != nil {
		user, err := t.userService.FindUserByEmail(*input.LastWorkerEmail)
		if err != nil {
			return nil, app_errors.Unprocessable("Worker does not exist", err)
		}
		existingTech.LastWorkerID = &user.ID
	}

	// --- optional numeric
	if input.OccupiedCells != 0 {
		existingTech.OccupiedCells = &input.OccupiedCells
	}

	if input.PurchasePrice != 0 {
		existingTech.PurchasePrice = &input.PurchasePrice
	}

	// --- movement
	if input.PostNumber != nil && *input.PostNumber != "" {
		existingTech.PostNumber = input.PostNumber
	}

	if input.MovementFrom != nil {
		existingTech.MovementFrom = input.MovementFrom
	}

	if input.MovementTo != nil {
		existingTech.MovementTo = input.MovementTo
	}

	if input.SendedAt != nil && !input.SendedAt.IsZero() {
		existingTech.SendedAt = input.SendedAt
	}

	if input.ArrivedAt != nil && !input.ArrivedAt.IsZero() {
		existingTech.ArrivedAt = input.ArrivedAt
	}

	if input.IsActual != nil {
		existingTech.IsActual = input.IsActual
	}

	// --- quality
	if input.QualityStatus != "" {
		existingTech.QualityStatus = &input.QualityStatus
	}

	// --- transfer status
	if input.TransferStatus != "" {
		existingTech.TransferStatus = input.TransferStatus
	}

	// --- universal name
	existingTech.UniversalName = existingTech.Brand + " " + existingTech.Model

	return t.repo.Change(techID, existingTech)
}

func (t *TechService) GetCategoriesByTypeID(typeID int) ([]string, error) {
	logger.Info("Getting categories by typeID:", typeID)

	categories, err := t.repo.GetCategoriesByTypeID(typeID)
	if err != nil {
		return nil, err
	}

	if categories == nil {
		return []string{}, nil
	}

	return categories, nil
}

func (t *TechService) FindTechByUserID(userID uuid.UUID) ([]*dto.TechItemPublic, error) {
	logger.Info("Finding tech by userID:", userID)

	// проверим что пользователь существует
	_, err := t.userService.FindUserByID(userID)
	if err != nil {
		return nil, app_errors.Unprocessable("user does not exist", err)
	}

	techs, err := t.repo.FindByUserID(userID)
	if err != nil {
		wErr := custom_errors.New(err, 500)
		wErr.AddResponseData("Internal server error")
		wErr.AddLogData(err.Error())
		return nil, wErr
	}

	if techs == nil {
		return []*dto.TechItemPublic{}, nil
	}

	return techs, nil
}