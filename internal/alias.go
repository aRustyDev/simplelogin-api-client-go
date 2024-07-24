package simplelogin_api_client_go

import "context"

// GET /api/v5/alias/options: Get alias options. Used by create alias process.
func (c *Client) GetAliasOptions(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// GET /api/v2/aliases: Get user's aliases.
func (c *Client) GetAliases(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// GET /api/aliases/:alias_id: Get alias information.
func (c *Client) GetAlias(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// GET /api/aliases/:alias_id/activities: Get alias activities.
func (c *Client) GetAliasActivities(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// GET /api/aliases/:alias_id/contacts: Get alias contacts.
func (c *Client) GetAliasContacts(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// DELETE /api/aliases/:alias_id: Delete an alias.
func (c *Client) DeleteAlias(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// POST /api/v3/alias/custom/new: Create new alias.
func (c *Client) PostAlias(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// TODO: merge w/ PostAlias
// POST /api/alias/random/new: Random an alias.
// func (c *Client) PostLogin(ctx context.Context, options *AccountOptions) (*Client, error) {
//     return nil, nil
// }

// POST /api/aliases/:alias_id/toggle: Enable/disable an alias.
func (c *Client) PostToggleAlias(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// POST /api/aliases/:alias_id/contacts: Create a new contact for an alias.
func (c *Client) PostAliasContact(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// PATCH /api/aliases/:alias_id: Update alias information.
func (c *Client) PatchAlias(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}
