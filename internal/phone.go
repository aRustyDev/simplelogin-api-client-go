package simplelogin_api_client_go

import "context"

type PhoneOptions struct {
	Id string `json:"reservation_id"`
}

// GET /api/phone/reservations/:reservation_id: Get messages received during a reservation.
func (c *Client) GetReservation(ctx context.Context, options *PhoneOptions) (*Client, error) {
	return nil, nil
}
