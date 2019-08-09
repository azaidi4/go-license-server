package apiutils

import "net/http"

// Send status code in response
func ReturnWithStatus(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}
