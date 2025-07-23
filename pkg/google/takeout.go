package google

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"go.farcloser.world/core/log"
)

var (
	// ErrTakeOutReadDirectoryFailed is returned when reading the takeout directory fails.
	ErrTakeOutReadDirectoryFailed = errors.New("error reading directory")
	// ErrTakeOutRecordParsingFailed is returned when parsing a takeout record fails.
	ErrTakeOutRecordParsingFailed = errors.New("failed parsing record")
)

//revive:disable:cognitive-complexity
func readTakeOutDirectory(directoryPath string) ([]*TakeOutList, error) {
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		return nil, errors.Join(ErrTakeOutReadDirectoryFailed, err)
	}

	var listSet []*TakeOutList

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".csv") {
			list := &TakeOutList{}
			fpath := filepath.Join(directoryPath, entry.Name())
			list.Name = strings.ReplaceAll(unespaceName(path.Base(fpath)), ".csv", "")

			var file *os.File

			file, err = os.Open(fpath)
			if err != nil {
				log.Error().Err(err).Str("file", fpath).Msg("unable to open file")

				continue
			}

			log.Info().Str("file", fpath).Msg("Reading takeout file")

			err = list.Read(file)
			_ = file.Close()

			if err != nil {
				log.Error().Err(err).Str("file", fpath).Msg("unable to read file")

				continue
			}

			if len(list.Records) == 0 {
				log.Warn().Str("file", fpath).Msg("No records found")

				continue
			}

			listSet = append(listSet, list)
		}
	}

	return listSet, nil
}

// TakeOutList represents an entire Google Takeout list (eg: a single Takeout CSV file).
type TakeOutList struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Records     []*TakeOutRecord `json:"records"`
}

// TakeOutRecord represents a single record in a Google Takeout CSV file.
type TakeOutRecord struct {
	// Title is the name of the place.
	Title string `json:"title"`
	// LegacyID is the legacy place ID, extracted from the URL.
	LegacyID string `json:"legacyId"`
	// Note is a user-provided note about the place.
	Note string `json:"note,omitempty"`
	// Tags is a comma-separated list of tags associated with the place.
	Tags string `json:"tags,omitempty"`
	// Comment is a user-provided comment about the place. XXX is this used?
	Comment string `json:"comment,omitempty"`
}

//revive:disable:cognitive-complexity,cyclomatic
func (list *TakeOutList) Read(reader io.Reader) error {
	var (
		err  error
		body []byte
	)

	list.Records = [](*TakeOutRecord)(nil)

	// First few lines MAY be the list description, so, read and discard them
	list.Description, body, err = preParse(reader)
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(bytes.NewReader(body))

	for {
		var (
			line []string
			pID  string
		)

		line, err = csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Error().Err(ErrTakeOutRecordParsingFailed).Err(err).Send()

			continue
		}

		pID, err = extractLegacyPlaceID(line[2])
		if err != nil {
			log.Error().Err(err).Str("line", strings.Join(line, " ")).
				Msg("Unable to extract place legacy ID")

			continue
		}

		//revive:disable:add-constant
		list.Records = append(list.Records, &TakeOutRecord{
			Title:    line[0],
			Note:     line[1],
			LegacyID: pID,
			Tags:     line[3],
			Comment:  line[4],
		})
	}

	return nil
}

// Google takeout CSV files may start with the description of the list.
// Ridiculous, but it is what it is.
func preParse(reader io.Reader) (string, []byte, error) {
	r := bufio.NewReader(reader)
	body := []byte{}
	headerSeen := false
	description := ""

	for {
		line, err := readLine(r)
		if err != nil {
			if err == io.EOF {
				break
			}

			return "", nil, errors.Join(ErrTakeOutRecordParsingFailed, err)
		}

		strLine := string(line)
		if !headerSeen {
			if strLine != "Title,Note,URL,Tags,Comment" {
				if strLine != "" {
					description += strLine
				}

				continue
			}

			headerSeen = true

			continue
		}

		if strLine != ",,,," {
			body = append(body, append(line, '\n')...)
		}
	}

	return description, body, nil
}

func readLine(reader *bufio.Reader) ([]byte, error) {
	var (
		isPrefix   = true
		err        error
		buff, line []byte
	)

	for isPrefix && err == nil {
		buff, isPrefix, err = reader.ReadLine()
		line = append(line, buff...)
	}

	//nolint:wrapcheck
	return line, err
}

func unespaceName(name string) string {
	return strings.ReplaceAll(name, "_", ":")
}
