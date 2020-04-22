package client

import (
	"fmt"
	"net/http"
)

// Geofence describes a Geofence schema object
type Geofence struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Area        string     `json:"area"`
	CalendarID  int        `json:"calendarId"`
	Attributes  Attributes `json:"attributes"`
}

// ListGeofences fetches a list of Geofences
// Without params, it returns a list of Geofences the user has access to
//
// Params:
//   WithAll - Can only be used by admins or managers to fetch all entities
//   WithUserID - Standard users can use this only with their own userId
//   WithDeviceID - Standard users can use this only with _deviceId_s, they have access to
//   WithGroupID - Standard users can use this only with _groupId_s, they have access to
//   Refresh
func (c *Client) ListGeofences(params ...QueryParameter) (geofences []Geofence, err error) {
	err = c.doRequest(http.MethodGet, "geofences"+query(params), nil, &geofences)
	return
}

// CreateGeofence creates a Geofence
func (c *Client) CreateGeofence(geofence Geofence) (Geofence, error) {
	return c.doGeofence(http.MethodPost, "geofences", geofence)
}

// DeleteGeofence deletes a Geofence
func (c *Client) DeleteGeofence(id int) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("geofences/%d", id), nil, nil)
}

// UpdateGeofence updates a Geofence
func (c *Client) UpdateGeofence(id int, geofence Geofence) (Geofence, error) {
	geofence.ID = id
	return c.doGeofence(http.MethodPut, fmt.Sprintf("geofence/%d", id), geofence)
}

func (c *Client) doGeofence(method, path string, geofence Geofence) (Geofence, error) {
	body, err := jsonBody(geofence)
	if err != nil {
		return geofence, err
	}

	err = c.doRequest(method, path, body, &geofence)

	return geofence, err
}
