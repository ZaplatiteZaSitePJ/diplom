package jwt_config

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTConfig struct {
	accessSecret  []byte
	refreshSecret []byte
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

func New(ac []byte, rs []byte, aTTL time.Duration, rTTL time.Duration) *JWTConfig {
	return &JWTConfig{
		accessSecret:  ac,
		refreshSecret: rs,
		accessTTL:     aTTL,
		refreshTTL:    rTTL,
	}
}

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateTokens создает пару access + refresh токенов для пользователя
func (jc *JWTConfig) GenerateTokens(userID uuid.UUID, role string) (access, refresh string, err error) {
	accessClaims := Claims{
		UserID: userID.String(),
		Role:   role, 
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jc.accessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	access, err = at.SignedString(jc.accessSecret)
	if err != nil {
		return "", "", err
	}

	refreshClaims := Claims{
		UserID: userID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jc.refreshTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refresh, err = rt.SignedString(jc.refreshSecret)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

// ValidateAccess проверяет access токен
func (jc *JWTConfig) ValidateAccess(token string) (*Claims, error) {
	return jc.validate(token, jc.accessSecret)
}

// ValidateRefresh проверяет refresh токен
func (jc *JWTConfig) ValidateRefresh(token string) (*Claims, error) {
	return jc.validate(token, jc.refreshSecret)
}

// validate общий метод для валидации токена
func (jc *JWTConfig) validate(tokenStr string, secret []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token expired")
		}
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, fmt.Errorf("invalid signature")
		}
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}