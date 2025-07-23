package wego

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"go.einride.tech/here/geocodingsearchv7"
	"go.farcloser.world/core/log"
)

var (
	// ErrMissingAPIKey is returned if the provided key is empty.
	ErrMissingAPIKey = errors.New("missing API key")
	// ErrFailedQuery is returned when a WeGo API call fails.
	ErrFailedQuery = errors.New("failed querying Wego API")
)

// Client is a WeGo API client.
type Client struct {
	apiKey string
	inner  *geocodingsearchv7.Client
}

// NewClient returns a new WeGo API client.
func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, ErrMissingAPIKey
	}
	cli := &Client{
		apiKey: apiKey,
	}

	cli.inner = geocodingsearchv7.NewClient(
		geocodingsearchv7.NewAPIKeyHTTPClient(cli.apiKey, http.DefaultClient.Transport),
	)

	return cli, nil
}

// ReverseGeocode uses WeGo API to retrieve a place name and additional information based on coordinates.
// Note: it does not work well in many places, especially outside the US and Europe.
func (client *Client) ReverseGeocode(
	ctx context.Context,
	latitude, longitude float64,
) (*geocodingsearchv7.ReverseGeocodingResponse, error) {
	log.Info().Float64("lat", latitude).Float64("long", longitude).Msg("Reverse geocoding")

	response, err := client.inner.ReverseGeocoding.ReverseGeocoding(
		ctx,
		&geocodingsearchv7.ReverseGeocodingRequest{
			GeoPosition: &geocodingsearchv7.GeoWaypoint{
				Lat:  latitude,
				Long: longitude,
			},
		},
	)
	if err != nil {
		return nil, errors.Join(ErrFailedQuery, err)
	}

	return response, nil
}

// Geocode uses WeGo API to retrieve a place coordinates, id, and additional information.
// "https://geocode.search.hereapi.com/v1/geocode?at=37.78,-122.43&q=Music+Lovers+Audio+Video+2295+Bush+St,+San+Francisco,+CA+94115,+USA&apiKey="
func (client *Client) Geocode(
	ctx context.Context,
	name string,
	latitude, longitude float64,
) ([]geocodingsearchv7.GeocodingItem, error) {
	// Names from Google may have spurious information that wego is not happy with, so, strip out some of it.

	// eg: Casa da Praia Tapas Bar & Wine Bar (After 20.00h is dinner time.)
	ind := strings.Index(name, "(")
	if ind != -1 {
		name = name[:ind]
	}

	// eg: La Veranda Resort Phu Quoc - MGallery Collection
	ind = strings.Index(name, " - ")
	if ind != -1 {
		name = name[:ind]
	}

	// eg: Blackheath Lodge by The Oyster Collection
	ind = strings.Index(name, " by ")
	if ind != -1 {
		name = name[:ind]
	}

	/*
		Currently not working with the API, but does work on the map :s
		D.T. Suzuki Museum
		https://geocode.search.hereapi.com/v1/geocode?q=
		D.T. Suzuki Museum
		Museu da Farmacia
		Toriyaki Ohana
		Mori Mori Sushi Kanazawa
		11-1 Murasakino Higashifujinomorichō, Kita Ward, Kyoto, 603-8223, Japan
		Le Cotte Rôti
		GENDY 南青山店

		and many others. WeGo geocode is just awful outside the US and Europe (and then Finland is very bad too).
	*/

	query := name

	log.Info().Str("query", query).Msg("Querying")

	// The query to geocode
	response, err := client.inner.Geocoding.Geocoding(ctx, &geocodingsearchv7.GeocodingRequest{
		Q: &query,
		GeoPosition: &geocodingsearchv7.GeoWaypoint{
			Lat:  latitude,
			Long: longitude,
		},
	})
	if err != nil {
		return nil, errors.Join(ErrFailedQuery, err)
	}

	return response.Items, nil
}
