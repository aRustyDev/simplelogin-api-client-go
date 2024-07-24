package simplelogin_api_client_go

import "context"

type MiscOptions struct {
	Receipt  string `json:"receipt_data"`
	IsMacApp bool   `json:"is_macapp"`
}

// POST /api/apple/process_payment: Process Apple's receipt.
func (c *Client) PostAppleReceipt(ctx context.Context, options *MiscOptions) (*Client, error) {
	return nil, nil
}
