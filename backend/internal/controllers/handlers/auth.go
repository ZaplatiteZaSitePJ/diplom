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

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response_message.WrapperResponseJSON(w, 400, "invalid request")
		return
	}

	access, refresh, role, err := h.Auth.Login(req.Email, req.Password)
	if err != nil {
		response_message.WrapperResponseJSON(w, 401, "invalid credentials")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refresh,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		MaxAge:   7 * 24 * 3600,
		SameSite: http.SameSiteStrictMode,
	})

	response_message.WrapperResponseJSON(w, 200, map[string]string{
		"access": access,
		"role":   role,
	})
}

func (h *Handlers) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		response_message.WrapperResponseJSON(w, 401, "no refresh token")
		return
	}

	access, newRefresh, role, err := h.Auth.Refresh(cookie.Value)
	if err != nil {
		response_message.WrapperResponseJSON(w, 401, "invalid refresh token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    newRefresh,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		MaxAge:   7 * 24 * 3600,
		SameSite: http.SameSiteStrictMode,
	})

	response_message.WrapperResponseJSON(w, 200, map[string]string{
		"access": access,
		"role":   role,
	})
}

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err == nil {
		_ = h.Auth.Logout(cookie.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		MaxAge:   -1,
	})

	response_message.WrapperResponseJSON(w, 200, "logout success")
}