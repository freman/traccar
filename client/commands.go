package client

import (
	"fmt"
	"net/http"
)

// Command describes a Command schema object
type Command struct {
	ID          int        `json:"id"`
	DeviceID    int        `json:"deviceId"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	SendSMS     bool       `json:"textChannel"`
	Attributes  Attributes `json:"attributes"`
}

// CommandType describes a CommandType schema object
type CommandType struct {
	Type string `json:"type"`
}

// ListCommands fetches a list of Saved Commands
// Without params, it returns a list of Commands the user has access to
//
// Params:
//   WithAll - Can only be used by admins or managers to fetch all entities
//   WithUserID - Standard users can use this only with their own userId
//   WithDeviceID - Standard users can use this only with _deviceId_s, they have access to
//   WithGroupID - Standard users can use this only with _groupId_s, they have access to
//   Refresh
func (c *Client) ListCommands(params ...QueryParameter) (commands []Command, err error) {
	err = c.doRequest(http.MethodGet, "commands"+query(params), nil, &commands)
	return
}

// CreateCommand creates a Saved Command
func (c *Client) CreateCommand(command Command) (Command, error) {
	return c.doCommand(http.MethodPost, "commands", command)
}

// ListDeviceCommands fetches a list of Saved Commands supported by Device at the moment
// This api makes no sense
//
// Errors:
//   400 Bad Request - Could happen when the user doesn't have permission for the device
func (c Client) ListDeviceCommands(deviceID int) (commands []Command, err error) {
	err = c.doRequest(http.MethodGet, fmt.Sprintf("commands/send?deviceId=%d", deviceID), nil, &commands)
	return
}

// SendCommand dispatches commands to device
// Dispatch a new command or Saved Command if Command.ID set
//
// Errors:
//   400 Bad Request - Could happen when the user doesn't have permission or an incorrect command type for the device
func (c *Client) SendCommand(command Command) (Command, error) {
	return c.doCommand(http.MethodPost, "commands/send", command)
}

// ListCommandTypes returns a list of available Commands for the Device or all
// possible Commands if Device omitted
//
// Params:
//   WithDeviceID
//   WithTextChannel
//
// Errors:
//  400 Bad Request - Could happen when trying to fetch from a device the user does not have permission
func (c Client) ListCommandTypes(params ...QueryParameter) (commandTypes []CommandType, err error) {
	err = c.doRequest(http.MethodGet, "commands/types"+query(params), nil, &commandTypes)
	return
}

// DeleteCommand deletes a saved command
func (c *Client) DeleteCommand(id int) error {
	return c.doRequest(http.MethodDelete, fmt.Sprintf("commands/%d", id), nil, nil)
}

// UpdateCommand updates a saved command
func (c *Client) UpdateCommand(id int, command Command) (Command, error) {
	command.ID = id
	return c.doCommand(http.MethodPut, fmt.Sprintf("commands/%d", id), command)
}

func (c *Client) doCommand(method, path string, command Command) (Command, error) {
	body, err := jsonBody(command)
	if err != nil {
		return command, err
	}

	err = c.doRequest(method, path, body, &command)

	return command, err
}
