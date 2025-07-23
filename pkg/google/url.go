package google

import (
	"errors"
	"regexp"
	"strings"
)

// ErrIDExtractionFailed is returned when the legacy place ID cannot be extracted from the URL.
var ErrIDExtractionFailed = errors.New("failed extracting legacy place ID from URL")

// urls in Google takeout files are in the form:
// https://www.google.com/maps/place/Music+Lovers+Audio+%26+Home+Theater+Systems/data=!4m2!3m1!1s0x808580c7fa33a725:0x19785b313f4653e9
// where 0x808580c7fa33a725:0x19785b313f4653e9 is the *legacy* place ID.
// This function just extracts that legacy place ID from the URL.
//
//revive:disable:add-constant
func extractLegacyPlaceID(mapURL string) (string, error) {
	pathDataRegex := regexp.MustCompile(`/data=([!][^?&#]+)`)
	matches := pathDataRegex.FindStringSubmatch(mapURL)
	if len(matches) < 2 {
		return "", ErrIDExtractionFailed
	}

	dataParam := matches[1]

	elements := strings.Split(dataParam, "!")
	for _, element := range elements {
		if element == "" {
			continue
		}

		re := regexp.MustCompile(`^(\d+)s(.+)$`)

		elementMatches := re.FindStringSubmatch(element)
		if len(elementMatches) == 3 && strings.Contains(elementMatches[2], "0x") {
			return elementMatches[2], nil
		}
	}

	return "", ErrIDExtractionFailed
}
