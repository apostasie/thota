package wego

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

	"go.farcloser.world/core/log"
	"go.farcloser.world/core/uuid"

	"github.com/apostasie/thota/pkg/away"
)

type shareData struct {
	Version    string  `json:"version"`
	ProviderID string  `json:"providerId"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

func shareLink(placeID string, latitude, longitude float64) string {
	sd := &shareData{
		Version:    "1.0.3",
		Longitude:  longitude,
		Latitude:   latitude,
		ProviderID: placeID,
	}
	jsonEnc, _ := json.Marshal(sd)
	b64Enc := base64.URLEncoding.EncodeToString(jsonEnc)

	return "https://share.here.com/p/e-" + b64Enc
}

// AwayToWeGo convert a take-away set of lists into WeGo collections
func AwayToWeGo(apiKey string, awayLists []*away.List) ([]*Collection, error) {
	client, err := NewClient(apiKey)
	if err != nil {
		return nil, err
	}

	collections := make([]*Collection, 0, len(awayLists))

	for _, awayList := range awayLists {
		log.Info().Str("list", awayList.Name).Msg("Processing away list")
		collections = append(collections, awayListToWeGoCollection(client, awayList))
	}

	return collections, nil
}

//revive:disable:function-length
func awayListToWeGoCollection(client *Client, awayList *away.List) *Collection {
	collection := &Collection{
		Type:   "FeatureCollection",
		Places: []*Place{},
	}

	collectionID := uuid.New()

	for _, place := range awayList.Places {
		candidates, err := client.Geocode(
			context.Background(),
			place.Name,
			place.Latitude,
			place.Longitude,
		)
		if err != nil {
			log.Error().Err(err).Msg("failed to query Wego API for place")
			continue
		}

		if len(candidates) == 0 {
			log.Error().
				Str("name", place.Name).
				Str("address", place.Address).
				Float64("latitude", place.Latitude).
				Float64("longitude", place.Longitude).
				Msg("No candidates found for place. Trying again with the address instead.")

			candidates, err = client.Geocode(
				context.Background(),
				place.Address,
				place.Latitude,
				place.Longitude,
			)
			if err != nil {
				log.Error().Err(err).Msg("failed to query Wego API for place")
				continue
			}

			if len(candidates) == 0 {
				log.Error().Msg("Giving up!!!!!!")

				continue
			}
		}

		candidate := candidates[0]

		geoCoordinates := &GeoCoordinates{
			Latitude:  candidate.Position.Lat,  // candidate.Position.Lat,
			Longitude: candidate.Position.Long, // candidate.Position.Long,
		}

		location := &Location{
			Address:        candidate.Address.Label,
			GeoCoordinates: geoCoordinates,
		}

		ind := strings.LastIndex(candidate.ID, ":")
		shortID := candidate.ID[ind+1:]

		properties := &Properties{
			CollectionMame: awayList.Name,
			// Collection description will be lost
			CollectionID: collectionID,
			WEGOURL:      shareLink(shortID, candidate.Position.Lat, candidate.Position.Long),
			Location:     location,
		}

		geometry := &Geometry{
			Type:        "Point",
			Coordinates: []float64{candidate.Position.Long, candidate.Position.Lat},
		}

		name := candidate.Title
		if name == candidate.Address.Label {
			name = place.Name
		}

		p := &Place{
			Type:       "Feature",
			ID:         candidate.ID,
			Title:      name,
			Properties: properties,
			Geom:       geometry,
			BBox: []float64{
				candidate.Position.Long,
				candidate.Position.Lat,
				candidate.Position.Long,
				candidate.Position.Lat,
			},
		}

		collection.Places = append(collection.Places, p)
	}

	return collection
}
