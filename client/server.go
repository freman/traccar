package client

import (
	"net/http"
)

// Server describes a Server schema object
type Server struct {
	ID               int        `json:"id"`
	Registration     bool       `json:"registration"`
	Readonly         bool       `json:"readonly"`
	DeviceReadonly   bool       `json:"deviceReadonly"`
	LimitCommands    bool       `json:"limitCommands"`
	Map              string     `json:"map"`
	BingKey          string     `json:"bingKey"`
	MapURL           string     `json:"mapUrl"`
	PoiLayer         string     `json:"poiLayer"`
	Latitude         float64    `json:"latitude"`
	Longitude        float64    `json:"longitude"`
	Zoom             int        `json:"zoom"`
	TwelveHourFormat bool       `json:"twelveHourFormat"`
	Version          string     `json:"version"`
	ForceSettings    bool       `json:"forceSettings"`
	CoordinateFormat string     `json:"coordinateFormat"`
	Attributes       Attributes `json:"attributes"`
}

// GetServer fetches Server information
func (c *Client) GetServer() (server Server, err error) {
	err = c.doRequest(http.MethodGet, "server", nil, &server)
	return
}

// UpdateServer updates Server information
func (c *Client) UpdateServer(server Server) (Server, error) {
	body, err := jsonBody(server)
	if err != nil {
		return server, err
	}

	err = c.doRequest(http.MethodPut, "server", body, &server)

	return server, err
}
