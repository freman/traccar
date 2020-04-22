package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// type Permissions struct {
// 	UserID        int `json:"userId,omitempty"`        // User Id, can be only first parameter
// 	DeviceID      int `json:"deviceId,omitempty`       // Device Id, can be first parameter or second only in combination with userId
// 	GroupID       int `json:"groupId,omitempty"`       // Group Id, can be first parameter or second only in combination with userId
// 	GeofenceID    int `json:"geofenceId,ompitempty"`   // Geofence Id, can be second parameter only
// 	CalendarID    int `json:"calendarId,omitempty"`    // Geofence Id, can be second parameter only and only in combination with userId
// 	AttributeID   int `json:"attributeId,omitempty"`   // Computed Attribute Id, can be second parameter only
// 	DriverID      int `json:"driverId,omitempty"`      // Driver Id, can be second parameter only
// 	ManagedUserID int `json:"managedUserId,omitempty"` // User Id, can be second parameter only and only in combination with userId
// }

// DeletePermission unlinks an Object from another Object
//
// Eg:
//   DeletePermission(PermitUserID(5).DeviceID(30))
func (c *Client) DeletePermission(request PermissionsRequest) error {
	return c.doPermissions(http.MethodDelete, request)
}

// CreatePermission links an Object to another Object
//
// Eg:
//   CreatePermission(PermitUserID(5).DeviceID(30))
func (c *Client) CreatePermission(request PermissionsRequest) error {
	return c.doPermissions(http.MethodPost, request)
}

func (c *Client) doPermissions(method string, request PermissionsRequest) error {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(request); err != nil {
		return err
	}

	return c.doRequest(method, "permissions", body, nil)
}

// PermissionsRequest is placeholder for the json object that's sent as
// apparently the order of the object fields is iportant.
type PermissionsRequest [2]struct {
	name  string
	value int
}

// PermitUserID will pair a UserID with whatever method you call on it
type PermitUserID int

// PermitDeviceID will pair a DeviceID with whatever method you call on it
type PermitDeviceID int

// PermitGroupID will pair a GroupID with whatever method you call on it
type PermitGroupID int

// DeviceID pairs the given DeviceID with the UserID
func (p PermitUserID) DeviceID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "userId", value: int(p)},
		{name: "deviceId", value: id},
	}
}

// GroupID pairs the given GroupID with the UserID
func (p PermitUserID) GroupID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "userId", value: int(p)},
		{name: "groupId", value: id},
	}
}

// GeofenceID pairs the given GeofenceID with the UserID
func (p PermitUserID) GeofenceID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "userId", value: int(p)},
		{name: "geofenceId", value: id},
	}
}

// CalendarID pairs the given CalendarID with the UserID
func (p PermitUserID) CalendarID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "userId", value: int(p)},
		{name: "calendarID", value: id},
	}
}

// AttributeID pairs the given AttributeID with the UserID
func (p PermitUserID) AttributeID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "userId", value: int(p)},
		{name: "attributeID", value: id},
	}
}

// DriverID pairs the given DriverID with the UserID
func (p PermitUserID) DriverID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "userId", value: int(p)},
		{name: "driverId", value: id},
	}
}

// ManagedUserID pairs the given ManagedUserID with the UserID
func (p PermitUserID) ManagedUserID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "userId", value: int(p)},
		{name: "managedUserID", value: id},
	}
}

// GeofenceID pairs the given GeofenceID with the DeviceID
func (p PermitDeviceID) GeofenceID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "deviceId", value: int(p)},
		{name: "geofenceId", value: id},
	}
}

// AttributeID pairs the given AttributeID with the DeviceID
func (p PermitDeviceID) AttributeID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "deviceId", value: int(p)},
		{name: "attributeId", value: id},
	}
}

// DriverID pairs the given DriverID with the DeviceID
func (p PermitDeviceID) DriverID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "deviceId", value: int(p)},
		{name: "driverId", value: id},
	}
}

// GeofenceID pairs the given GeofenceID with the GroupID
func (p PermitGroupID) GeofenceID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "groupId", value: int(p)},
		{name: "geofenceId", value: id},
	}
}

// AttributeID pairs the given AttributeID with the GroupID
func (p PermitGroupID) AttributeID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "groupId", value: int(p)},
		{name: "attributeId", value: id},
	}
}

// DriverID pairs the given DriverID with the GroupID
func (p PermitGroupID) DriverID(id int) PermissionsRequest {
	return PermissionsRequest{
		{name: "groupId", value: int(p)},
		{name: "driverId", value: id},
	}
}

// MarshalJSON converts the placeholder into json notation for sending
func (p PermissionsRequest) MarshalJSON() (result []byte, err error) {
	result = append(result, '{')

	for i, v := range p {
		var name, value []byte

		if i > 0 {
			result = append(result, ',', ' ')
		}

		if name, err = json.Marshal(v.name); err != nil {
			return
		}

		if value, err = json.Marshal(v.value); err != nil {
			return
		}

		result = append(result, name...)
		result = append(result, ':', ' ')
		result = append(result, value...)
	}

	result = append(result, '}')

	return
}
