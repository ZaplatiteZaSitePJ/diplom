package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"inno-accounting/internal/domain"
	"inno-accounting/internal/dto"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	"inno-accounting/pkg/server_utils/configure_headers"
	"inno-accounting/pkg/server_utils/response_message"
	"net/http"
)

// CreateTech handles HTTP request for creating a new tech entity.
// Validates input DTO and returns created tech or error response.
func (h *Handlers) CreateTech(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	var input dto.TechItemPublic
	json.NewDecoder(r.Body).Decode(&input)

	if input.Brand == "" || input.Model == "" {
		errStr := "one or more required fields are missing"
		jsonErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(jsonErr.Code), jsonErr.Message)
		return
	}

	tech, err := h.TechItems.CreateTech(&input)

	if err != nil {
		var appErr *app_errors.AppError

		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(appErr.Code), appErr.Message)
		}
		return
	}

	logger.Info(fmt.Sprintf("Tech created successfully: %+v", tech))
	response_message.WrapperResponseJSON(w, 201, tech)
}

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
		techs = []*domain.Tech{}
	}

	response_message.WrapperResponseJSON(w, 200, techs)
}