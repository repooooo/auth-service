package handler

import (
	"encoding/json"
	"github.com/repooooo/go-utils/sl"
	"log/slog"
	"net/http"
)

// HealthCheck handles the health check requests
// It returns a JSON response indicating the health status of the service.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{"status": "healthy"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		slog.Error("Failed to encode health check response", sl.Err(err))
	}
}
