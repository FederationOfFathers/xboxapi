package xboxapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Account returns the XUID (int) and Gamertag (strong) of the xbox api account
func (c *Client) Account() (int, string, error) {
	rsp, err := c.Get("https://xboxapi.com/v2/accountXuid")
	defer safeHTTPResponseClose(rsp)
	if err := rspError(rsp); err != nil {
		if isHTTPError(err) {
			return 0, "", nil
		}
		return 0, "", err
	}
	var data = &struct {
		XUID     int    `json:"xuid"`
		GamerTag string `json:"gamerTag"`
	}{}
	err = json.NewDecoder(rsp.Body).Decode(&data)
	return data.XUID, data.GamerTag, err
}

// XUID returns the XUID for a given GamerTag
func (c *Client) XUID(gt string) (int, error) {
	rsp, err := c.Get(fmt.Sprintf("https://xboxapi.com/v2/xuid/%s", gt))
	defer safeHTTPResponseClose(rsp)
	if err != nil {
		return 0, err
	}
	if err := rspError(rsp); err != nil {
		if err == ErrNotFound {
			return 0, nil
		}
	}
	var data int
	err = json.NewDecoder(rsp.Body).Decode(&data)
	return data, err
}

// GamerTag returns the GamerTag for a given XUID
func (c *Client) GamerTag(xuid int) (string, error) {
	rsp, err := c.Get(fmt.Sprintf("https://xboxapi.com/v2/gamertag/%d", xuid))
	defer safeHTTPResponseClose(rsp)
	if err != nil {
		return "", err
	}
	if err := rspError(rsp); err != nil {
		if err == ErrNotFound {
			return "", nil
		}
	}
	buf, err := ioutil.ReadAll(rsp.Body)
	return string(buf), err
}
