package xboxapi

import (
	"encoding/json"
	"fmt"
)

// GamerCard as returned by /v2/{xuid}/gamercard
type GamerCard struct {
	GamerTag                  string `json:"gamertag"`
	Name                      string `json:"name"`
	Location                  string `json:"location"`
	Bio                       string `json:"bio"`
	GamerScore                int    `json:"gamerscore"`
	Tier                      string `json:"tier"`
	Motto                     string `json:"motto"`
	AvatarBodyImagePath       string `json:"avatarBodyImagePath"`
	GamerpicSmallImagePath    string `json:"gamerpicSmallImagePath"`
	GamerpicLargeImagePath    string `json:"gamerpicLargeImagePath"`
	GamerpicSmallSslImagePath string `json:"gamerpicSmallSslImagePath"`
	GamerpicLargeSslImagePath string `json:"gamerpicLargeSslImagePath"`
	AvatarManifest            string `json:"avatarManifest"`
}

// GamerCard fetches the account GamerCard for a given XUID
func (c *Client) GamerCard(xuid int) (*GamerCard, error) {
	rsp, err := c.Get(fmt.Sprintf("https://xboxapi.com/v2/%d/gamercard", xuid))
	defer safeHTTPResponseClose(rsp)
	if err != nil {
		return nil, err
	}
	var data *GamerCard
	err = json.NewDecoder(rsp.Body).Decode(&data)
	return data, err
}
