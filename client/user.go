package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// LoginUserPath computes a request path to the login action of user.
func LoginUserPath(username string) string {
	return fmt.Sprintf("/u/%v", username)
}

// Login
func (c *Client) LoginUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewLoginUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginUserRequest create the request corresponding to the login action endpoint of the user resource.
func (c *Client) NewLoginUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
