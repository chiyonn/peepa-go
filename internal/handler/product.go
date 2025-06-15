package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/chiyonn/peepa-go/internal/service"
)

type ProductHandler struct {
	srv *service.ProductService
	log *slog.Logger
}

func NewProductHandler(srv *service.ProductService, log *slog.Logger) *ProductHandler {
	return &ProductHandler{
		srv: srv,
		log: log,
	}
}

func (h *ProductHandler) GetByASIN(w http.ResponseWriter, r *http.Request) {
	asin := chi.URLParam(r, "asin")
	if asin == "" {
		http.Error(w, "asin required", http.StatusBadRequest)
	}

	product, err := h.srv.GetByASIN(asin)
	if err != nil {
		h.log.Error("failed to get product: %v", "error", err)
		http.Error(w, "failed to fetch product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(product); err != nil {
		h.log.Error("failed to encode response: %v", "error", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
