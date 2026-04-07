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
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Creating user
func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)
	var newUser dto.CreateUser

	json.NewDecoder(r.Body).Decode(&newUser)

	if newUser.Email=="" || newUser.Password=="" || newUser.Username=="" {
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
		
	} else {
		safetyUser := dto.PublicUser{Username: user.Username, Email: user.Email}
		logger.Info(fmt.Sprintf("User created succesfully: %+v", safetyUser))
		response_message.WrapperResponseJSON(w, 201, safetyUser)
	}
}

// Getting user by ID
func (h *Handlers) GetUserByID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	//GETTING USER ID FROM REQUEST
	idStr := mux.Vars(r)["id"]
	userUUID, err := uuid.Parse(idStr)

	if err != nil {
		errStr := "invalid user ID"
		idErr := app_errors.InvalidInput(errStr, errors.New(errStr))
		response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(idErr.Code), idErr.Message)
		return
	}
	
	findedUser, err := h.User.FindUserByID(userUUID)

		if err != nil {
		var appErr *app_errors.AppError
		
		if errors.As(err, &appErr) {
			response_message.WrapperResponseJSON(w, app_errors.HTTPStatusFromCode(appErr.Code), appErr.Message)
		}
	} else {
		safetyUser := dto.PublicUserFromModel(findedUser)
		logger.Info("User finded succesfully")
		response_message.WrapperResponseJSON(w, 200, safetyUser)
	}
}

// Getting ME user

func (h *Handlers) GetUserByMe(w http.ResponseWriter, r *http.Request) {
	authData := r.Context().Value(middleware.GetUserIDKey()).(jwt.AuthData)

	user, err := h.User.FindUserByID(authData.UserID)
	if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
	} else {
		convertedUsers := dto.PublicUserFromModel(user)
		logger.Info("Users finded succesfully")
		response_message.WrapperResponseJSON(w, 200, convertedUsers)
	}
}

// Getting all user
func (h *Handlers) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	users, err := h.User.FindAllUsers()
	if err != nil {
		custom_errors.ErrorResponse(w, err, logger.GetLoger())
	} else {
		convertedUsers := dto.SeveralUsersToPublic(users)
		logger.Info("Users finded succesfully")
		response_message.WrapperResponseJSON(w, 200, convertedUsers)
	}
}

// Deleting user by ID
func (h *Handlers) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	configure_headers.DefaultHeader(w)

	//GETTING USER ID FROM REQUEST
	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		wError := custom_errors.New(err, 400)
		wError.AddLogData(fmt.Sprintf("Invalid user ID: %v. ID must be decimal", mux.Vars(r)["id"]))
		wError.AddResponseData(fmt.Sprintf("Invalid user ID: %v. ID must be decimal", mux.Vars(r)["id"]))
		custom_errors.ErrorResponse(w, wError, logger.GetLoger())
		return
	}

	err = h.User.DeleteByID(userID)
}