package client

import (
	"fmt"
	"net/http"
)

// Group describes a Group schema object
type Group struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	GroupID    int        `json:"groupId"`
	Attributes Attributes `json:"attributes"`
}

// ListGroups fetches a list of Groups
// Without any params, returns a list of the Groups the user belongs to
//
// Params:
//   WithAll - Can only be used by admins or managers to fetch all entities
//   WithUserID - Standard users can use this only with their own userId
func (c *Client) ListGroups(params ...QueryParameter) (groups []Group, err error) {
	err = c.doRequest(http.MethodGet, "groups"+query(params), nil, &groups)
	return
}

// CreateGroup creates a Group
//
// Errors:
//   400 Bad Request - No permission
func (c *Client) CreateGroup(group Group) (Group, error) {
	return c.doGroup(http.MethodPost, "groups", group)
}

// DeleteGroup deletes a group
func (c *Client) DeleteGroup(id int) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("groups/%d", id), nil, nil)
}

// UpdateGroup updates a group
func (c *Client) UpdateGroup(id int, group Group) (Group, error) {
	group.ID = id
	return c.doGroup(http.MethodPut, fmt.Sprintf("groups/%d", id), group)
}

func (c *Client) doGroup(method, path string, group Group) (Group, error) {
	body, err := jsonBody(group)
	if err != nil {
		return group, err
	}

	err = c.doRequest(method, path, body, &group)

	return group, err
}
