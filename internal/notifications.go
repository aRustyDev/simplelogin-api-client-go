package simplelogin_api_client_go

import "context"

type NotificationOptions struct {
	Page string `json:"page"`
}

// GET /api/notifications: Get notifications.
func (c *Client) GetNotifications(ctx context.Context, options *NotificationOptions) (*Client, error) {
	return nil, nil
}

// POST /api/notifications/:notification_id: Mark as read a notification.
func (c *Client) PostNotification(ctx context.Context, options *NotificationOptions) (*Client, error) {
	return nil, nil
}
