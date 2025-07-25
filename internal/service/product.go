package service

import (
	"log/slog"

	"github.com/chiyonn/peepa-go/internal/client"
	"github.com/chiyonn/peepa-go/internal/model"
)

type ProductService struct {
	cli *client.PeepaClient
	log *slog.Logger
}

func NewProductService(cli *client.PeepaClient, log *slog.Logger) *ProductService {
	return &ProductService{
		cli: cli,
		log: log,
	}
}

func (s *ProductService) GetByASIN(asin string) (*model.Product, error) {
	raw, err := s.cli.GetByASIN(asin)
	if err != nil {
		s.log.Error("failed to get product from client", "error", err)
		return nil, err
	}
	product, err := model.NewProduct(raw)
	if err != nil {
		s.log.Error("failed to create product from raw data", "error", err)
		return nil, err
	}
	return product, nil
}
