package xboxapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var ErrUnknownAchievementType = fmt.Errorf("Unknown Achievement Type")

type Achievement struct {
	ID                int       `json:"id"`
	TitleID           int       `json:"title_id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	LockedSescription string    `json:"locked_description"`
	Unlocked          bool      `json:"unlocked"`
	Secret            bool      `json:"secret"`
	Image             string    `json:"image"`
	TimeUnlocked      time.Time `json:"time_unlocked"`
}

func (a *Achievement) Ingest(any interface{}) error {
	if a == nil {
		a = &Achievement{}
	}
	switch any := any.(type) {
	case AchievementOne:
		a.ID = any.ID
		a.TitleID = any.TitleAssociations[0].ID
		a.Name = any.Name
		a.Description = any.Description
		a.LockedSescription = any.LockedDescription
		a.Secret = any.IsSecret
		for _, i := range any.MediaAssets {
			if i.Type == "Icon" {
				a.Image = i.URL
				break
			}
		}
		if strings.ToLower(any.ProgressState) == "achieved" {
			a.Unlocked = true
			a.TimeUnlocked = any.Progression.TimeUnlocked.Time()
		}
	case Achievement360:
		a.ID = any.ID
		a.TitleID = any.TitleID
		a.Name = any.Name
		a.Description = any.Description
		a.LockedSescription = any.LockedDescription
		a.Secret = any.Secret
		a.Image = any.Image
		if any.Unlocked {
			a.Unlocked = any.Unlocked
			a.TimeUnlocked = any.TimeUnlocked.Time()
		}
	default:
		return ErrUnknownAchievementType
	}
	return nil
}

type AchievementList []*Achievement

func (c *Client) Achievements(xuid string, titleID int) (AchievementList, error) {
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
	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	var rval = AchievementList{}
	if bytes.Contains(buf, []byte(`"serviceConfigId"`)) {
		var list = []AchievementOne{}
		if err := json.Unmarshal(buf, &list); err != nil {
			return nil, err
		}
		for _, cheevo := range list {
			var entry = &Achievement{}
			if err := entry.Ingest(cheevo); err != nil {
				return nil, err
			}
			rval = append(rval, entry)
		}
	} else {
		var list = []Achievement360{}
		if err := json.Unmarshal(buf, &list); err != nil {
			return nil, err
		}
		for _, cheevo := range list {
			var entry = &Achievement{}
			if err := entry.Ingest(cheevo); err != nil {
				return nil, err
			}
			rval = append(rval, entry)
		}
	}
	return rval, err
}
