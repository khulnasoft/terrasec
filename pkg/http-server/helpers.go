package httpserver

import (
	"fmt"
	"net/http"
)

// apiResponse creates an API response
func apiResponse(w http.ResponseWriter, msg string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprint(w, msg)
}

// apiErrorResponse creates an API error response
func apiErrorResponse(w http.ResponseWriter, errMsg string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	http.Error(w, errMsg, statusCode)
}
