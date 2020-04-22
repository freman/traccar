package client

import (
	"fmt"
	"net/http"
)

// Notification describes a Notification schema object
type Notification struct {
	ID         int        `json:"id"`
	Type       string     `json:"type"`
	Always     bool       `json:"always"`
	Web        bool       `json:"web"`
	Mail       bool       `json:"mail"`
	SMS        bool       `json:"sms"`
	CalendarID int        `json:"calendarId"`
	Attributes Attributes `json:"attributes"`
}

// NotificationType describes a NotificationType schema object
type NotificationType struct {
	Type string `json:"type"`
}

// ListNotifications fetches a list of Notifications
// Without params, it returns a list of Notifications the user has access to
//
// Params:
//   WithAll - Can only be used by admins or managers to fetch all entities
//   WithUserID - Standard users can use this only with their own userId
//   WithDeviceID - Standard users can use this only with _deviceId_s, they have access to
//   WithGroupID - Standard users can use this only with _groupId_s, they have access to
//   Refresh
func (c *Client) ListNotifications(params ...QueryParameter) (notifications []Notification, err error) {
	err = c.doRequest(http.MethodGet, "notifications"+query(params), nil, &notifications)
	return
}

// CreateNotification creates a Notification
func (c *Client) CreateNotification(notification Notification) (Notification, error) {
	return c.doNotification(http.MethodPost, "notifications", notification)
}

// TestNotification sends test notification to current user via Email and SMS
//
// Errors:
// 400 Bad Request - Could happen if sending has failed
func (c Client) TestNotification() error {
	return c.doRequest(http.MethodPost, "notifications/test", nil, nil)
}

// ListNotificationTypes fetches a list of available Notification types
func (c Client) ListNotificationTypes() (notificationTypes []NotificationType, err error) {
	err = c.doRequest(http.MethodGet, "notifications/types", nil, &notificationTypes)
	return
}

// DeleteNotification deletes a notification
func (c *Client) DeleteNotification(id int) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("notifications/%d", id), nil, nil)
}

// UpdateNotification updates a notification
func (c *Client) UpdateNotification(id int, notification Notification) (Notification, error) {
	notification.ID = id
	return c.doNotification(http.MethodPut, fmt.Sprintf("notifications/%d", id), notification)
}

func (c *Client) doNotification(method, path string, notification Notification) (Notification, error) {
	body, err := jsonBody(notification)
	if err != nil {
		return notification, err
	}

	err = c.doRequest(method, path, body, &notification)

	return notification, err
}
