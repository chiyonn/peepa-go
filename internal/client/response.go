package client

type AuthResponse struct {
	Result struct {
		AccessToken string `json:"AccessToken"`
	} `json:"AuthenticationResult"`
}

type rawProductDetailResponse struct {
	Data struct {
		GetProductDetail MetaData `json:"getProductDetail"`
	} `json:"data"`
}

type MetaData struct {
	ASIN      string `json:"asin"`
	CreatedAt string `json:"createdAt"`
	JSON      string `json:"json"`
}

type RawProduct struct {
	ASIN            string     `json:"asin"`
	Title           string     `json:"title"`
	RootCategory    int64      `json:"rootCategory"`
	Categories      []int64    `json:"categories"`
	ImagesCSV       string     `json:"imagesCSV"`
	Brand           string     `json:"brand"`
	Manifacturer    string     `json:"manifacturer"`
	LastPriceChange int64      `json:"lastPriceChange"`
	LastUpdate      int64      `json:"lastUpdated"`
	Offers          []RawOffer `json:"offers"`
}

type RawOffer struct {
	LastSeen        int    `json:"lastSeen"`
	SellerID        string `json:"sellerId"`
	OfferCSV        []int  `json:"offerCSV"`
	Condition       int    `json:"condition"`
	IsPrime         bool   `json:"isPrime"`
	IsMAP           bool   `json:"isMAP"`
	IsShippable     bool   `json:"isShippable"`
	IsAddonItem     bool   `json:"isAddonItem"`
	IsPreorder      bool   `json:"isPreorder"`
	IsWarehouseDeal bool   `json:"isWarehouseDeal"`
	IsScam          bool   `json:"isScam"`
	IsAmazon        bool   `json:"isAmazon"`
	IsPrimeExcl     bool   `json:"isPrimeExcl"`
	OfferID         int    `json:"offerId"`
	IsFBA           bool   `json:"isFBA"`
	ShipsFromChina  bool   `json:"shipsFromChina"`
	MinOrderQty     int    `json:"minOrderQty"`
	CouponHistory   []int  `json:"couponHistory"`
	LastStockUpdate int    `json:"lastStockUpdate"`
}
