package xboxapi

import (
	"encoding/json"
	"fmt"
)

// Presence as returned by /v2/{xuid}/presence
type Presence struct {
	XUID     int    `json:"xuid"`
	State    string `json:"state"`
	LastSeen struct {
		DeviceType string  `json:"deviceType"`
		TitleID    int     `json:"titleId"`
		TitleName  string  `json:"titleName"`
		TimeStamp  APITime `json:"timestamp"`
	} `json:"lastSeen"`
}

// Presence fetches the presence information for a given XUID
func (c *Client) Presence(xuid int) (*Presence, error) {
	rsp, err := c.Get(fmt.Sprintf("https://xboxapi.com/v2/%d/presence", xuid))
	defer safeHTTPResponseClose(rsp)
	defer rsp.Body.Close()
	var data *Presence
	err = json.NewDecoder(rsp.Body).Decode(&data)
	return data, err
}
