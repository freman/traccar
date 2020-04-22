package client

import (
	"net/http"
	"time"
)

// ReportStops describes a ReportStops schema object
type ReportStops struct {
	DeviceID    int       `json:"deviceId"`
	DeviceName  string    `json:"deviceName"`
	Duration    int       `json:"duration"`
	StartTime   time.Time `json:"startTime"`
	Address     string    `json:"address"`
	Lat         float64   `json:"lat"`
	Lon         float64   `json:"lon"`
	EndTime     time.Time `json:"end"`
	SpentFuel   float64   `json:"spentFuel"`
	EngineHours int       `json:"engineHours"`
}

// ReportSummary describes a ReportSummary schema object
type ReportSummary struct {
	DeviceID     int     `json:"deviceId"`
	DeviceName   string  `json:"deviceName"`
	MaxSpeed     float64 `json:"maxSpeed"`
	AverageSpeed float64 `json:"averageSpeed"`
	Distance     float64 `json:"distance"`
	SpentFuel    float64 `json:"spentFuel"`
	EngineHours  int     `json:"engineHours"`
}

// ReportTrips describes a ReportTrips schema object
type ReportTrips struct {
	DeviceID       int       `json:"deviceId"`
	DeviceName     string    `json:"deviceName"`
	MaxSpeed       float64   `json:"maxSpeed"`
	AverageSpeed   float64   `json:"averageSpeed"`
	Distance       float64   `json:"distance"`
	SpentFuel      float64   `json:"spentFuel"`
	Duration       int       `json:"duration"`
	StartTime      time.Time `json:"startTime"`
	StartAddress   string    `json:"startAddress"`
	StartLat       float64   `json:"startLat"`
	StartLon       float64   `json:"startLon"`
	EndTime        time.Time `json:"endTime"`
	EndAddress     string    `json:"endAddress"`
	EndLat         float64   `json:"endLat"`
	EndLon         float64   `json:"endLon"`
	DriverUniqueID int       `json:"driverUniqueId"`
	DriverName     string    `json:"driverName"`
}

// GetReportEvents Fetch a list of Events within the time period for the Devices or Groups
// At least one deviceId or one groupId must be passed
//
// Params:
//   WithDeviceID - multiple parameters can be passed
//   WithGroupID -  multiple parameters can be passed
// type not implemented
//   WithFrom
//   WithTo
func (c *Client) GetReportEvents(params ...QueryParameter) (events []Event, err error) {
	err = c.doRequest(http.MethodGet, "report/events"+query(params), nil, &events)
	return
}

// GetReportRoute Fetch a list of Positions within the time period for the Devices or Groups
// At least one deviceId or one groupId must be passed
//
// Params:
//   WithDeviceID - multiple parameters can be passed
//   WithGroupID -  multiple parameters can be passed
// type not implemented
//   WithFrom
//   WithTo
func (c *Client) GetReportRoute(params ...QueryParameter) (route []Position, err error) {
	err = c.doRequest(http.MethodGet, "reports/route"+query(params), nil, &route)
	return
}

// GetReportStops Fetch a list of ReportStops within the time period for the Devices or Groups
// At least one deviceId or one groupId must be passed
//
// Params:
//   WithDeviceID - multiple parameters can be passed
//   WithGroupID -  multiple parameters can be passed
// type not implemented
//   WithFrom
//   WithTo
func (c *Client) GetReportStops(params ...QueryParameter) (stops []ReportStops, err error) {
	err = c.doRequest(http.MethodGet, "reports/stops"+query(params), nil, &stops)
	return
}

// GetReportSummary Fetch a list of ReportSummary within the time period for the Devices or Groups
// At least one deviceId or one groupId must be passed
//
// Params:
//   WithDeviceID - multiple parameters can be passed
//   WithGroupID -  multiple parameters can be passed
// type not implemented
//   WithFrom
//   WithTo
func (c *Client) GetReportSummary(params ...QueryParameter) (summary []ReportSummary, err error) {
	err = c.doRequest(http.MethodGet, "reports/summary"+query(params), nil, &summary)
	return
}

// GetReportTrips Fetch a list of ReportTrips within the time period for the Devices or Groups
// At least one deviceId or one groupId must be passed
//
// Params:
//   WithDeviceID - multiple parameters can be passed
//   WithGroupID -  multiple parameters can be passed
// type not implemented
//   WithFrom
//   WithTo
func (c *Client) GetReportTrips(params ...QueryParameter) (trips []ReportTrips, err error) {
	err = c.doRequest(http.MethodGet, "reports/trips"+query(params), nil, &trips)
	return
}
