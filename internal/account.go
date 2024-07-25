package simplelogin_api_client_go

import (
	"context"
	"fmt"
	"net/http"
)

type AccountOptions struct {
	Device         string `json:"device"`
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
}

type UserInfo struct {
	Name                  string `json:"name"`
	CanCreateReverseAlias bool   `json:"can_create_reverse_alias"`
	ConnectedProton       string `json:"connected_proton_address"`
	IsPremium             bool   `json:"is_premium"`
	Email                 string `json:"email"`
	InTrial               bool   `json:"in_trial"`
	ProfilePictureUrl     string `json:"profile_picture_url"`
	MaxAliasCount         int    `json:"max_alias_free_plan"`
}

type CookieToken struct {
	Token string `json:"token"`
}

// POST /api/auth/login: Authentication
func (c *Client) PostLogin(ctx context.Context, options *AccountOptions) (*AuthBundle, error) {

	// Set the variables for the request
	reqType := "POST"
	path := fmt.Sprintf("%s%s", c.Auth.BaseURL, "/api/auth/login")

	// Create a new request object
	req, err := http.NewRequest(reqType, path, nil)
	if err != nil {
		return nil, err
	}

	// Add the context to the request
	req = req.WithContext(ctx)

	req.Header.Set("email", "")
	req.Header.Set("password", "")
	req.Header.Set("device", options.Device)
	fmt.Println(req.Header)

	// Make the request
	if err := c.NewRequest(req, &c.Auth); err != nil {
		return nil, err
	}

	return &c.Auth, nil
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
func (c *Client) PatchUser(ctx context.Context, options *AccountOptions) (*UserInfo, error) {
	return nil, nil
}

// DELETE /api/user: Delete the current user.
func (c *Client) DeleteUser(ctx context.Context, options *AccountOptions) (*Client, error) {
	return nil, nil
}

// GET /api/user_info: Get user's information.
func (c *Client) GetUserInfo(ctx context.Context, options *AccountOptions) (*UserInfo, error) {

	// Set the variables for the request
	reqType := "GET"
	path := fmt.Sprintf("%s%s", c.Auth.BaseURL, "/api/user_info")
	ui := UserInfo{}

	// Create a new request object
	req, err := http.NewRequest(reqType, path, nil)
	if err != nil {
		return nil, err
	}

	// Add the context to the request
	req = req.WithContext(ctx)

	// Set Request Headers
	req.Header.Set("Authentication", c.Auth.ApiKey)

	// Make the request
	if err := c.NewRequest(req, &ui); err != nil {
		return nil, err
	}

	return &ui, nil
}

// GET /api/user/cookie_token: Get a one time use token to exchange it for a valid cookie
func (c *Client) GetCookie(ctx context.Context, options *AccountOptions) (*Client, error) {

	// Set the variables for the request
	reqType := "GET"
	path := fmt.Sprintf("%s%s", c.Auth.BaseURL, "/api/user/cookie_token")
	ui := CookieToken{}

	// Create a new request object
	req, err := http.NewRequest(reqType, path, nil)
	if err != nil {
		return nil, err
	}

	// Add the context to the request
	req = req.WithContext(ctx)

	// Set Request Headers
	req.Header.Set("Authentication", c.Auth.ApiKey)

	// Make the request
	if err := c.NewRequest(req, &ui); err != nil {
		return nil, err
	}

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
