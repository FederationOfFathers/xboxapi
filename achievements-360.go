package xboxapi

import (
	"encoding/json"
	"fmt"
)

type Achievement360 struct {
	ID                int     `json:"id"`
	TitleID           int     `json:"titleId"`
	Name              string  `json:"name"`
	Sequence          int     `json:"sequence"`
	Flags             int     `json:"flags"`
	UnlockedOnline    bool    `json:"unlockedOnline"`
	Unlocked          bool    `json:"unlocked"`
	Secret            bool    `json:"isSecret"`
	Platform          int     `json:"platform"`
	GamerScore        int     `json:"gamerscore"`
	ImageID           int     `json:"imageId"`
	Description       string  `json:"description"`
	LockedDescription string  `json:"lockedDescription"`
	Type              int     `json:"type"`
	Revoked           bool    `json:"isRevoked"`
	TimeUnlocked      APITime `json:"timeUnlocked"`
	Image             string  `json:"imageUnlocked"`
}

type Achievement360List []*Achievement360

func (c *Client) Achievements360(xuid string, titleID int) (Achievement360List, error) {
	rsp, err := c.Get(fmt.Sprintf("https://xboxapi.com/v2/%s/achievements/%d", xuid, titleID))
	defer safeHTTPResponseClose(rsp)
	if err != nil {
		return nil, err
	}
	if err := rspError(rsp); err != nil {
		if isHTTPError(err) {
			return nil, nil
		}
		return nil, err
	}
	var rval = Achievement360List{}
	err = json.NewDecoder(rsp.Body).Decode(&rval)
	return rval, err
}
