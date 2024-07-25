package simplelogin_api_client_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Auth       AuthBundle
	HTTPClient *http.Client
}

type AuthBundle struct {
	BaseURL       string
	ApiKey        string
	Email         string
	Password      string
	MfaToken      string
	MfaKey        string
	FacebookToken string
	GoogleToken   string
	Code          string
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type OutputOptions interface {
	AccountOptions
	AliasOptions
}

func NewClient(auth AuthBundle) *Client {
	return &Client{
		Auth: auth,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// TODO: scope 'result' to a specific type; implement some kind of enum for the structs
func (c *Client) NewRequest(req *http.Request, result any) error {
	// Set common headers for requests
	// req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("Accept", "*/*")
	// req.Header.Set("Connection", "Keep-Alive")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	// req.Header.Set("User-Agent", "PostmanRuntime/7.40.0")
	// req.Header.Set("Cookie", "slapp=a7bd211f-45c6-410b-bcec-a8f36cd6097d.tcVYAUskb9y2IW4SHeD1YnlZGPE")

	// Make the request
	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	// Close the response body when the function returns
	defer response.Body.Close()

	// If the status code is not in the 200 range, return an error
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		// Decode the response body into the errorResponse struct
		if err = json.NewDecoder(response.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", response.StatusCode)
	}

	// Read the Body of the response & parse it into the 'result' Object
	body, err := io.ReadAll(response.Body) // response body is []byte
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}

	return nil
}
