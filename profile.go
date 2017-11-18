package xboxapi

import (
	"encoding/json"
	"fmt"
)

// Profile as returned by /v2/{xuid}/profile
type Profile struct {
	ID                int    `json:"id"`
	HostID            int    `json:"hostId"`
	GamerTag          string `json:"Gamertag"`
	GameDisplayName   string `json:"GameDisplayName"`
	AppDisplayName    string `json:"AppDisplayName"`
	GamerScore        int    `json:"Gamerscore"`
	GameDisplayPicRaw string `json:"GameDisplayPicRaw"`
	AccountTier       string `json:"AccountTier"`
	XboxOneRep        string `json:"XboxOneRep"`
	PreferredColor    string `json:"PreferredColor"`
	TenureLevel       int    `json:"TenureLevel"`
	IsSponsoredUser   bool   `json:"isSponsoredUser"`
}

// Profile fetches the account profile for a given XUID
func (c *Client) Profile(xuid int) (*Profile, error) {
	rsp, err := c.Get(fmt.Sprintf("https://xboxapi.com/v2/%d/profile", xuid))
	defer safeHTTPResponseClose(rsp)
	if err != nil {
		return nil, err
	}
	var data *Profile
	err = json.NewDecoder(rsp.Body).Decode(&data)
	return data, err
}
