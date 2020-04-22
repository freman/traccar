package client

import (
	"fmt"
	"net/http"
)

// Calendar describes a calendar schema object
//
// Data is the base64 encoded ical
type Calendar struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Data       string     `json:"data"`
	Attributes Attributes `json:"attributes"`
}

// ListCalendars fetches a list of Calendars
// Without params, it returns a list of Calendars the user has access to
//
// Params:
//   WithAll - Can only be used by admins or managers to fetch all entities
//   WithUserID - Standard users can use this only with their own userId
func (c *Client) ListCalendars(params ...QueryParameter) (calendars []Calendar, err error) {
	err = c.doRequest(http.MethodGet, "calendars"+query(params), nil, &calendars)
	return
}

// CreateCalendar creates a Calendar
func (c *Client) CreateCalendar(calendar Calendar) (Calendar, error) {
	return c.doCalendar(http.MethodPost, "calendars", calendar)
}

// DeleteCalendar deletes a Calendar
func (c *Client) DeleteCalendar(id int) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("calendars/%d", id), nil, nil)
}

// UpdateCalendar updates a Calendar
func (c *Client) UpdateCalendar(id int, calendar Calendar) (Calendar, error) {
	calendar.ID = id
	return c.doCalendar(http.MethodPut, fmt.Sprintf("calendars/%d", id), calendar)
}

func (c *Client) doCalendar(method, path string, calendar Calendar) (Calendar, error) {
	body, err := jsonBody(calendar)
	if err != nil {
		return calendar, err
	}

	err = c.doRequest(method, path, body, &calendar)

	return calendar, err
}
