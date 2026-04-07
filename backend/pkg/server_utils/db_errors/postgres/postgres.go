package pg_err

import (
	"errors"

	"github.com/lib/pq"
)

// unwrap pq.Error from err
func getPgError(err error) (*pq.Error, bool) {
	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		return pgErr, true
	}
	return nil, false
}

func IsUniqueViolation(err error) bool {
	if pgErr, ok := getPgError(err); ok {
		return pgErr.Code == CodeUniqueViolation
	}
	return false
}

func IsForeignKeyViolation(err error) bool {
	if pgErr, ok := getPgError(err); ok {
		return pgErr.Code == CodeForeignKeyViolation
	}
	return false
}

func IsNotNullViolation(err error) bool {
	if pgErr, ok := getPgError(err); ok {
		return pgErr.Code == CodeNotNullViolation
	}
	return false
}

func IsCheckViolation(err error) bool {
	if pgErr, ok := getPgError(err); ok {
		return pgErr.Code == CodeCheckViolation
	}
	return false
}

func IsInvalidInputSyntax(err error) bool {
	if pgErr, ok := getPgError(err); ok {
		return pgErr.Code == CodeInvalidTextRepresentation
	}
	return false
}

func IsSyntaxError(err error) bool {
	if pgErr, ok := getPgError(err); ok {
		return pgErr.Code == CodeSyntaxError
	}
	return false
}

func IsSerializationFailure(err error) bool {
	if pgErr, ok := getPgError(err); ok {
		return pgErr.Code == CodeSerializationFailure
	}
	return false
}

func IsConnectionException(err error) bool {
	if pgErr, ok := getPgError(err); ok {
		return pgErr.Code == CodeConnectionException
	}
	return false
}
