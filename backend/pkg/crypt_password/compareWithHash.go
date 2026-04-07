package crypt_password

import "golang.org/x/crypto/bcrypt"

func CompareWithHash(hashedPassword string, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}