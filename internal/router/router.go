package router

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/chiyonn/peepa-go/internal/handler"
	"github.com/chiyonn/peepa-go/internal/service"
)

func NewRouter(srv *service.ProductService, log *slog.Logger) *chi.Mux {
	r := chi.NewRouter()
	h := handler.NewProductHandler(srv, log)

	r.Get("/products/{asin}", h.GetByASIN)

	return r
}
