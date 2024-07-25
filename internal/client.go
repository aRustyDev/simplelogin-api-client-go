package simplelogin_api_client_go

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

func NewClient(auth AuthBundle) *Client {
	return &Client{
		Auth: auth,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

// TODO: scope 'result' to a specific type; implement some kind of enum for the structs.
func (c *Client) NewRequest(req *http.Request, result any) error {
	// Set common headers for requests.
	// req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("Accept", "*/*")
	// req.Header.Set("Connection", "Keep-Alive")

	log.Debug().Msg(fmt.Sprintf("URL: %+v\n", req.URL))
	log.Debug().Msg(fmt.Sprintf("Headers: %+v\n", req.Header))

	// Make the request.
	response, err := c.HTTPClient.Do(req)
	HandleError(err, zerolog.ErrorLevel, "Failed to make request")
	log.Debug().Msg(fmt.Sprintf("StatusCode: %+v\n", response.StatusCode))

	// Close the response body when the function returns.
	defer response.Body.Close()

	// TODO: review this to see if we can remove it? (do we need the errorResponse struct?).
	// If the status code is not in the 200 range, return an error.
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		// Decode the response body into the errorResponse struct.
		err = json.NewDecoder(response.Body).Decode(&errRes)
		HandleError(err, zerolog.FatalLevel, errRes.Message)

		HandleError(err, zerolog.FatalLevel, fmt.Sprintf("unknown error, status code: %d", response.StatusCode))
	}

	// Read the Body of the response & parse it into the 'result' Object.
	body, err := io.ReadAll(response.Body) // response body is []byte.
	HandleError(err, zerolog.FatalLevel, "Failed to make request")

	// Parse []byte to go struct pointer.
	err = json.Unmarshal(body, &result)
	HandleError(err, zerolog.FatalLevel, "Can not unmarshal JSON")
	log.Debug().Msg(fmt.Sprintf("Result: %+v\n", result))

	return nil
}
