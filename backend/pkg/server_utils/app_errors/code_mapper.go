package app_errors

func HTTPStatusFromCode(code Code) int {
	switch code {
	case CodeNotFound:
		return 404
	case CodeAlreadyExists, CodeConflict:
		return 409
	case CodeInvalidInput:
		return 400
	case CodeUnprocessable:
		return 422
	case CodeForbidden:
		return 403
	case CodeUnauthorized:
		return 401
	case CodeTooManyRequests:
		return 429
	case CodeServiceUnavailable:
		return 503
	default:
		return 500
	}
}