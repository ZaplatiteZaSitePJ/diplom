package crypt_password

import (
	"golang.org/x/crypto/bcrypt"
)

// Хэширование пароля
func EncryptPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
