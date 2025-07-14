package router

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/chiyonn/peepa-go/internal/handler"
	"github.com/chiyonn/peepa-go/internal/service"
)

func NewRouter(srv *service.ProductService, log *slog.Logger) *chi.Mux {
	r := chi.NewRouter()
	hh := handler.NewHealthHandler(log)
	ph := handler.NewProductHandler(srv, log)

	r.Get("/health", hh.GetHealth)
	r.Get("/products/{asin}", ph.GetByASIN)

	return r
}
