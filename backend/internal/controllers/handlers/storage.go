package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"inno-accounting/internal/dto"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	"inno-accounting/pkg/server_utils/configure_headers"
	custom_errors "inno-accounting/pkg/server_utils/errors"
	"inno-accounting/pkg/server_utils/response_message"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// CreateStorage handles HTTP request for creating a new storage entity.
// Validates input DTO and returns created storage or error response.
func (h *Handlers) CreateStorage(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	var input dto.CreateStorage
	json.NewDecoder(r.Body).Decode(&input)

	if input.StorageName == "" || input.City == "" || input.Capacity == 0 {
		errStr := "one or more required fields are missing"
		jsonErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(jsonErr.Code), jsonErr.Message)
		return
	}

	storage, err := h.Storage.CreateStorage(&input)

	if err != nil {
		var appErr *app_errors.AppError

		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(appErr.Code), appErr.Message)
		}
		return
	}

	logger.Info(fmt.Sprintf("Storage created successfully: %+v", storage))
	response_message.WrapperResponseJSON(w, 201, storage)
}

// GetStorageByID handles HTTP request for retrieving a storage by its UUID.
// Parses ID from request path and returns storage or error response.
func (h *Handlers) GetStorageByID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	idStr := mux.Vars(r)["id"]
	storageUUID, err := uuid.Parse(idStr)

	if err != nil {
		errStr := "invalid storage ID"
		idErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(idErr.Code), idErr.Message)
		return
	}

	storage, err := h.Storage.FindStorageByID(storageUUID)

	if err != nil {
		var appErr *app_errors.AppError

		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(appErr.Code), appErr.Message)
		}
		return
	}

	logger.Info("Storage found successfully")
	response_message.WrapperResponseJSON(w, 200, storage)
}

// GetAllStorages handles HTTP request for retrieving all storages.
// Returns a list of storages or error response.
func (h *Handlers) GetAllStorages(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	storages, err := h.Storage.FindAllStorages()
	if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
		return
	}

	logger.Info("Storages found successfully")
	response_message.WrapperResponseJSON(w, 200, storages)
}

// DeleteStorageByID handles HTTP request for soft-deleting a storage by UUID.
// Marks storage as deleted and returns status response.
func (h *Handlers) DeleteStorageByID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	idStr := mux.Vars(r)["id"]
	storageUUID, err := uuid.Parse(idStr)
	if err != nil {
		custom_errors.ErrorResponse(w, custom_errors.New(err, 400), logger.GetLoger())
		return
	}

	newStorageName := r.URL.Query().Get("newStorageName")

	var namePtr *string
	if newStorageName != "" {
		namePtr = &newStorageName
	}

	err = h.Storage.DeleteStorageByID(storageUUID, namePtr)
	if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
		return
	}

	response_message.WrapperResponseJSON(w, 200, "storage deleted successfully")
}


func (h *Handlers) UpdateStorageByID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	idStr := mux.Vars(r)["id"]
	storageUUID, err := uuid.Parse(idStr)
	if err != nil {
		errStr := "invalid storage ID"
		idErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(idErr.Code), idErr.Message)
		return
	}

	var input dto.CreateStorage
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		errStr := "invalid input data"
		jsonErr := app_errors.InvalidInput(errStr, err)
		response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(jsonErr.Code), jsonErr.Message)
		return
	}

	updatedStorage, err := h.Storage.ChangeStorageByID(storageUUID, &input)
	if err != nil {
		var appErr *app_errors.AppError
		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(appErr.Code), appErr.Message)
		} else {
			custom_errors.ErrorResponse(w, err, logger.GetLoger())
		}
		return
	}

	logger.Info(fmt.Sprintf("Storage updated successfully: %+v", updatedStorage))
	response_message.WrapperResponseJSON(w, 200, updatedStorage)
}

// GetStoragesByName handles HTTP request for retrieving storages by name.
// Returns a list of storages matching the given name.
func (h *Handlers) GetStorageByExactName(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	name := r.URL.Query().Get("name")
	if name == "" {
		err := app_errors.InvalidInput("name query parameter is required", errors.New("missing name"))
		response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(err.Code), err.Message)
		return
	}

	storage, err := h.Storage.FindStorageByExactName(name)
	if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
		return
	}

	response_message.WrapperResponseJSON(w, 200, storage)
}