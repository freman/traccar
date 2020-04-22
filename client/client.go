package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// Client is an API client for Traccar
type Client struct {
	c       *http.Client
	baseurl *url.URL
	token   string

	jar http.CookieJar
}

// Message is what's received via streaming sockets
type Message struct {
	Devices   []Device   `json:"devices"`
	Positions []Position `json:"positions"`
	Events    []Event    `json:"events"`
}

func (c *Client) startSession() error {
	_, err := c.FetchSession(c.token)
	return err
}

// New returns a new API Client for Traccar
func New(token string, opts ...Option) (*Client, error) {
	c := &Client{token: token}
	for _, opt := range opts {
		opt(c)
	}

	if c.baseurl == nil {
		return nil, errors.New("no default baseURL exists so the WithBaseURL is currently mandatory")
	}

	if c.c == nil {
		c.c = &http.Client{
			Timeout: time.Minute,
		}
	}

	if c.jar == nil {
		if c.c.Jar == nil {
			jar, err := cookiejar.New(&cookiejar.Options{})
			if err != nil {
				return nil, err
			}

			c.jar = jar
		} else {
			c.jar = c.c.Jar
		}
	}

	if c.c.Jar == nil {
		c.c.Jar = c.jar
	}

	if c.c.Jar != c.jar {
		return nil, errors.New("cookie jars don't match")
	}

	if err := c.startSession(); err != nil {
		return nil, err
	}

	return c, nil
}

// Stream messages from TRACCAR via websocket
func (c *Client) Stream(stop chan struct{}) (chan Message, error) {
	dialer := websocket.DefaultDialer
	dialer.Jar = c.jar

	wsURL := "ws" + strings.TrimPrefix(c.resolveURI("socket"), "http")
	wsc, _, err := dialer.Dial(wsURL, nil)
	if err != nil {
		return nil, err
	}

	messages := make(chan Message, 5)

	go func() {
		defer wsc.Close()

		for {
			select {
			case <-stop:
				return
			default:
				_, message, err := wsc.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					return
				}

				var m Message
				if err := json.Unmarshal(message, &m); err != nil {
					log.Println("decode:", err)
					continue
				}

				select {
				case messages <- m:
					// All is well
				default:
					// silently discard the message
				}
			}
		}
	}()

	return messages, nil
}

func (c *Client) resolveURI(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}

	return c.baseurl.ResolveReference(u).String()
}

func (c *Client) doRequest(method, uri string, body io.Reader, response interface{}) error {
	uri = c.resolveURI(uri)
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	if body != nil {
		if rb, isa := body.(requestBody); isa {
			req.Header.Set("Content-Type", rb.ContentType)
		}
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(resp.Status)
	}

	if response != nil {
		if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
			return err
		}
	}

	return nil
}

type requestBody struct {
	ContentType string
	Body        *bytes.Buffer
}

func (b requestBody) Read(p []byte) (n int, err error) {
	return b.Body.Read(p)
}

func jsonBody(v interface{}) (b requestBody, err error) {
	b.ContentType = "application/json"
	b.Body = new(bytes.Buffer)
	err = json.NewEncoder(b.Body).Encode(v)

	return
}

func formBody(v url.Values) (b requestBody) {
	b.ContentType = "application/x-www-form-urlencoded"
	b.Body = bytes.NewBufferString(v.Encode())

	return
}
