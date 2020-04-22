package client

import (
	"fmt"
	"net/http"
)

// Driver describes a Driver schema object
type Driver struct {
	ID         int        `json:"id"`
	Attributes Attributes `json:"attributes"`
	Name       string     `json:"name"`
	UniqueID   string     `json:"uniqueId"`
}

// ListDrivers fetches a list of Drivers
// Without params, it returns a list of Drivers the user has access to
//
// Params:
//   WithAll - Can only be used by admins or managers to fetch all entities
//   WithUserID - Standard users can use this only with their own userId
//   WithDeviceID - Standard users can use this only with _deviceId_s, they have access to
//   WithGroupID - Standard users can use this only with _groupId_s, they have access to
//   Refresh
func (c *Client) ListDrivers(params ...QueryParameter) (driverList []Driver, err error) {
	err = c.doRequest(http.MethodPost, "drivers"+query(params), nil, &driverList)
	return
}

// CreateDriver creates a Driver
func (c *Client) CreateDriver(driver Driver) (Driver, error) {
	return c.doDriver(http.MethodPost, "drivers", driver)
}

// DeleteDriver deletes a Driver
func (c *Client) DeleteDriver(id int) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("drivers/%d", id), nil, nil)
}

// UpdateDriver updates a Driver
func (c *Client) UpdateDriver(id int, driver Driver) (Driver, error) {
	driver.ID = id
	return c.doDriver(http.MethodPut, fmt.Sprintf("drivers/%d", id), driver)
}

func (c *Client) doDriver(method, path string, driver Driver) (Driver, error) {
	body, err := jsonBody(driver)
	if err != nil {
		return driver, err
	}

	err = c.doRequest(method, path, body, &driver)

	return driver, err
}
