package simplelogin_api_client_go

import "context"

// GET /api/custom_domains: Get custom domains.
func (c *Client) GetDomain(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// TODO: consider new Fn name? (domain vs alias focus)
// GET /api/custom_domains/:custom_domain_id/trash: Get deleted aliases of a custom domain.
func (c *Client) GetDeletedAlias(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// PATCH /api/custom_domains/:custom_domain_id: Update custom domain's information.
func (c *Client) PatchDomain(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}
