package xboxapi

import (
	"encoding/json"
	"fmt"
)

type XboxOneTitle struct {
	LastUnlock         APITime `json:"lastUnlock"`
	TitleID            int     `json:"titleId"`
	ServiceConfigID    string  `json:"serviceConfigId"`
	TitleType          string  `json:"titleType"`
	Platform           string  `json:"platform"`
	Name               string  `json:"name"`
	EarnedAchievements int     `json:"earnedAchievements"`
	CurrentGamerscore  int     `json:"currentGamerscore"`
	MaxGamerscore      int     `json:"maxGamerscore"`
}

// XboxOneGames is one page of games as returned by /v2/{xuid}/xboxonegames
type XboxOneGames struct {
	Titles     []XboxOneTitle `json:"titles"`
	PagingInfo struct {
		ContinuationToken *json.Number `json:"continuationToken"`
		TotalRecords      int          `json:"totalRecords"`
	} `json:"pagingInfo"`
}

func (c *Client) XboxOneGames(xuid int, continuationToken *json.Number) (*XboxOneGames, error) {
	var url string
	if continuationToken == nil {
		url = fmt.Sprintf("https://xboxapi.com/v2/%d/xboxonegames", xuid)
	} else {
		url = fmt.Sprintf("https://xboxapi.com/v2/%d/xbox360games?continuationToken=%s", xuid, continuationToken.String())
	}
	rsp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if err := rspError(rsp); err != nil {
		if isHTTPError(err) {
			return nil, nil
		}
		return nil, err
	}
	defer rsp.Body.Close()
	var data *XboxOneGames
	err = json.NewDecoder(rsp.Body).Decode(&data)
	return data, err
}

type Xbox360Title struct {
	LastPlayed          APITime `json:"lastPlayed"`
	CurrentAchievements int     `json:"currentAchievements"`
	CurrentGamerscore   int     `json:"currentGamerscore"`
	Sequence            int     `json:"sequence"`
	TitleID             int     `json:"titleId"`
	TitleType           int     `json:"titleType"`
	Platforms           []int   `json:"platforms"`
	Name                string  `json:"name"`
	TotalAchievements   int     `json:"totalAchievements"`
	TotalGamerscore     int     `json:"totalGamerscore"`
}

type Xbox360Games struct {
	Titles     []Xbox360Title `json:"titles"`
	PagingInfo struct {
		ContinuationToken *json.Number `json:"continuationToken"`
		TotalRecords      int          `json:"totalRecords"`
	} `json:"pagingInfo"`
	Version APITime `json:"version"`
}

func (c *Client) Xbox360Games(xuid int, continuationToken *json.Number) (*Xbox360Games, error) {
	var url string
	if continuationToken == nil {
		url = fmt.Sprintf("https://xboxapi.com/v2/%d/xbox360games", xuid)
	} else {
		url = fmt.Sprintf("https://xboxapi.com/v2/%d/xbox360games?continuationToken=%s", xuid, continuationToken.String())
	}
	rsp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if err := rspError(rsp); err != nil {
		if isHTTPError(err) {
			return nil, nil
		}
		return nil, err
	}
	defer rsp.Body.Close()
	var data *Xbox360Games
	err = json.NewDecoder(rsp.Body).Decode(&data)
	return data, err
}
