package simplelogin_api_client_go

import (
	"context"
	"fmt"
	"net/http"
	"path"
)

type AccountOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// POST /api/auth/login: Authentication
func (c *Client) PostLogin(ctx context.Context, options *AccountOptions) (*AuthBundle, error) {

	// Set the variables for the request
	path := path.Join(c.BaseURL, "/api/auth/login")
	res := AuthBundle{}

	// Create a new request object
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.BaseURL, path), nil)
	if err != nil {
		return nil, err
	}

	// Add the context to the request
	req = req.WithContext(ctx)

	// Make the request
	if err := c.NewRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// POST /api/auth/mfa: 2FA authentication
func (c *Client) PostMFA(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// POST /api/auth/register: Register a new account.
func (c *Client) PostRegister(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// POST /api/auth/activate: Activate new account.
func (c *Client) PostActivate(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// POST /api/auth/reactivate: Request a new activation code.
func (c *Client) PostReactivate(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// POST /api/auth/forgot_password: Request reset password link.
func (c *Client) PostForgotPassword(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// POST /api/api_key: Create a new API key.
func (c *Client) PostApiKey(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// PATCH /api/sudo: Enable sudo mode.
func (c *Client) PatchSudo(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// PATCH /api/user_info: Update user's information.
func (c *Client) PatchUser(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// DELETE /api/user: Delete the current user.
func (c *Client) DeleteUser(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// GET /api/user_info: Get user's information.
func (c *Client) GetUserInfo(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// GET /api/user/cookie_token: Get a one time use token to exchange it for a valid cookie
func (c *Client) GetCookie(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// GET /api/stats: Get user's stats.
func (c *Client) GetUserStats(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// GET /api/logout: Log out.
func (c *Client) Logout(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}
