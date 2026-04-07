package custom_errors

import "errors"

var (
	ErrNotFound      = New(errors.New("not found"), 404)
	ErrAlreadyExists = New(errors.New("already exists"), 409)
	ErrUnauthorized  = New(errors.New("unauthorized"), 401)
	ErrForbidden     = New(errors.New("forbidden"), 403)
	ErrInvalidInput  = New(errors.New("invalid input"), 422)
	ErrInternal      = New(errors.New("internal error"), 500)
	//ErrBadRequest    = errors.New("bad request")
)
