package handlers

import (
	"encoding/json"
	"inno-accounting/internal/dto"
	"inno-accounting/pkg/server_utils/response_message"
	"net/http"

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
	filter := &dto.SoftwareFilter{}

	items, err := h.Software.FindAllSoftware(filter)
	if err != nil {
		response_message.WrapperResponseJSON(w, 500, err.Error())
		return
	}

	response_message.WrapperResponseJSON(w, 200, items)
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
