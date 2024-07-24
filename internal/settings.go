package simplelogin_api_client_go

import "context"

type SettingOptions struct {
	Device       string `json:"alias_generator"`
	Notification bool   `json:"notification"`
	Domain       string `json:"random_alias_default_domain"`
	Format       string `json:"sender_format"`
	Suffix       string `json:"random_alias_suffix"`
}

// GET /api/setting: Get user's settings.
func (c *Client) GetUserSettings(ctx context.Context, options *SettingOptions) (*Client, error) {
	return nil, nil
}

// PATCH /api/setting: Update user's settings.
func (c *Client) PatchUserSettings(ctx context.Context, options *SettingOptions) (*Client, error) {
	return nil, nil
}

// GET /api/v2/setting/domains: Get domains that user can use to create random alias.
func (c *Client) GetUserDomains(ctx context.Context, options *SettingOptions) (*Client, error) {
	return nil, nil
}
