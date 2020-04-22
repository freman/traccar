package client

import (
	"fmt"
	"net/http"
)

// Maintenance describes a Maintenance schema object
type Maintenance struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Start      int64      `json:"start"`
	Period     int64      `json:"period"`
	Attributes Attributes `json:"attributes"`
}

// ListMaintenances fetches a list of Maintenance
// Without params, it returns a list of Maintenance the user has access to
//
// Params:
//   WithAll - Can only be used by admins or managers to fetch all entities
//   WithUserID - Standard users can use this only with their own userId
//   WithDeviceID - Standard users can use this only with _deviceId_s, they have access to
//   WithGroupID - Standard users can use this only with _groupId_s, they have access to
//   Refresh
func (c *Client) ListMaintenances(params ...QueryParameter) (maintenances []Maintenance, err error) {
	err = c.doRequest(http.MethodGet, "maintenances"+query(params), nil, &maintenances)
	return
}

// CreateMaintenance creates a Maintenance
func (c *Client) CreateMaintenance(maintenance Maintenance) (Maintenance, error) {
	return c.doMaintenance(http.MethodPost, "maintenances", maintenance)
}

// DeleteMaintenance deletes a Maintenance
func (c *Client) DeleteMaintenance(id int) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("maintenances/%d", id), nil, nil)
}

// UpdateMaintenance updates a Maintenance
func (c *Client) UpdateMaintenance(id int, maintenance Maintenance) (Maintenance, error) {
	maintenance.ID = id
	return c.doMaintenance(http.MethodPut, fmt.Sprintf("maintenances/%d", id), maintenance)
}

func (c *Client) doMaintenance(method, path string, maintenance Maintenance) (Maintenance, error) {
	body, err := jsonBody(maintenance)
	if err != nil {
		return maintenance, err
	}

	err = c.doRequest(method, path, body, &maintenance)

	return maintenance, err
}
