package client

import (
	"net/http"
	"time"
)

// Statistics describes a Statistics schema object
type Statistics struct {
	CaptureTime      time.Time `json:"captureTime"`
	ActiveUsers      int       `json:"activeUsers"`
	ActiveDevices    int       `json:"activeDevices"`
	Requests         int       `json:"requests"`
	MessagesReceived int       `json:"messagesReceived"`
	MessagesStored   int       `json:"messagesStored"`
}

// GetStatistics fetches server statistics
//
// Params:
//   WithFrom
//   WithTo
func (c *Client) GetStatistics(params ...QueryParameter) (statistics []Statistics, err error) {
	err = c.doRequest(http.MethodGet, "statistics"+query(params), nil, &statistics)
	return
}
