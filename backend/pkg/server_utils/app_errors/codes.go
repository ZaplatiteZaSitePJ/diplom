package app_errors

type Code string

const (
	CodeNotFound           Code = "NOT_FOUND"           // 404
	CodeAlreadyExists      Code = "ALREADY_EXISTS"      // 409
	CodeInvalidInput       Code = "INVALID_INPUT"       // 400
	CodeForbidden          Code = "FORBIDDEN"           // 403
	CodeInternal           Code = "INTERNAL"            // 500
	CodeUnprocessable      Code = "UNPROCESSABLE"       // 422
	CodeUnauthorized       Code = "UNAUTHORIZED"        // 401
	CodeConflict           Code = "CONFLICT"            // 409
	CodeTooManyRequests    Code = "TOO_MANY_REQUESTS"   // 429
	CodeServiceUnavailable Code = "SERVICE_UNAVAILABLE" // 503
)

// Fabric

func NotFound(msg string, err error) *AppError {
	return &AppError{Code: CodeNotFound, Message: msg, Err: err}
}

func AlreadyExists(msg string, err error) *AppError {
	return &AppError{Code: CodeAlreadyExists, Message: msg, Err: err}
}

func InvalidInput(msg string, err error) *AppError {
	return &AppError{Code: CodeInvalidInput, Message: msg, Err: err}
}

func Forbidden(msg string, err error) *AppError {
	return &AppError{Code: CodeForbidden, Message: msg, Err: err}
}

func Internal(msg string, err error) *AppError {
	return &AppError{Code: CodeInternal, Message: msg, Err: err}
}

func Unprocessable(msg string, err error) *AppError {
	return &AppError{Code: CodeUnprocessable, Message: msg, Err: err}
}

func Unauthorized(msg string, err error) *AppError {
	return &AppError{Code: CodeUnauthorized, Message: msg, Err: err}
}

func Conflict(msg string, err error) *AppError {
	return &AppError{Code: CodeConflict, Message: msg, Err: err}
}

func TooManyRequests(msg string, err error) *AppError {
	return &AppError{Code: CodeTooManyRequests, Message: msg, Err: err}
}

func ServiceUnavailable(msg string, err error) *AppError {
	return &AppError{Code: CodeServiceUnavailable, Message: msg, Err: err}
}