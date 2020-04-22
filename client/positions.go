package client

import (
	"net/http"
	"time"
)

// Position describes a Position schema object
type Position struct {
	ID         int64      `json:"id"`
	Attributes Attributes `json:"attributes"`
	DeviceID   int64      `json:"deviceId"`
	Protocol   string     `json:"protocol"`
	ServerTime time.Time  `json:"serverTime"`
	DeviceTime time.Time  `json:"deviceTime"`
	FixTime    time.Time  `json:"fixTime"`
	Outdated   bool       `json:"outdated"`
	Valid      bool       `json:"valid"`
	Latitude   float64    `json:"latitude"`
	Longitude  float64    `json:"longitude"`
	Altitude   float64    `json:"altitude"`
	Speed      float64    `json:"speed"`
	Course     float64    `json:"course"`
	Address    string     `json:"address"`
	Accuracy   float64    `json:"accuracy"`
	Network    Attributes `json:"network"`
}

// ListPositions fetches a list of Positions
// Without any params, it returns a list of last known positions for all
// the user's Devices. from and _to_ fields are not required with _id_
//
// Params:
//   WithDeviceID - deviceId is optional, but requires the from and _to_ parameters when used
//   WithFrom
//   WithTo
//   WithID - To fetch one or more positions. Multiple params can be passed
func (c *Client) ListPositions(params ...QueryParameter) (positions []Position, err error) {
	err = c.doRequest(http.MethodGet, "positions"+query(params), nil, &positions)
	return
}
