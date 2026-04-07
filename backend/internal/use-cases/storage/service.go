package storages

import (
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