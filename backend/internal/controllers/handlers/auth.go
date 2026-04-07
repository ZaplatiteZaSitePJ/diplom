package handlers

import (
	"encoding/json"
	"inno-accounting/pkg/server_utils/response_message"
	"net/http"
)

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	access, refresh, err := h.Auth.Login(req.Email, req.Password)
	if err != nil {
		response_message.WrapperResponseJSON(w, 401, "invalid credentials")
		return
	}

	response_message.WrapperResponseJSON(w, 200, map[string]string{
		"access":  access,
		"refresh": refresh,
	})
}