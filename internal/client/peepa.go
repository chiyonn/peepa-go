package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type PeepaClient struct {
	hc          *http.Client
	cfg         *PeepaConfig
	accessToken string
}

var ErrInvalidConfig = errors.New("config must be set")

func NewPeepaClient(cfg *PeepaConfig) (*PeepaClient, error) {
	if cfg.Host == "" || cfg.AuthHost == "" || cfg.ClientID == "" || cfg.RefreshToken == "" {
		return nil, ErrInvalidConfig
	}
	return &PeepaClient{
		hc:  &http.Client{Timeout: 5 * time.Second},
		cfg: cfg,
	}, nil
}

func (c *PeepaClient) getAccessToken() (string, error) {
	body := map[string]any{
		"ClientId": c.cfg.ClientID,
		"AuthFlow": "REFRESH_TOKEN_AUTH",
		"AuthParameters": map[string]string{
			"REFRESH_TOKEN": c.cfg.RefreshToken,
		},
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.cfg.AuthHost, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

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

func (c *PeepaClient) GetByASIN(asin string) error {
	if err := c.ensureAccessToken(); err != nil {
		return err
	}

	reqBody := ProductDetailRequest{
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
		Variables: ProductDetailVariables{
			ASIN:     asin,
			Domain:   "5",
			IsLite:   false,
			IsDetail: true,
			NoCache:  false,
			CountPV:  false,
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/graphql", c.cfg.Host)
	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s – %s", resp.Status, string(bodyBytes))
	}

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Response body: %s\n", string(bodyBytes))
	return nil
}
