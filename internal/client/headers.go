package client

var commonHeaders = map[string]string{
	"Origin":             "https://search.eresa.jp",
	"Priority":           "u=1, i",
	"Referer":            "https://search.eresa.jp/",
	"Sec-CH-UA":          `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`,
	"Sec-CH-UA-Mobile":   "?0",
	"Sec-CH-UA-Platform": `"Windows"`,
	"Sec-Fetch-Dest":     "empty",
	"Sec-Fetch-Mode":     "cors",
	"Sec-Fetch-Site":     "cross-site",
	"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
}

var tokenHeaders = map[string]string{
	"Content-Type":  "application/x-amz-json-1.1",
	"X-Amz-Target":   "AWSCognitoIdentityProviderService.InitiateAuth",
	"X-Amz-User-Agent": "aws-amplify/0.1.x js",
}

var graphqlHeaders = map[string]string{
	"Content-Type":       "application/json",
	"X-Amz-User-Agent":   "aws-amplify/3.8.23 js",
}

