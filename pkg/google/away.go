package google

import (
	"go.farcloser.world/core/log"

	"github.com/apostasie/thota/pkg/away"
)

// TakeAway reads a Google Takeout archive from the specified directory path and returns a slice of away.Lists.
func TakeAway(apiKey, directoryPath string) ([]*away.List, error) {
	log.Info().Str("directory", directoryPath).Msg("Gathering data from Google takeout archive")

	listSet, err := readTakeOutDirectory(directoryPath)
	if err != nil {
		return nil, err
	}

	log.Info().Int("number of lists", len(listSet)).Msg("Done reading takeout files")

	for _, list := range listSet {
		log.Info().Str("name", list.Name).Int("places", len(list.Records)).Send()
	}

	log.Info().Msg("Analyzing records")

	cli := &legacyPlacesClient{
		APIKey: apiKey,
	}

	lists := make([]*away.List, 0, len(listSet))

	for _, list := range listSet {
		awayList := &away.List{
			Name:        list.Name,
			Description: list.Description,
			Places:      []*away.Place{},
		}

		lists = append(lists, awayList)

		for _, record := range list.Records {
			var details *legacyPlace

			details, err = cli.GetPlaceDetails(record.LegacyID)
			if err != nil {
				log.Error().
					Err(err).
					Str("list", list.Name).
					Str("title", record.Title).
					Str("id", record.LegacyID).
					Msg("failed to get place details")
				// Maybe we *could* prompt the user and offer suggestions using search on the place name.
				continue
			}

			place := &away.Place{
				Name:      details.Result.Name,
				Address:   details.Result.FormattedAddress,
				Latitude:  details.Result.Geometry.Location.Lat,
				Longitude: details.Result.Geometry.Location.Lng,
				Note:      record.Note,
			}
			awayList.Places = append(awayList.Places, place)
		}
	}

	return lists, nil
}
