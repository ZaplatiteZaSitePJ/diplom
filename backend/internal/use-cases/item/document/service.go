package document

import (
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"
	storages "inno-accounting/internal/use-cases/storage"
	"inno-accounting/internal/use-cases/user"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	custom_errors "inno-accounting/pkg/server_utils/errors"
	"time"

	"github.com/google/uuid"
)

type DocumentService struct {
	repo           DocumentRepository
	storageService *storages.StorageService
	userService    *user.UserService
}

func New(repo DocumentRepository, storageService *storages.StorageService, userService *user.UserService) *DocumentService {
	return &DocumentService{
		repo:           repo,
		storageService: storageService,
		userService:    userService,
	}
}

// ===================== CREATE =====================

func (s *DocumentService) CreateDocument(input *dto.DocsItemPublic) (*domain.Document, error) {
	logger.Info("Creating document:", input)

	// --- storage
	var storageID *uuid.UUID
	if input.LastStorage != nil {
		storage, err := s.storageService.FindStorageByExactName(*input.LastStorage)
		if err != nil {
			return nil, app_errors.Unprocessable("Storage does not exist", err)
		}
		storageID = &storage.ID
	}

	// --- worker (LAST WORKER EMAIL)
	var workerID *uuid.UUID
	if input.LastWorkerEmail != nil {
		user, err := s.userService.FindUserByEmail(*input.LastWorkerEmail)
		if err != nil {
			return nil, app_errors.Unprocessable("Worker does not exist", err)
		}
		workerID = &user.ID
	}

	// --- responsible worker
	responsibleUser, err := s.userService.FindUserByEmail(input.ResponsibleWorkerEmail)
	if err != nil {
		return nil, app_errors.Unprocessable("Responsible worker does not exist", err)
	}

	// --- category
	var categoryID *int
	if input.Category != "" {
		id, err := s.repo.FindCategoryIDByName(input.Category)
		if err != nil {
			return nil, app_errors.Unprocessable("Category does not exist", err)
		}
		categoryID = &id
	}

	// --- transfer status
	var transferStatus string
	if input.TransferStatus != nil {
		transferStatus = *input.TransferStatus
	}

	qualityDefault := "new" 

	id := uuid.New()

	doc := domain.Document{
		Item: domain.Item{
			ID:             id,
			UniversalName:  input.Category + " №" + input.DocNumber,
			CategoryID:     categoryID,
			LastStorageID:  storageID,
			LastWorkerID:   workerID,
			TransferStatus: transferStatus,
			TypeID:         1,
			QualityStatus: &qualityDefault,
		},
		ItemID:                 id,
		ResponsibleWorkerID:    responsibleUser.ID,
		ResponsibleWorkerEmail: input.ResponsibleWorkerEmail,
		FullSignedAt:           timeValue(input.FullSignedAt),
		NeededSigns:            input.NeededSigns,
		ReceivedSigns:          input.ReceivedSigns,
		DocNumber:              input.DocNumber,
	}

	return s.repo.Save(&doc)
}

// ===================== GET BY ID =====================

func (s *DocumentService) FindDocumentByID(id uuid.UUID) (*dto.DocsItemPublic, error) {
	doc, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// --- category
	var category string
	if doc.CategoryID != nil {
		c, _ := s.repo.FindCategoryNameByID(*doc.CategoryID)
		if c != nil {
			category = *c
		}
	}

	// --- storage
	var lastStorage *string
	if doc.LastStorageID != nil {
		sr, err := s.storageService.FindStorageByID(*doc.LastStorageID)
		if err == nil && sr != nil {
			lastStorage = &sr.StorageName
		}
	}

	// --- LAST WORKER EMAIL (FIXED)
	var lastWorkerEmail *string
	if doc.LastWorkerID != nil {
		u, err := s.userService.FindUserByID(*doc.LastWorkerID)
		if err == nil && u != nil {
			lastWorkerEmail = &u.Email
		}
	}

	// --- transfer status
	var transferStatus *string
	if doc.TransferStatus != "" {
		transferStatus = &doc.TransferStatus
	}

	return &dto.DocsItemPublic{
		ID:                     doc.ID,
		UniversalName:         doc.UniversalName,
		Type:                  "docs",
		Category:              category,
		LastStorage:           lastStorage,
		LastWorkerEmail:       lastWorkerEmail,
		TransferStatus:        transferStatus,

		FullSignedAt:           &doc.FullSignedAt,
		ResponsibleWorkerEmail: doc.ResponsibleWorkerEmail,
		NeededSigns:            doc.NeededSigns,
		ReceivedSigns:          doc.ReceivedSigns,
		DocNumber:              doc.DocNumber,
	}, nil
}

// ===================== GET ALL =====================

func (s *DocumentService) FindAllDocuments(filter *dto.DocsFilter) ([]*dto.DocsItemPublic, error) {
	items, err := s.repo.FindAll(filter)
	if err != nil {
		return nil, custom_errors.New(err, 500)
	}

	if items == nil {
		return []*dto.DocsItemPublic{}, nil
	}

	return items, nil
}

// ===================== PATCH =====================

func (s *DocumentService) ChangeDocumentByID(id uuid.UUID, input *dto.DocsItemPublic) (*domain.Document, error) {
	doc, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// --- doc number
	if input.DocNumber != "" {
		doc.DocNumber = input.DocNumber
		doc.UniversalName = input.DocNumber
	}

	// --- signed date
	if input.FullSignedAt != nil {
		doc.FullSignedAt = *input.FullSignedAt
	}

	if input.TransferStatus != nil {
		doc.TransferStatus = *input.TransferStatus
	}

	// --- storage
	if input.LastStorage != nil {
		storage, err := s.storageService.FindStorageByExactName(*input.LastStorage)
		if err != nil {
			return nil, app_errors.Unprocessable("Storage does not exist", err)
		}
		doc.LastStorageID = &storage.ID
	}

	// --- last worker
	if input.LastWorkerEmail != nil {
		user, err := s.userService.FindUserByEmail(*input.LastWorkerEmail)
		if err != nil {
			return nil, app_errors.Unprocessable("Worker does not exist", err)
		}
		doc.LastWorkerID = &user.ID
	}

	// --- category
	var categoryName string

	if input.Category != "" {
		id, err := s.repo.FindCategoryIDByName(input.Category)
		if err != nil {
			return nil, app_errors.Unprocessable("Category does not exist", err)
		}
		doc.CategoryID = &id
		categoryName = input.Category
	} else if doc.CategoryID != nil {
		c, _ := s.repo.FindCategoryNameByID(*doc.CategoryID)
		if c != nil {
			categoryName = *c
		}
	}
	
	if input.DocNumber != "" {
		doc.DocNumber = input.DocNumber
	}

	universalName := categoryName + " №" + input.DocNumber

	// --- flags
	doc.NeededSigns = input.NeededSigns
	doc.ReceivedSigns = input.ReceivedSigns

	// --- responsible worker
	if input.ResponsibleWorkerEmail != "" {
		user, err := s.userService.FindUserByEmail(input.ResponsibleWorkerEmail)
		if err != nil {
			return nil, app_errors.Unprocessable("Worker does not exist", err)
		}
		doc.ResponsibleWorkerID = user.ID
		doc.ResponsibleWorkerEmail = user.Email
		doc.UniversalName = universalName
	}

	return s.repo.Change(id, doc)
}

// ===================== HELPERS =====================

func timeValue(t *time.Time) time.Time {
	if t != nil {
		return *t
	}
	return time.Time{}
}