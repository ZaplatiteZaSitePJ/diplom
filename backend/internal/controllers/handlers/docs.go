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

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// ===================== CREATE =====================

func (h *Handlers) CreateDocument(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	var input dto.DocsItemPublic

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid json body")
		return
	}

	if input.DocNumber == "" || input.ResponsibleWorkerEmail == "" {
		errStr := "doc_number and responsible_worker_email are required"
		jsonErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(
			w,
			app_errors.HTTPStatusFromCode(jsonErr.Code),
			jsonErr.Message,
		)
		return
	}

	doc, err := h.Documents.CreateDocument(&input)
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

		logger.Error("CreateDocument error:", err)
		response_message.WrapperResponseJSON(w, 500, "internal server error")
		return
	}

	logger.Info(fmt.Sprintf("Document created successfully: %+v", doc))
	response_message.WrapperResponseJSON(w, 201, doc)
}

// ===================== GET ALL =====================

func (h *Handlers) GetAllDocuments(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	query := r.URL.Query()

	filter := &dto.DocsFilter{}

	if v := query.Get("id"); v != "" {
		id, err := uuid.Parse(v)
		if err == nil {
			filter.ID = &id
		} else {
			response_message.WrapperResponseJSON(w, 400, "invalid uuid format")
			return
		}
	}
	if v := query.Get("doc_number"); v != "" {
		filter.DocNumber = &v
	}
	if v := query.Get("last_worker_email"); v != "" {
		filter.LastWorkerEmail = &v
	}
	if v := query.Get("last_storage"); v != "" {
		filter.LastStorage = &v
	}
	if v := query.Get("category"); v != "" {
		filter.Category = &v
	}
	if v := query.Get("transfer_status"); v != "" {
		filter.TransferStatus = &v
	}

	docs, err := h.Documents.FindAllDocuments(filter)
	if err != nil {
		var appErr *app_errors.AppError
		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(
				w,
				app_errors.HTTPStatusFromCode(appErr.Code),
				appErr.Message,
			)
		} else {
			logger.Error("Unhandled error in GetAllDocuments:", err)
			response_message.WrapperResponseJSON(w, 500, "internal server error")
		}
		return
	}

	if docs == nil {
		docs = []*dto.DocsItemPublic{}
	}

	response_message.WrapperResponseJSON(w, 200, docs)
}

// ===================== GET BY ID =====================

func (h *Handlers) GetDocumentByID(w http.ResponseWriter, r *http.Request) {
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

	docID, err := uuid.Parse(idStr)
	if err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid uuid format")
		return
	}

	doc, err := h.Documents.FindDocumentByID(docID)
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

		logger.Error("GetDocumentByID error:", err)
		response_message.WrapperResponseJSON(w, 500, "internal server error")
		return
	}

	response_message.WrapperResponseJSON(w, 200, doc)
}

// ===================== PATCH =====================

func (h *Handlers) PatchDocumentByID(w http.ResponseWriter, r *http.Request) {
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

	docID, err := uuid.Parse(idStr)
	if err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid uuid format")
		return
	}

	var input dto.DocsItemPublic
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid json body")
		return
	}

	updated, err := h.Documents.ChangeDocumentByID(docID, &input)
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

		logger.Error("PatchDocumentByID error:", err)
		response_message.WrapperResponseJSON(w, 500, "internal server error")
		return
	}

	response_message.WrapperResponseJSON(w, 200, updated)
}