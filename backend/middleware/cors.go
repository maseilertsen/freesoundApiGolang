package middleware

import (
	"net/http"
	"os"
)

// HandlePreflight returns a handler for OPTIONS preflight requests
func HandlePreflight(w http.ResponseWriter, r *http.Request) {
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusNoContent)
}
