package client

import (
	"fmt"
	"net/http"
)

// ComputedAttribute describes a ComputedAttribute schema Object
type ComputedAttribute struct {
	ID          int                   `json:"id"`
	Description string                `json:"description"`
	Attribute   string                `json:"attribute"`
	Expression  string                `json:"expression"`
	Type        ComputedAttributeType `json:"type"`
}

// ComputedAttributeType is a string type of ComputedAttribute Type
type ComputedAttributeType string

// Computed attribute types
const (
	StringComputedAttribute  ComputedAttributeType = "String"
	NumberComputedAttribute  ComputedAttributeType = "Number"
	BooleanComputedAttribute ComputedAttributeType = "Boolean"
)

// ListComputedAttributes fetches a list of Attributes
// Without params, it returns a list of Attributes the user has access to
//
// Params:
//   WithAll - Can only be used by admins or managers to fetch all entities
//   WithUserID - Standard users can use this only with their own userId
//   WithDeviceID - Standard users can use this only with _deviceId_s, they have access to
//   WithGroupID - Standard users can use this only with _groupId_s, they have access to
//   Refresh
func (c *Client) ListComputedAttributes(params ...QueryParameter) (attributeList []ComputedAttribute, err error) {
	err = c.doRequest(http.MethodGet, "attributes/computed"+query(params), nil, &attributeList)
	return
}

// CreateComputedAttribute creates an Attribute
func (c *Client) CreateComputedAttribute(attribute ComputedAttribute) (ComputedAttribute, error) {
	return c.doComputedAttribute(http.MethodPost, "attributes/computed", attribute)
}

// DeleteComputedAttribute deletes an Attribute
func (c *Client) DeleteComputedAttribute(id int) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("attributes/computed/%d", id), nil, nil)
}

// UpdateComputedAttribute updates an Attribute
func (c *Client) UpdateComputedAttribute(id int, attribute ComputedAttribute) (ComputedAttribute, error) {
	attribute.ID = id
	return c.doComputedAttribute(http.MethodPut, fmt.Sprintf("attributes/computed/%d", id), attribute)
}

func (c *Client) doComputedAttribute(method, path string, attribute ComputedAttribute) (ComputedAttribute, error) {
	body, err := jsonBody(attribute)
	if err != nil {
		return attribute, err
	}

	err = c.doRequest(method, path, body, &attribute)

	return attribute, err
}
