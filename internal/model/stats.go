package model

import "github.com/chiyonn/peepa-go/internal/client"

type Stats struct {
	BuyBoxPrice int64 `json:"buyBoxPrice"`
}

func NewStats(s client.RawStats) *Stats {
	return &Stats{
		BuyBoxPrice: s.BuyBoxPrice,
	}
}
