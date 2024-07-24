package simplelogin_api_client_go

import (
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://app.simplelogin.io"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

type AuthBundle struct {
	BaseURL string
	apiKey  string
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewClient(auth AuthBundle) *Client {
	return &Client{
		BaseURL: auth.BaseURL,
		apiKey:  auth.apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
