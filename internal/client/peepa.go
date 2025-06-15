package client

import (
	"strings"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"log/slog"
)

type PeepaClient struct {
	hc          *http.Client
	log         *slog.Logger
	cfg         *PeepaConfig
	accessToken string
}

var ErrInvalidConfig = errors.New("config must be set")

func NewPeepaClient(cfg *PeepaConfig, log *slog.Logger) (*PeepaClient, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}
	if log == nil {
		log = slog.Default()
	}
	return &PeepaClient{
		hc:  &http.Client{Timeout: 5 * time.Second},
		cfg: cfg,
		log: log,
	}, nil
}

func validateConfig(cfg *PeepaConfig) error {
	var missing []string
	if cfg.Host == "" {
		missing = append(missing, "Host")
	}
	if cfg.AuthHost == "" {
		missing = append(missing, "AuthHost")
	}
	if cfg.ClientID == "" {
		missing = append(missing, "ClientID")
	}
	if cfg.RefreshToken == "" {
		missing = append(missing, "RefreshToken")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing config fields: %v", missing)
	}
	return nil
}

func buildTokenRequestBody(cfg *PeepaConfig) (io.Reader, error) {
	body := map[string]any{
		"ClientId": cfg.ClientID,
		"AuthFlow": "REFRESH_TOKEN_AUTH",
		"AuthParameters": map[string]string{
			"REFRESH_TOKEN": cfg.RefreshToken,
		},
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(bodyBytes), nil
}

func buildProductRequestBody(vars ProductDetailVariables) (io.Reader, error) {
	body := ProductDetailRequest{
		Query: `
			query GetProductDetail($asin: String, $domain: String, $isLite: Boolean, $isDetail: Boolean, $nocache: Boolean, $countpv: Boolean) {
				getProductDetail(asin: $asin, domain: $domain, isLite: $isLite, isDetail: $isDetail, nocache: $nocache, countpv: $countpv) {
					asin
					json
					createdAt
					updatedAt
					__typename
				}
			}`,
		Variables: vars,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(bodyBytes), nil
}

func (c *PeepaClient) ensureAccessToken() error {
	if c.accessToken != "" {
		return nil
	}
	t, err := c.getAccessToken()
	if err != nil {
		return fmt.Errorf("failed to initialize access token: %w", err)
	}
	c.accessToken = t
	return nil
}

func (c *PeepaClient) getAccessToken() (string, error) {
	body, err := buildTokenRequestBody(c.cfg)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.cfg.AuthHost, body)
	if err != nil {
		return "", err
	}

	for k, v := range commonHeaders {
		req.Header.Set(k, v)
	}
	for k, v := range tokenHeaders {
		req.Header.Set(k, v)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("auth request failed with status: %s", resp.Status)
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return "", err
	}

	return authResp.Result.AccessToken, nil
}

func (c *PeepaClient) GetByASIN(asin string) (*RawProduct, error) {
	if err := c.ensureAccessToken(); err != nil {
		return nil, err
	}

	respBody, err := c.doGraphQLRequest(asin)
	if err != nil {
		return nil, err
	}

	return c.parseProductDetailResponse(respBody)
}

func (c *PeepaClient) doGraphQLRequest(asin string) ([]byte, error) {
	body, err := buildProductRequestBody(ProductDetailVariables{
		ASIN:     asin,
		Domain:   "5",
		IsLite:   false,
		IsDetail: true,
		NoCache:  false,
		CountPV:  false,
	})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/graphql", c.cfg.Host)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	for k, v := range commonHeaders {
		req.Header.Set(k, v)
	}
	for k, v := range graphqlHeaders {
		req.Header.Set(k, v)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		c.log.Error("GraphQL API error", "status", resp.Status, "body", string(bodyBytes))
		return nil, fmt.Errorf("GraphQL request failed with status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func (c *PeepaClient) parseProductDetailResponse(data []byte) (*RawProduct, error) {
	var rawResp rawProductDetailResponse
	if err := json.Unmarshal(data, &rawResp); err != nil {
		return nil, fmt.Errorf("failed to parse raw response: %w", err)
	}

	jsonStr := rawResp.Data.GetProductDetail.JSON
	if strings.TrimSpace(jsonStr) == "" {
    return nil, errors.New("product detail JSON is empty")
	}
	var products []RawProduct
	if err := json.Unmarshal([]byte(jsonStr), &products); err != nil {
		return nil, fmt.Errorf("failed to parse product JSON: %w", err)
	}

	if len(products) == 0 {
		return nil, errors.New("no product found")
	}

	c.log.Info("Fetched product detail", "product", products[0].ASIN)
	return &products[0], nil
}
