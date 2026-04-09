package storages

import (
	"fmt"
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"
	"inno-accounting/pkg/logger"
	custom_errors "inno-accounting/pkg/server_utils/errors"
	"time"

	"github.com/google/uuid"
)

type StorageService struct {
	repo StorageRepository
}

func New(repo StorageRepository) *StorageService {
	return &StorageService{
		repo: repo,
	}
}


func (u *StorageService) CreateStorage(input *dto.CreateStorage) (*domain.Storage, error){
	logger.Info("Creating new storage: ", input)
	
	newStorage := domain.Storage{
		StorageName: input.StorageName,
		City: input.City,
		Capacity: input.Capacity,
		OccupiedCells: 0,
		ItemsAmount: 0,
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return u.repo.Save(&newStorage)
}

func (s *StorageService) ChangeStorageByID(storageID uuid.UUID, input *dto.CreateStorage) (*domain.Storage, error) {
	logger.Info("Trying to change storage: ", storageID)

	// Найдем существующее хранилище
	existingStorage, err := s.repo.FindByID(storageID)
	if err != nil {
		return nil, err
	}

	if input.StorageName != "" {
		existingStorage.StorageName = input.StorageName
	}
	if input.City != "" {
		existingStorage.City = input.City
	}
	if input.Capacity != 0 {
		existingStorage.Capacity = input.Capacity
	}

	existingStorage.UpdatedAt = time.Now()

	updatedStorage, err := s.repo.Change(storageID, existingStorage)
	if err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf("Storage updated successfully: %+v", updatedStorage))
	return updatedStorage, nil
}

func (s *StorageService) FindStorageByID(storageID uuid.UUID) (*domain.Storage,  error){
	logger.Info("Trying to find storage: ", storageID)

	findedStorage, err := s.repo.FindByID(storageID)
	if err != nil {
	 	return nil,  err
	}

	return findedStorage,  nil
}

func (u *StorageService) FindAllStorages() ([]*domain.Storage, error) {
	logger.Info("Trying to find all storages")
	findedStorages, err := u.repo.FindAll()

	if err != nil {
	 	wErr := custom_errors.New(err, 500)
	 	wErr.AddResponseData("Internal server error")
	 	wErr.AddLogData(err.Error())
		return nil, wErr
	}
	return findedStorages, nil
}

func (s *StorageService) DeleteStorageByID(storageID uuid.UUID) error {
	return nil
}

func (s *StorageService) FindStorageByExactName(name string) (*domain.Storage, error) {
	if name == "" {
		wErr := custom_errors.New(nil, 400)
		wErr.AddResponseData("Name parameter is required")
		wErr.AddLogData("Empty name parameter provided")
		return nil, wErr
	}

	storage, err := s.repo.FindByExactName(name)
	if err != nil {
		return nil, custom_errors.New(err, 500)
	}

	if storage == nil {
		wErr := custom_errors.New(nil, 404)
		wErr.AddResponseData(fmt.Sprintf("Storage with name '%s' not found", name))
		return nil, wErr
	}

	return storage, nil
}