package client

import (
	"fmt"
	"net/http"
)

// Device describes a Device schema object
type Device struct {
	ID          int           `json:"id"`
	Attributes  Attributes    `json:"attributes"`
	GroupID     int           `json:"groupId"`
	Name        string        `json:"name"`
	UniqueID    string        `json:"uniqueId"`
	Status      string        `json:"status"`
	LastUpdate  string        `json:"lastUpdate"`
	PositionID  int           `json:"positionId"`
	GeofenceIds []interface{} `json:"geofenceIds"`
	Phone       string        `json:"phone"`
	Model       string        `json:"model"`
	Contact     string        `json:"contact"`
	Category    string        `json:"category"`
	Disabled    bool          `json:"disabled"`
}

// DeviceAccumulators describes a DeviceAccumulators object
type DeviceAccumulators struct {
	DeviceID      int   `json:"deviceId"`
	TotalDistance int64 `json:"totalDistance"`
	Hours         int64 `json:"hours"`
}

// ListDevices fetches a list of Devices
// Without any params, returns a list of the user's devices
//
// Params:
//   WithAll - Can only be used by admins or managers to fetch all entities
//   WithUserID - Standard users can use this only with their own userId
//   WithID - To fetch one or more devices. Multiple params can be passed
//   WithUniqueID - To fetch one or more devices. Multiple params can be passed
//
// Errors:
//   400 Bad Request - No permission
func (c *Client) ListDevices(params ...QueryParameter) (deviceList []Device, err error) {
	err = c.doRequest(http.MethodPost, "devices"+query(params), nil, &deviceList)
	return
}

// CreateDevice creates a Device
func (c *Client) CreateDevice(device Device) (Device, error) {
	return c.doDevice(http.MethodPost, "devices", device)
}

// DeleteDevice deletes a Device
func (c *Client) DeleteDevice(id int) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("devices/%d", id), nil, nil)
}

// UpdateDevice updates a Device
func (c *Client) UpdateDevice(id int, device Device) (Device, error) {
	device.ID = id
	return c.doDevice(http.MethodPut, fmt.Sprintf("devices/%d", id), device)
}

// UpdateDeviceAccumulators updates total distance and hours of the Device
func (c *Client) UpdateDeviceAccumulators(id int, accumulators DeviceAccumulators) error {
	accumulators.DeviceID = id
	body, err := jsonBody(accumulators)
	if err != nil {
		return err
	}

	return c.doRequest(http.MethodPut, fmt.Sprintf("devices/%d/accumulators", id), body, nil)
}

func (c *Client) doDevice(method, path string, device Device) (Device, error) {
	body, err := jsonBody(device)
	if err != nil {
		return device, err
	}

	err = c.doRequest(method, path, body, &device)

	return device, err
}
