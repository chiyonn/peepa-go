package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/chiyonn/peepa-go/internal/client"
)

type ProductHandler struct {
	cli *client.PeepaClient
}

func NewProductHandler(cli *client.PeepaClient) *ProductHandler {
	return &ProductHandler{
		cli: cli,
	}
}

func (h *ProductHandler) GetByASIN(w http.ResponseWriter, r *http.Request) {
	asin := chi.URLParam(r, "asin")

	product, err := h.cli.GetByASIN(asin)
	if err != nil {
		log.Printf("failed to get product: %v", err)
		http.Error(w, "failed to fetch product", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(product); err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
