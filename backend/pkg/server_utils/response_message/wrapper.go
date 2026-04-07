package response_message

import (
	"encoding/json"
	"net/http"
)

func WrapperResponseJSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	var is_error bool

	if statusCode > 400 {
		is_error = true
	}

	res := NewResponseMessage(statusCode, data, is_error)
	json.NewEncoder(w).Encode(res)
}
