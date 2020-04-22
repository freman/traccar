package client

import (
	"net/http"
	"net/url"
)

// Option is a function that operates on a Client object
type Option func(c *Client)

// WithBaseURL permits you to pass a base API url to the client consturctor
func WithBaseURL(uri string) Option {
	return func(c *Client) {
		u, err := url.Parse(uri)
		if err != nil {
			panic(err)
		}

		c.baseurl = u
	}
}

// WithHTTPClient permits you to pass a non-standard or otherwise tinkered
// with http.Client object to the client consturctor
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.c = client
	}
}

// WithCookieJar permits you to pass a given cookie jar to the client constructor
// this is useful if you would prefer to save your cookies to disk
// Note that if you set this cookie jar and pass one WithHTTPClient than you
// will return an error if they're not the same jar
func WithCookieJar(jar http.CookieJar) Option {
	return func(c *Client) {
		c.jar = jar
	}
}
