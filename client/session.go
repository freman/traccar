package client

import (
	"net/http"
	"net/url"
)

// CloseSession closes the session
func (c *Client) CloseSession() error {
	return c.doRequest(http.MethodDelete, "session", nil, nil)
}

// FetchSession fetches session information from a token
//
// Errors:
//   404 Not Found - Not Found
func (c *Client) FetchSession(token string) (user User, err error) {
	err = c.doRequest(http.MethodGet, "session?token="+token, nil, &user)
	return
}

// CreateSession creates a new session from user credentials
//
// Errors:
//   401 Unauthorized - Unauthorized
func (c *Client) CreateSession(email, password string) (user User, err error) {
	body := formBody(url.Values{
		"email":    []string{email},
		"password": []string{password},
	})

	err = c.doRequest(http.MethodPost, "session", body, &user)

	return
}
