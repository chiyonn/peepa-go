package model

import (
	"errors"
	"strings"

	"github.com/chiyonn/peepa-go/internal/client"
)

type Product struct {
	ASIN            string
	Title           string
	RootCategory    int64
	Categories      []int64
	Images          []string
	Brand           string
	Manifacturer    string
	Offers          []*Offer
	Stats           *Stats
	BuyBoxPrice     int64
	LastPriceChange int64
	LastUpdated     int64
}

func toOffers(raws []client.RawOffer) []*Offer {
	var offers []*Offer
	for _, o := range raws {
		offers = append(offers, NewOffer(o))
	}
	return offers
}

func NewProduct(p *client.RawProduct) (*Product, error) {
	if p == nil {
		return nil, ErrRawNill
	}
	return &Product{
		ASIN:            p.ASIN,
		Title:           p.Title,
		RootCategory:    p.RootCategory,
		Categories:      p.Categories,
		Images:          strings.Split(p.ImagesCSV, ","),
		Brand:           p.Brand,
		Manifacturer:    p.Manifacturer,
		Offers:          toOffers(p.Offers),
		Stats:           NewStats(p.Stats),
		LastPriceChange: p.LastPriceChange,
		LastUpdated:     p.LastUpdate,
	}, nil
}

var ErrRawNill = errors.New("cannot perse nil")
