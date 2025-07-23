package google

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

var (
	// ErrLegacyRequestFailed is returned when the API request fails, or the returned data cannot be parsed.
	ErrLegacyRequestFailed = errors.New("failed to make API request")
	// ErrLegacyAPIError is returned when the API returns an error in the Status property of the payload JSON.
	ErrLegacyAPIError = errors.New("API returned error status")
)

//revive:disable:nested-structs
type legacyPlace struct {
	Result struct {
		PlaceID          string `json:"place_id"`
		Name             string `json:"name"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
		Types       []string `json:"types,omitempty"`
		Website     string   `json:"website,omitempty"`
		PhoneNumber string   `json:"formatted_phone_number,omitempty"`
		Rating      float64  `json:"rating,omitempty"`
		Photos      []struct {
			PhotoReference string `json:"photo_reference"`
		} `json:"photos,omitempty"`
	} `json:"result"`
	Status string `json:"status"`
}

const baseLegacyURL = "https://maps.googleapis.com/maps/api/place/details/json"

type legacyPlacesClient struct {
	APIKey string
}

func (c *legacyPlacesClient) GetPlaceDetails(placeID string) (*legacyPlace, error) {
	params := url.Values{}
	params.Set("ftid", placeID)
	params.Set("key", c.APIKey)
	params.Set("fields", "place_id,name,formatted_address,geometry,types")
	// ,website,formatted_phone_number,rating,photos

	requestURL := baseLegacyURL + "?" + params.Encode()

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, errors.Join(ErrLegacyRequestFailed, err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Join(ErrLegacyRequestFailed, err)
	}

	var placeDetails legacyPlace
	if err := json.Unmarshal(body, &placeDetails); err != nil {
		return nil, errors.Join(ErrLegacyRequestFailed, err)
	}

	if placeDetails.Status != "OK" {
		return nil, errors.Join(ErrLegacyAPIError, errors.New("Google error: "+placeDetails.Status))
	}

	return &placeDetails, nil
}
