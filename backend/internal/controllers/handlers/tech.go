package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"inno-accounting/internal/dto"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	"inno-accounting/pkg/server_utils/configure_headers"
	"inno-accounting/pkg/server_utils/response_message"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// CreateTech creates new tech object
func (h *Handlers) CreateTech(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	var input dto.TechItemPublic

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid json body")
		return
	}

	if input.Brand == "" || input.Model == "" {
		errStr := "brand and model are required"
		jsonErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(
			w,
			app_errors.HTTPStatusFromCode(jsonErr.Code),
			jsonErr.Message,
		)
		return
	}

	tech, err := h.TechItems.CreateTech(&input)
	if err != nil {
		var appErr *app_errors.AppError
		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(
				w,
				app_errors.HTTPStatusFromCode(appErr.Code),
				appErr.Message,
			)
			return
		}

		logger.Error("CreateTech error:", err)
		response_message.WrapperResponseJSON(w, 500, "internal server error")
		return
	}

	logger.Info(fmt.Sprintf("Tech created successfully: %+v", tech))
	response_message.WrapperResponseJSON(w, 201, tech)
}

// GetAllTech returns filtered list of tech objects
func (h *Handlers) GetAllTech(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	query := r.URL.Query()

	filter := &dto.TechFilter{}

	if v := query.Get("id"); v != "" {
		filter.ID = &v
	}
	if v := query.Get("brand"); v != "" {
		filter.Brand = &v
	}
	if v := query.Get("model"); v != "" {
		filter.Model = &v
	}
	if v := query.Get("last_worker"); v != "" {
		filter.LastWorker = &v
	}
	if v := query.Get("last_storage"); v != "" {
		filter.LastStorage = &v
	}
	if v := query.Get("category"); v != "" {
		filter.Category = &v
	}
	if v := query.Get("quality_status"); v != "" {
		filter.QualityStatus = &v
	}
	if v := query.Get("transfer_status"); v != "" {
		filter.TransferStatus = &v
	}

	techs, err := h.TechItems.FindAllTechs(filter)
	if err != nil {
		var appErr *app_errors.AppError
		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(
				w,
				app_errors.HTTPStatusFromCode(appErr.Code),
				appErr.Message,
			)
		} else {
			logger.Error("Unhandled error in GetAllTech:", err)
			response_message.WrapperResponseJSON(w, 500, "internal server error")
		}
		return
	}

	if techs == nil {
		techs = []*dto.TechItemPublic{}
	}

	response_message.WrapperResponseJSON(w, 200, techs)
}

// GetTechByID returns single tech object by UUID
func (h *Handlers) GetTechByID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	// ожидаем ?id=uuid
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		errStr := "id is required"
		jsonErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(
			w,
			app_errors.HTTPStatusFromCode(jsonErr.Code),
			jsonErr.Message,
		)
		return
	}

	techID, err := uuid.Parse(idStr)
	if err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid uuid format")
		return
	}

	tech, err := h.TechItems.FindTechByID(techID)
	if err != nil {
		var appErr *app_errors.AppError
		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(
				w,
				app_errors.HTTPStatusFromCode(appErr.Code),
				appErr.Message,
			)
			return
		}

		logger.Error("GetTechByID error:", err)
		response_message.WrapperResponseJSON(w, 500, "internal server error")
		return
	}

	response_message.WrapperResponseJSON(w, 200, tech)
}

// PatchTechByID partially updates tech item
func (h *Handlers) PatchTechByID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		errStr := "id is required"
		jsonErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(
			w,
			app_errors.HTTPStatusFromCode(jsonErr.Code),
			jsonErr.Message,
		)
		return
	}

	techID, err := uuid.Parse(idStr)
	if err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid uuid format")
		return
	}

	var input dto.TechItemPublic
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid json body")
		return
	}

	updatedTech, err := h.TechItems.ChangeTechByID(techID, &input)
	if err != nil {
		var appErr *app_errors.AppError
		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(
				w,
				app_errors.HTTPStatusFromCode(appErr.Code),
				appErr.Message,
			)
			return
		}

		logger.Error("PatchTechByID error:", err)
		response_message.WrapperResponseJSON(w, 500, "internal server error")
		return
	}

	response_message.WrapperResponseJSON(w, 200, updatedTech)
}

// DeleteTechByID deletes tech item by UUID
func (h *Handlers) DeleteTechByID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		errStr := "id is required"
		jsonErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(
			w,
			app_errors.HTTPStatusFromCode(jsonErr.Code),
			jsonErr.Message,
		)
		return
	}

	techID, err := uuid.Parse(idStr)
	if err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid uuid format")
		return
	}

	err = h.TechItems.DeleteTechByID(techID)
	if err != nil {
		var appErr *app_errors.AppError
		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(
				w,
				app_errors.HTTPStatusFromCode(appErr.Code),
				appErr.Message,
			)
			return
		}

		logger.Error("DeleteTechByID error:", err)
		response_message.WrapperResponseJSON(w, 500, "internal server error")
		return
	}

	logger.Info(fmt.Sprintf("Tech deleted successfully: %s", techID))
	response_message.WrapperResponseJSON(w, 204, "tech deleted successfully")
}

func (h *Handlers) GetCategoriesByTypeID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	vars := mux.Vars(r)
	typeStr := vars["type_index"]

	if typeStr == "" {
		response_message.WrapperResponseJSON(w, 400, "type_index is required")
		return
	}

	typeID, err := strconv.Atoi(typeStr)
	if err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid type_index")
		return
	}

	categories, err := h.TechItems.GetCategoriesByTypeID(typeID)
	if err != nil {
		var appErr *app_errors.AppError
		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(
				w,
				app_errors.HTTPStatusFromCode(appErr.Code),
				appErr.Message,
			)
			return
		}

		logger.Error("GetCategoriesByTypeID error:", err)
		response_message.WrapperResponseJSON(w, 500, "internal server error")
		return
	}

	response_message.WrapperResponseJSON(w, 200, categories)
}