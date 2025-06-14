package client

type ProductDetailRequest struct {
	Query     string                 `json:"query"`
	Variables ProductDetailVariables `json:"variables"`
}

type ProductDetailVariables struct {
	ASIN     string `json:"asin"`
	Domain   string `json:"domain"`
	IsLite   bool   `json:"isLite"`
	IsDetail bool   `json:"isDetail"`
	NoCache  bool   `json:"nocache"`
	CountPV  bool   `json:"countpv"`
}
