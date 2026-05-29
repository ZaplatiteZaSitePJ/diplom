package handlers

import (
	"encoding/json"
	"inno-accounting/internal/controllers/middleware"
	"inno-accounting/internal/dto"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/configure_headers"
	custom_errors "inno-accounting/pkg/server_utils/errors"
	"inno-accounting/pkg/server_utils/response_message"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handlers) CreateSoftware(w http.ResponseWriter, r *http.Request) {
	var input dto.SoftwareItemPublic

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid json")
		return
	}

	software, err := h.Software.CreateSoftware(&input)
	if err != nil {
		response_message.WrapperResponseJSON(w, 500, err.Error())
		return
	}

	response_message.WrapperResponseJSON(w, 201, software)
}

func (h *Handlers) GetAllSoftware(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	query := r.URL.Query()
	filter := &dto.SoftwareFilter{}

	if v := query.Get("id"); v != "" {
		filter.ID = &v
	}
	if v := query.Get("category"); v != "" {
		filter.Category = &v
	}
	if v := query.Get("last_worker_email"); v != "" {
		filter.LastWorkerEmail = &v
	}
	if v := query.Get("last_storage"); v != "" {
		filter.LastStorage = &v
	}
	if v := query.Get("transfer_status"); v != "" {
		filter.TransferStatus = &v
	}
	if v := query.Get("vendor"); v != "" {
		filter.Vendor = &v
	}
	if v := query.Get("title"); v != "" {
		filter.Title = &v
	}
	if v := query.Get("license_key"); v != "" {
		filter.LicenseKey = &v
	}

	// float
	if v := query.Get("purchase_price"); v != "" {
		if parsed, err := strconv.ParseFloat(v, 64); err == nil {
			filter.PurchasePrice = &parsed
		}
	}

	// дата строкой (как у тебя в repo)
	if v := query.Get("expired_at"); v != "" {
		filter.ExpiredAt = &v
	}

	items, err := h.Software.FindAllSoftware(filter)
	if err != nil {
		response_message.WrapperResponseJSON(w, 500, err.Error())
		return
	}

	if items == nil {
		items = []*dto.SoftwareItemPublic{}
	}

	response_message.WrapperResponseJSON(w, 200, items)
}

// ME variation
func (h *Handlers) GetMySoftware(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	authData := middleware.GetAuthData(r.Context())

	filter := &dto.SoftwareFilter{
		UserID: &authData.UserID,
	}

	software, err := h.Software.FindAllSoftware(filter)
	if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
		return
	}

	if software == nil {
		software = []*dto.SoftwareItemPublic{}
	}

	response_message.WrapperResponseJSON(w, 200, software)
}

func (h *Handlers) GetSoftwareByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid uuid")
		return
	}

	item, err := h.Software.FindSoftwareByID(id)
	if err != nil {
		response_message.WrapperResponseJSON(w, 500, err.Error())
		return
	}

	response_message.WrapperResponseJSON(w, 200, item)
}

func (h *Handlers) PatchSoftwareByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid uuid")
		return
	}

	var input dto.SoftwareItemPublic
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid json")
		return
	}

	updated, err := h.Software.ChangeSoftwareByID(id, &input)
	if err != nil {
		response_message.WrapperResponseJSON(w, 500, err.Error())
		return
	}

	response_message.WrapperResponseJSON(w, 200, updated)
}
