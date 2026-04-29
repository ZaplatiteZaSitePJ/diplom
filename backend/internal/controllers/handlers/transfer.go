package handlers

// import (
// 	"encoding/json"
// 	"errors"
// 	"inno-accounting/internal/domain"
// 	"inno-accounting/internal/dto"
// 	"inno-accounting/pkg/server_utils/app_errors"
// 	"inno-accounting/pkg/server_utils/configure_headers"
// 	"inno-accounting/pkg/server_utils/response_message"
// 	"net/http"

// 	"github.com/google/uuid"
// )

// func (h *Handlers) MoveItem(w http.ResponseWriter, r *http.Request) {
// 	configure_headers.DefaultHeader(w)

// 	var input dto.MoveItemRequest

// 	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
// 		response_message.WrapperResponseJSON(w, 400, "invalid json body")
// 		return
// 	}

// 	if input.ItemID == "" || input.ToType == "" {
// 		errStr := "item_id and to_type are required"
// 		jsonErr := app_errors.InvalidInput(errStr, errors.New(errStr))
// 		response_message.WrapperResponseJSON(
// 			w,
// 			app_errors.HTTPStatusFromCode(jsonErr.Code),
// 			jsonErr.Message,
// 		)
// 		return
// 	}

// 	itemID, err := uuid.Parse(input.ItemID)
// 	if err != nil {
// 		response_message.WrapperResponseJSON(w, 400, "invalid item_id")
// 		return
// 	}

// 	var toID *uuid.UUID
// 	if input.ToID != nil {
// 		parsed, err := uuid.Parse(*input.ToID)
// 		if err != nil {
// 			response_message.WrapperResponseJSON(w, 400, "invalid to_id")
// 			return
// 		}
// 		toID = &parsed
// 	}

// 	err = h.Location.MoveItem(
// 		itemID,
// 		domain.LocationType(input.ToType),
// 		toID,
// 	)

// 	if err != nil {
// 		var appErr *app_errors.AppError
// 		if errors.As(err, &appErr) {
// 			response_message.WrapperResponseJSON(
// 				w,
// 				app_errors.HTTPStatusFromCode(appErr.Code),
// 				appErr.Message,
// 			)
// 			return
// 		}

// 		response_message.WrapperResponseJSON(w, 500, "internal server error")
// 		return
// 	}

// 	response_message.WrapperResponseJSON(w, 200, "item moved successfully")
// }

// func (h *Handlers) GetLocation(w http.ResponseWriter, r *http.Request) {
// 	configure_headers.DefaultHeader(w)

// 	itemIDStr := r.URL.Query().Get("item_id")
// 	if itemIDStr == "" {
// 		response_message.WrapperResponseJSON(w, 400, "item_id is required")
// 		return
// 	}

// 	itemID, err := uuid.Parse(itemIDStr)
// 	if err != nil {
// 		response_message.WrapperResponseJSON(w, 400, "invalid item_id")
// 		return
// 	}

// 	loc, err := h.Location.GetLocation(itemID)
// 	if err != nil {
// 		var appErr *app_errors.AppError
// 		if errors.As(err, &appErr) {
// 			response_message.WrapperResponseJSON(
// 				w,
// 				app_errors.HTTPStatusFromCode(appErr.Code),
// 				appErr.Message,
// 			)
// 			return
// 		}

// 		response_message.WrapperResponseJSON(w, 500, "internal server error")
// 		return
// 	}

// 	response_message.WrapperResponseJSON(w, 200, loc)
// }