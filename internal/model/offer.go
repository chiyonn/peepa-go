package model

import "github.com/chiyonn/peepa-go/internal/client"

type Offer struct {
	LastSeen         int
	SellerID         string
	OfferCSV         []int
	Condition        int
	IsPrime          bool
	IsMAP            bool
	IsShippable      bool
	IsAddonItem      bool
	IsPreorder       bool
	IsWarehouseDeal  bool
	IsScam           bool
	IsAmazon         bool
	IsPrimeExcl      bool
	OfferID          int
	IsFBA            bool
	ShipsFromChina   bool
	MinOrderQty      int
	CouponHistory    []int
	LastStockUpdate  int
}

func NewOffer(o client.RawOffer) *Offer {
	return &Offer{
		SellerID: o.SellerID,
		OfferCSV: o.OfferCSV,
		Condition: o.Condition,
		IsPrime: o.IsPrime,
		IsMAP: o.IsMAP,
		IsShippable: o.IsShippable,
		IsAddonItem: o.IsAddonItem,
		IsPreorder: o.IsPreorder,
		IsWarehouseDeal: o.IsWarehouseDeal,
		IsScam: o.IsScam,
		IsAmazon: o.IsAmazon,
		IsPrimeExcl: o.IsPrimeExcl,
		OfferID: o.OfferID,
		IsFBA: o.IsFBA,
		ShipsFromChina: o.ShipsFromChina,
		MinOrderQty: o.MinOrderQty,
		CouponHistory: o.CouponHistory,
		LastStockUpdate: o.LastStockUpdate,
	}
}
