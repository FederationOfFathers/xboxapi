package xboxapi

import (
	"encoding/json"
	"fmt"
)

type TilehubTitle struct {
	TitleID               string   `json:"titleId"`
	PFN                   *string  `json:"pfn"`
	BingID                string   `json:"bingId"`
	WindowsPhoneProductID *string  `json:"windowsPhoneProductId"`
	Name                  string   `json:"name"`
	Type                  string   `json:"type"`
	Devices               []string `json:"devices"`
	DisplayImage          string   `json:"displayImage"`
	MediaItemType         string   `json:"mediaItemType"`
	ModernTitleID         string   `json:"modernTitleId"`
	IsBundle              bool     `json:"isBundle"`
	Achievement           struct {
		CurrentAchievements int `json:"currentAchievements"`
		TotalAchievements   int `json:"totalAchievements"`
		CurrentGamerscore   int `json:"currentGamerscore"`
		TotalGamerscore     int `json:"totalGamerscore"`
		ProgressPercentage  int `json:"progressPercentage"`
		SourceVersion       int `json:"sourceVersion"`
	} `json:"achievement"`
	Images       interface{} `json:"images"`
	TitleHistory struct {
		LastTimePlayed APITime `json:"lastTimePlayed"`
		Visible        bool    `json:"visible"`
		CanHide        bool    `json:"canHide"`
	} `json:""`
	Detail           interface{} `json:"detail"`
	FriendsWhoPlayed struct {
		CurrentlyPlayingCount int `json:"currentlyPlayingCount"`
		HavePlayedCount       int `json:"havePlayedCount"`
		People                []struct {
			XUID               string  `json:"xuid"`
			IsFavorite         bool    `json:"isFavorite"`
			DisplayPicRaw      string  `json:"displayPicRaw"`
			UseAvatar          bool    `json:"useAvatar"`
			IsCurrentlyPlaying bool    `json:"isCurrentlyPlaying"`
			Gamertag           string  `json:"gamertag"`
			LastTimePlayed     APITime `json:"lastTimePlayed"`
			PresenceState      string  `json:"presenceState"`
			PreferredColor     struct {
				PrimaryColor   string `json:"primaryColor"`
				PecondaryColor string `json:"secondaryColor"`
				PertiaryColor  string `json:"tertiaryColor"`
			} `json:"preferredColor"`
		} `json:"people"`
	} `json:"friendsWhoPlayed"`
}

type TilehubAchievementList struct {
	XUID   string          `json:"xuid"`
	Titles []*TilehubTitle `json:"titles"`
}

func (c *Client) TileHub(xuid string) (*TilehubAchievementList, error) {
	rsp, err := c.Get(fmt.Sprintf("https://xboxapi.com/v2/%s/titlehub-achievement-list", xuid))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if err := rspError(rsp); err != nil {
		if isHTTPError(err) {
			return nil, nil
		}
		return nil, err
	}
	var rval *TilehubAchievementList
	err = json.NewDecoder(rsp.Body).Decode(&rval)
	return rval, err
}
