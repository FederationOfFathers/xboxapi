package xboxapi

import (
	"encoding/json"
	"fmt"
)

type ActivityItems struct {
	AchievementScid        string `json:"achievementScid"`
	AchievementID          int    `json:"achievementId"`
	AchievementType        string `json:"achievementType"`
	AchievementIcon        string `json:"achievementIcon"`
	Gamerscore             int    `json:"gamerscore"`
	AchievementName        string `json:"achievementName"`
	AchievementDescription string `json:"achievementDescription"`
	IsSecret               bool   `json:"isSecret"`
	HasAppAward            bool   `json:"hasAppAward"`
	HasArtAward            bool   `json:"hasArtAward"`
	ContentImageURI        string `json:"contentImageUri"`
	ContentTitle           string `json:"contentTitle"`
	Platform               string `json:"platform"`
	TitleID                int    `json:"titleId"`
	UserImageURIMd         string `json:"userImageUriMd"`
	UserImageURIXs         string `json:"userImageUriXs"`
	Description            string `json:"description"`
	Date                   string `json:"date"`
	HasUgc                 bool   `json:"hasUgc"`
	ActivityItemType       string `json:"activityItemType"`
	ContentType            string `json:"contentType"`
	ShortDescription       string `json:"shortDescription"`
	ItemText               string `json:"itemText"`
	ItemImage              string `json:"itemImage"`
	ShareRoot              string `json:"shareRoot"`
	FeedItemID             string `json:"feedItemId"`
	ItemRoot               string `json:"itemRoot"`
	NumLikes               int    `json:"numLikes"`
	HasLiked               bool   `json:"hasLiked"`
	Gamertag               string `json:"gamertag"`
	RealName               string `json:"realName"`
	DisplayName            string `json:"displayName"`
	UserImageURI           string `json:"userImageUri"`
	UserXuid               int    `json:"userXuid"`
	Activity               struct {
		NumShares              int         `json:"numShares"`
		NumLikes               int         `json:"numLikes"`
		NumComments            int         `json:"numComments"`
		UgcCaption             interface{} `json:"ugcCaption"`
		AuthorType             string      `json:"authorType"`
		AchievementScid        string      `json:"achievementScid"`
		AchievementID          int         `json:"achievementId"`
		ActivityItemType       string      `json:"activityItemType"`
		AchievementType        string      `json:"achievementType"`
		UserXuid               int         `json:"userXuid"`
		AchievementIcon        string      `json:"achievementIcon"`
		Date                   string      `json:"date"`
		Gamerscore             int         `json:"gamerscore"`
		ContentType            string      `json:"contentType"`
		AchievementName        string      `json:"achievementName"`
		TitleID                int         `json:"titleId"`
		AchievementDescription string      `json:"achievementDescription"`
		Platform               string      `json:"platform"`
		IsSecret               bool        `json:"isSecret"`
		Sandboxid              string      `json:"sandboxid"`
		SharedSourceUser       int         `json:"sharedSourceUser"`
		UserKey                interface{} `json:"userKey"`
		RarityCategory         string      `json:"rarityCategory"`
		ScID                   string      `json:"scid"`
		RarityPercentage       float32     `json:"rarityPercentage"`
	} `json:"activity"`
	AuthorInfo struct {
		Name       string `json:"name"`
		SecondName string `json:"secondName"`
		ImageURL   string `json:"imageUrl"`
		AuthorType string `json:"authorType"`
		ID         int    `json:"id"`
	} `json:"authorInfo"`
}

type Activity struct {
	ActivityItems          []ActivityItems `json:"activityItems"`
	NumItems               int             `json:"numItems"`
	PollingToken           int             `json:"pollingToken"`
	PollingIntervalSeconds int             `json:"pollingIntervalSeconds"`
	ContinuationToken      *json.Number    `json:"contToken"`
}

func (c *Client) Activity(xuid int, continuationToken *json.Number) (*Activity, error) {
	var url string
	if continuationToken == nil {
		url = fmt.Sprintf("https://xboxapi.com/v2/%d/activity", xuid)
	} else {
		url = fmt.Sprintf("https://xboxapi.com/v2/%d/activity?continuationToken=%s", xuid, continuationToken.String())
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
	var data *Activity
	err = json.NewDecoder(rsp.Body).Decode(&data)
	return data, err
}
