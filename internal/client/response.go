package client

type AuthResponse struct {
	Result struct {
		AccessToken string `json:"AccessToken"`
	} `json:"AuthenticationResult"`
}

type rawProductDetailResponse  struct {
	Data struct {
		GetProductDetail MetaData `json:"getProductDetail"`
	} `json:"data"`
}

type MetaData struct {
	ASIN		  string `json:"asin"`
	CreatedAt string `json:"createdAt"`
	JSON      string `json:"json"`
}

type Product struct {
	Title string `json:"title"`
	Categories []int64 `json:"categories"`
	ImagesCSV string `json:"imagesCSV"`
}
