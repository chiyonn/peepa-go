package model

import "github.com/chiyonn/peepa-go/internal/client"

type Stats struct {
	SalesRankDrops30  int64
	SalesRankDrops90  int64
	SalesRankDrops180 int64
	SalesRankDrops365 int64
	BuyBoxPrice       int64
}

func NewStats(s client.RawStats) *Stats {
	return &Stats{
		SalesRankDrops30:  s.SalesRankDrops30,
		SalesRankDrops90:  s.SalesRankDrops90,
		SalesRankDrops180: s.SalesRankDrops180,
		SalesRankDrops365: s.SalesRankDrops365,
		BuyBoxPrice:       s.BuyBoxPrice,
	}
}
