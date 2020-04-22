package client

import (
	"fmt"
	"net/http"
	"time"
)

// User describes a User schema object
type User struct {
	ID               int        `json:"id"`
	Name             string     `json:"name"`
	Email            string     `json:"email"`
	Readonly         bool       `json:"readonly"`
	Administrator    bool       `json:"administrator"`
	Map              string     `json:"map"`
	Latitude         float64    `json:"latitude"`
	Longitude        float64    `json:"longitude"`
	Zoom             int        `json:"zoom"`
	Password         string     `json:"password"`
	TwelveHourFormat bool       `json:"twelveHourFormat"`
	CoordinateFormat string     `json:"coordinateFormat"`
	Disabled         bool       `json:"disabled"`
	ExpirationTime   time.Time  `json:"expirationTime"`
	DeviceLimit      int        `json:"deviceLimit"`
	UserLimit        int        `json:"userLimit"`
	DeviceReadonly   bool       `json:"deviceReadonly"`
	LimitCommands    bool       `json:"limitCommands"`
	PoiLayer         string     `json:"poiLayer"`
	Token            string     `json:"token"`
	Attributes       Attributes `json:"attributes"`
}

// ListUsers fetches a list of Users
//
// Params:
//   WithUserId - Can only be used by admin or manager users
//
// Errors:
//   400 Bad Request - No Permission
func (c *Client) ListUsers(params ...QueryParameter) (users []User, err error) {
	err = c.doRequest(http.MethodGet, "users"+query(params), nil, &users)
	return
}

// CreateUser creates a User
func (c *Client) CreateUser(user User) (User, error) {
	return c.doUser(http.MethodPost, "user", user)
}

// DeleteUser deletes a User
func (c *Client) DeleteUser(id int) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("users/%d", id), nil, nil)
}

// UpdateUser updates a user
func (c *Client) UpdateUser(id int, user User) (User, error) {
	user.ID = id
	return c.doUser(http.MethodPut, fmt.Sprintf("users/%d", id), user)
}

func (c *Client) doUser(method, path string, user User) (User, error) {
	body, err := jsonBody(user)
	if err != nil {
		return user, err
	}

	err = c.doRequest(method, path, body, &user)

	return user, err
}
