package xboxapi

import (
	"encoding/json"
	"fmt"
	"time"
)

// XboxOneGames is one page of games as returned by /v2/{xuid}/xboxonegames
type XboxOneGames struct {
	Titles []struct {
		LastUnlock         time.Time `json:"lastUnlock"`
		TitleID            int       `json:"titleId"`
		ServiceConfigID    string    `json:"serviceConfigId"`
		TitleType          string    `json:"titleType"`
		Platform           string    `json:"platform"`
		Name               string    `json:"name"`
		EarnedAchievements int       `json:"earnedAchievements"`
		CurrentGamerscore  int       `json:"currentGamerscore"`
		MaxGamerscore      int       `json:"maxGamerscore"`
	} `json:"titles"`
	PagingInfo struct {
		ContinuationToken *string `json:"continuationToken"`
		TotalRecords      int     `json:"totalRecords"`
		xuid              int
	} `json:"pagingInfo"`
	c *Client
}

func (x *XboxOneGames) More() (*XboxOneGames, error) {
	if x.PagingInfo.ContinuationToken == nil {
		return nil, nil
	}
	rsp, err := x.c.Get(
		fmt.Sprintf(
			"https://xboxapi.com/v2/%d/xboxonegames?continuationToken=%s",
			x.PagingInfo.xuid,
			x.PagingInfo.ContinuationToken,
		),
	)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	var data *XboxOneGames
	err = json.NewDecoder(rsp.Body).Decode(&data)
	if err != nil {
		data.PagingInfo.xuid = x.PagingInfo.xuid
		data.c = x.c
	}
	return data, err
}

func (c *Client) XboxOneGames(xuid int) (*XboxOneGames, error) {
	rsp, err := c.Get(fmt.Sprintf("https://xboxapi.com/v2/%d/xboxonegames", xuid))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	var data *XboxOneGames
	err = json.NewDecoder(rsp.Body).Decode(&data)
	if err != nil {
		data.PagingInfo.xuid = xuid
		data.c = c
	}
	return data, err
}

type Xbox360Games struct {
	Titles []struct {
		LastPlayed          time.Time `json:"lastPlayed"`
		CurrentAchievements int       `json:"currentAchievements"`
		CurrentGamerscore   int       `json:"currentGamerscore"`
		Sequence            int       `json:"sequence"`
		TitleID             int       `json:"titleId"`
		TitleType           int       `json:"titleType"`
		Platforms           []int     `json:"platforms"`
		Name                string    `json:"name"`
		TotalAchievements   int       `json:"totalAchievements"`
		TotalGamerscore     int       `json:"totalGamerscore"`
	} `json:"titles"`
	PagingInfo struct {
		ContinuationToken *string `json:"continuationToken"`
		TotalRecords      int     `json:"totalRecords"`
		xuid              int
	} `json:"pagingInfo"`
	Version time.Time `json:"version"`
	c       *Client
}

func (x *Xbox360Games) More() (*Xbox360Games, error) {
	if x.PagingInfo.ContinuationToken == nil {
		return nil, nil
	}
	rsp, err := x.c.Get(
		fmt.Sprintf(
			"https://xboxapi.com/v2/%d/xbox360games?continuationToken=%s",
			x.PagingInfo.xuid,
			x.PagingInfo.ContinuationToken,
		),
	)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	var data *Xbox360Games
	err = json.NewDecoder(rsp.Body).Decode(&data)
	if err != nil {
		data.PagingInfo.xuid = x.PagingInfo.xuid
		data.c = x.c
	}
	return data, err
}

func (c *Client) Xbox360Games(xuid int) (*Xbox360Games, error) {
	rsp, err := c.Get(fmt.Sprintf("https://xboxapi.com/v2/%d/xbox360games", xuid))
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	var data *Xbox360Games
	err = json.NewDecoder(rsp.Body).Decode(&data)
	if err != nil {
		data.PagingInfo.xuid = xuid
		data.c = c
	}
	return data, err
}
