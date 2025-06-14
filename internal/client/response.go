package client

type AuthResponse struct {
	Result struct {
		AccessToken string `json:"AccessToken"`
	} `json:"AuthenticationResult"`
}
