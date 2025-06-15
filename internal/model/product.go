package model

import (
	"strconv"
	"strings"

	"github.com/chiyonn/peepa-go/internal/client"
)

type Product struct {
	ASIN       string
	Title      string
	Categories []string
	Images     []string
}

func NewProduct(p *client.RawProduct) *Product {
	return &Product{
		ASIN:       p.ASIN,
		Title:      p.Title,
		Categories: toStrs(p.Categories),
		Images:     strings.Split(p.ImagesCSV, ","),
	}
}

func toStrs(ints []int64) []string {
	strSli := make([]string, len(ints))
	for i, v := range ints {
		strSli[i] = strconv.FormatInt(v, 10)
	}
	return strSli
}
