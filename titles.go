package xboxapi

import (
	"encoding/json"
	"fmt"
)

type TitleDetails struct {
	MediaGroup           interface{}
	MediaItemType        interface{}
	ID                   string
	Name                 string
	Description          string
	ReducedDescription   string
	ReducedName          string
	ReleaseDate          string
	TitleID              int `json:"TitleId"`
	VuiDisplayName       string
	DeveloperName        string
	PublisherName        string
	Updated              string
	ParentalRating       string
	ParentalRatingSystem string
	SortName             string
	KValue               int
	KValueNamespace      string
	HexTitleID           int `json:"HexTitleId"`
	AllTimePlayCount     int
	SevenDaysPlayCount   int
	ThirtyDaysPlayCount  int
	AllTimeRatingCount   int
	AllTimeAverageRating float64
	ResourceAccess       string
	IsRetail             bool
	ManualURL            string `json:"ManualUrl"`
	Genres               []struct {
		Name string
	}
	Images []struct {
		ID       string
		URL      string `json:"Url"`
		Purposes []string
		Purpose  string
		Height   int
		Width    int
	}
	Capabilities []struct {
		NonLocalizedName string
		Value            interface{}
	}
	LegacyIds []struct {
		IDType string `json:"IdType"`
		Value  interface{}
	}
	Availabilities []struct {
		AvailabilityID      string
		ContentID           string `json:"ContentId"`
		LicensePolicyTicket string
		OfferDisplayData    interface{}
		/* // Turns out this can be a string sometimes. See title 911120656
		OfferDisplayData struct {
			AcceptablePaymentInstrumentTypes []string `json:"acceptablePaymentInstrumentTypes"`
			AvailabilityDescription          string   `json:"availabilityDescription"`
			AvailabilityTitle                string   `json:"availabilityTitle"`
			CurrencyCode                     string   `json:"currencyCode"`
			DisplayListPrice                 string   `json:"displayListPrice"`
			DisplayPositionTag               int      `json:"displayPositionTag"`
			DisplayPrice                     string   `json:"displayPrice"`
			DistributionType                 string   `json:"distributionType"`
			IsPurchasable                    bool     `json:"isPurchasable"`
			ListPrice                        float64  `json:"listPrice"`
			OfferID                          string   `json:"offerId"`
			PrerequisiteProductID            string   `json:"prerequisiteProductId"`
			PrerequisiteProductType          string   `json:"prerequisiteProductType"`
			Price                            float64  `json:"price"`
			PromotionalText                  string   `json:"promotionalText"`
			ReducedTitle                     string   `json:"reducedTitle"`
			TaxTypeCode                      string   `json:"taxTypeCode"`
			SignedOffer                      string   `json:"SignedOffer"`
			SubscriptionBenefits             []struct {
				ID       string
				Benefits []string
			} `json:"SubscriptionBenefits"`
		}
		*/
		Devices []struct {
			Name string
		}
	}
	ParentalRatings []struct {
		RatingID         string `json:"RatingId"`
		LegacyRatingID   int    `json:"LegacyRatingId"`
		Rating           string
		RatingSystem     string
		RatingMinimumAge int
		LocalizedDetails struct {
			ShortName    string
			LongName     string
			RatingImages []struct {
				URL string `json:"Url"`
			}
		}
		RatingDescriptors []struct {
			NonLocalizedDescriptor string
			ID                     int `json:"Id"`
		}
		RatingDisclaimers []struct {
			Text string
		}
	}
}

type Title struct {
	Items          []TitleDetails `json:"Items"`
	ImpressionGUID string         `json:"ImpressionGuid"`
}

func (t *Title) Devices() []string {
	var rval = []string{}
	for _, item := range t.Items {
		for _, a := range item.Availabilities {
			for _, d := range a.Devices {
				var found bool
				for _, n := range rval {
					if n == d.Name {
						found = true
						break
					}
				}
				if !found {
					rval = append(rval, d.Name)
				}
			}
		}
	}
	return rval
}

func (c *Client) GameDetails(guid string) (*Title, error) {
	rsp, err := c.Get(fmt.Sprintf("https://xboxapi.com/v2/game-details/%s", guid))
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
	var rval *Title
	err = json.NewDecoder(rsp.Body).Decode(&rval)
	return rval, err
}

func (c *Client) GameDetailsHex(hex string) (*Title, error) {
	rsp, err := c.Get(fmt.Sprintf("https://xboxapi.com/v2/game-details-hex/%s", hex))
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
	var rval *Title
	err = json.NewDecoder(rsp.Body).Decode(&rval)
	return rval, err
}
