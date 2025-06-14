package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/chiyonn/peepa-go/internal/client"
	"github.com/chiyonn/peepa-go/internal/handler"
)

func NewRouter(pc *client.PeepaClient) *chi.Mux {
	r := chi.NewRouter()
	h := handler.NewProductHandler(pc)

	r.Get("/products/{asin}", h.GetByASIN)

	return r
}
