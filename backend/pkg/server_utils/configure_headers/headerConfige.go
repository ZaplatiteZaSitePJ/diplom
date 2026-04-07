package configure_headers

import "net/http"

func DefaultHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
