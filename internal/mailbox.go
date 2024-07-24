package simplelogin_api_client_go

import "context"

type MailboxOptions struct {
	Id                string `json:"mailbox_id"`
	TrasferTo         string `json:"transfer_aliases_to"`
	Default           bool   `json:"default"`
	CancelEmailChange string `json:"cancel_email_change"`
}

// POST /api/mailboxes: Create a new mailbox.
func (c *Client) PostMailbox(ctx context.Context, options *MailboxOptions) (*Client, error) {
	return nil, nil
}

// DELETE /api/mailboxes/:mailbox_id: Delete a mailbox.
func (c *Client) DeleteMailbox(ctx context.Context, options *MailboxOptions) (*Client, error) {
	return nil, nil
}

// PUT /api/mailboxes/:mailbox_id: Update a mailbox.
func (c *Client) PutMailbox(ctx context.Context, options *MailboxOptions) (*Client, error) {
	return nil, nil
}
