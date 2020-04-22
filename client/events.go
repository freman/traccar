package client

import (
	"fmt"
	"net/http"
	"time"
)

// Event describes an Event schema object
type Event struct {
	ID            int        `json:"id"`
	Type          string     `json:"type"`
	ServerTime    time.Time  `json:"serverTime"`
	DeviceID      int        `json:"deviceId"`
	PositionID    int        `json:"positionId"`
	GeofenceID    int        `jsson:"geofenceId"`
	MaintenanceID int        `json:"maintenanceId"`
	Attributes    Attributes `json:"attributes"`
}

// GetEvent is undocumented
func (c *Client) GetEvent(id int) (event Event, err error) {
	err = c.doRequest(http.MethodGet, fmt.Sprintf("events/%d", id), nil, &event)
	return
}
