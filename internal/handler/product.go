package handler

import (
	"net/http"
	"fmt"

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

	product := h.cli.GetByASIN(asin)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"asin": "%s", "title": "Sample Product"}`, asin)
}
