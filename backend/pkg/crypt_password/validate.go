package crypt_password

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

var (
    uppercaseRegex   = regexp.MustCompile(`[A-Z]`)
    lowercaseRegex   = regexp.MustCompile(`[a-z]`)
    digitRegex       = regexp.MustCompile(`[0-9]`)
    specialCharRegex = regexp.MustCompile(`[!@#\$%\^&\*\(\)\-\_\=\+\[\]\{\}\\|;:'",<\.>\/\?]`)
)

func ValidatePassword(password string) error {
    length := utf8.RuneCountInString(password)
    if length < 8 {
        return errors.New("password must be at least 8 characters long")
    }

    if !uppercaseRegex.MatchString(password) {
        return errors.New("password must contain at least one uppercase letter")
    }
    if !lowercaseRegex.MatchString(password) {
        return errors.New("password must contain at least one lowercase letter")
    }
    if !digitRegex.MatchString(password) {
        return errors.New("password must contain at least one digit")
    }
    if !specialCharRegex.MatchString(password) {
        return errors.New("password must contain at least one special character")
    }

    return nil
}