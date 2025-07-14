package handler

import (
	"net/http"
	"encoding/json"
	"log/slog"
)

type HealthHandler struct {
	log *slog.Logger
}

func NewHealthHandler(log *slog.Logger) *HealthHandler {
	return &HealthHandler{
		log: log,
	}
}

func (h *HealthHandler) GetHealth(w http.ResponseWriter, r *http.Request) {
	var healthStatus struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}

	healthStatus.Message = "Service is healthy"
	healthStatus.Code = 200

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthStatus)
}

