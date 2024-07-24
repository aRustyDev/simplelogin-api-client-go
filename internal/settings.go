package simplelogin_api_client_go

import "context"

// GET /api/setting: Get user's settings.
func (c *Client) GetUserSettings(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// PATCH /api/setting: Update user's settings.
func (c *Client) PatchUserSettings(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// GET /api/v2/setting/domains: Get domains that user can use to create random alias.
func (c *Client) GetUserDomains(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}
