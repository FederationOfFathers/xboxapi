package xboxapi

import (
	"encoding/json"
	"fmt"
)

type AchievementOne struct {
	ID                int         `json:"id"`
	ServiceConfigID   string      `json:"serviceConfigId"`
	Name              string      `json:"name"`
	ProgressState     string      `json:"progressState"`
	IsSecret          bool        `json:"isSecret"`
	Description       string      `json:"description"`
	LockedDescription string      `json:"lockedDescription"`
	ProductID         string      `json:"productId"`
	AchievementType   string      `json:"achievementType"`
	ParticipationType string      `json:"participationType"`
	EstimatedTime     string      `json:"estimatedTime"`
	IsRevoked         bool        `json:"isRevoked"`
	TimeWindow        interface{} `json:"timeWindow"` // no idea what goes here when it's not null
	Deeplink          interface{} `json:"deeplink"`   // no idea what goes here when it's not null
	TitleAssociations []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"titleAssociations"`
	Progression struct {
		Requirements []struct {
			ID                    string `json:"id"`
			Current               int    `json:"current"`
			Target                int    `json:"target"`
			OperationType         string `json:"operationType"`
			ValueType             string `json:"valueType"`
			RuleParticipationType string `json:"ruleParticipationType"`
		} `json:"requirements"`
		TimeUnlocked APITime `json:"timeUnlocked"`
	} `json:"progression"`
	MediaAssets []struct {
		Name string `json:"name"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"mediaAssets"`
	Platforms []string `json:"platforms"`
	Rewards   []struct {
		Name        interface{} `json:"name"`        // no idea what goes here when it's not null
		Description interface{} `json:"description"` // no idea what goes here when it's not null
		MediaAsset  interface{} `json:"mediaAsset"`  // no idea what goes here when it's not null
		Value       int         `json:"value"`
		Type        string      `json:"type"`
		ValueType   string      `json:"valueType"`
	} `json:"rewards"`
	Rarity struct {
		CurrentCategory   string  `json:"currentCategory"`
		CurrentPercentage float64 `json:"currentPercentage"`
	} `json:"achievements"`
}

type AchievementOneList []*AchievementOne

func (c *Client) AchievementsOne(xuid string, titleID int) (AchievementOneList, error) {
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
	var rval = AchievementOneList{}
	err = json.NewDecoder(rsp.Body).Decode(&rval)
	return rval, err
}
