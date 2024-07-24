package simplelogin_api_client_go

import "context"

type ContactOptions struct {
	Id string `json:"contact_id"`
}

// DELETE /api/contacts/:contact_id: Delete a contact.
func (c *Client) DeleteContact(ctx context.Context, options *ContactOptions) (*Client, error) {
	return nil, nil
}

// POST /api/contacts/:contact_id/toggle: Block/unblock a contact.
func (c *Client) PostToggleContact(ctx context.Context, options *ContactOptions) (*Client, error) {
	return nil, nil
}
