package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"inno-accounting/internal/adapters/jwt"
	"inno-accounting/internal/controllers/middleware"
	"inno-accounting/internal/dto"
	"inno-accounting/pkg/logger"
	"inno-accounting/pkg/server_utils/app_errors"
	"inno-accounting/pkg/server_utils/configure_headers"
	custom_errors "inno-accounting/pkg/server_utils/errors"
	"inno-accounting/pkg/server_utils/response_message"
	"inno-accounting/pkg/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//
// 🔹 CREATE USER
//
func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	var newUser dto.CreateUser
	json.NewDecoder(r.Body).Decode(&newUser)

	if newUser.Email == "" || newUser.Password == "" || newUser.Name == "" || newUser.LastName == "" {
		errStr := "one or more required fields are missing"
		jsonErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(jsonErr.Code), jsonErr.Message)
		return
	}

	user, err := h.User.CreateUser(&newUser)
	if err != nil {
		var appErr *app_errors.AppError

		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(appErr.Code), appErr.Message)
		}
		return
	}

	safetyUser := dto.PublicUserFromModel(user)
	logger.Info(fmt.Sprintf("User created successfully: %+v", safetyUser))
	response_message.WrapperResponseJSON(w, 201, safetyUser)
}

//
// 🔹 GET USER BY ID
//
func (h *Handlers) GetUserByID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	idStr := mux.Vars(r)["id"]
	userUUID, err := uuid.Parse(idStr)
	if err != nil {
		errStr := "invalid user ID"
		idErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(idErr.Code), idErr.Message)
		return
	}

	user, err := h.User.FindUserByID(userUUID)
	if err != nil {
		var appErr *app_errors.AppError

		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(appErr.Code), appErr.Message)
		}
		return
	}

	safetyUser := dto.PublicUserFromModel(user)
	logger.Info("User found successfully")
	response_message.WrapperResponseJSON(w, 200, safetyUser)
}

//
// 🔹 GET ME
//
func (h *Handlers) GetUserByMe(w http.ResponseWriter, r *http.Request) {
	authData := r.Context().Value(middleware.GetUserIDKey()).(jwt.AuthData)

	user, err := h.User.FindUserByID(authData.UserID)
	if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
		return
	}

	convertedUser := dto.PublicUserFromModel(user)
	logger.Info("User fetched successfully")
	response_message.WrapperResponseJSON(w, 200, convertedUser)
}

//
// 🔹 GET ALL USERS
//
func (h *Handlers) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	q := r.URL.Query()

	filter := &dto.UserFilter{
		ID:       utils.StrPtr(q.Get("id")),
		Email:    utils.StrPtr(q.Get("email")),
		Name:     utils.StrPtr(q.Get("name")),
		LastName: utils.StrPtr(q.Get("lastname")),
		Post:     utils.StrPtr(q.Get("post")),
		Grade:    utils.StrPtr(q.Get("grade")),
		City:     utils.StrPtr(q.Get("city")),
	}

	users, err := h.User.FindAllUsers(filter)
	if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
		return
	}

	convertedUsers := dto.SeveralUsersToPublic(users)
	logger.Info("Users fetched successfully")
	response_message.WrapperResponseJSON(w, 200, convertedUsers)
}

//
// 🔹 DELETE USER
//
func (h *Handlers) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	idStr := mux.Vars(r)["id"]
	userUUID, err := uuid.Parse(idStr)
	if err != nil {
		wError := custom_errors.New(err, 400)
		wError.AddLogData(fmt.Sprintf("Invalid user ID: %v", idStr))
		wError.AddResponseData("Invalid user ID")
		custom_errors.ErrorResponse(w, wError, logger.GetLoger())
		return
	}

	err = h.User.DeleteByID(userUUID)
	if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
		return
	}

	logger.Info("User deleted successfully")
	response_message.WrapperResponseJSON(w, 200, "user deleted")
}