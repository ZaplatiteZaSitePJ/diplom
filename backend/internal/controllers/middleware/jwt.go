package middleware

import (
	"context"
	"fmt"
	"inno-accounting/internal/adapters/jwt"
	"inno-accounting/pkg/logger"
	custom_errors "inno-accounting/pkg/server_utils/errors"
	"net/http"
	"strings"
)

type contextKey string

const userIDKey contextKey = "user_id"

func GetUserIDKey() contextKey {
	return  userIDKey
}

// AuthData хранится в контексте запроса
type AuthData struct {
	UserID string
	Role   string
}

// JWTMiddleware проверяет access-токен и кладет AuthData в контекст
func JWTMiddleware(tokenManager jwt.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authError := custom_errors.New(fmt.Errorf("unauthorized"), 401)
			authError.AddLogData("unauthorized")
			authError.AddResponseData("unauthorized")

			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				custom_errors.ErrorResponse(w, authError, logger.GetLoger())
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			authData, err := tokenManager.ValidateAccess(token)
			if err != nil {
				custom_errors.ErrorResponse(w, authError, logger.GetLoger())
				return
			}

			// кладем в контекст
			ctx := context.WithValue(r.Context(), userIDKey, authData)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetAuthData достает AuthData из контекста
func GetAuthData(ctx context.Context) jwt.AuthData {
	auth, ok := ctx.Value(userIDKey).(jwt.AuthData)
	if !ok {
		return jwt.AuthData{}
	}
	return auth
}

// RoleMiddleware проверяет, есть ли у пользователя нужная роль
func RoleMiddleware(allowedRoles ...string) func(http.Handler) http.Handler {
	roleSet := make(map[string]bool)
	for _, r := range allowedRoles {
		roleSet[r] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := GetAuthData(r.Context())
			if !roleSet[auth.Role] {
				authError := custom_errors.New(fmt.Errorf("forbidden"), 403)
				authError.AddLogData("forbidden")
				authError.AddResponseData("forbidden")
				custom_errors.ErrorResponse(w, authError, logger.GetLoger())
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}