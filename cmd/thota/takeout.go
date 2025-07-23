package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli"
	"go.farcloser.world/core/log"

	"github.com/apostasie/thota/pkg/google"
)

var takeoutCommand = cli.Command{
	Name:  "takeout",
	Usage: "process a Google Takeout directory into a takeaway file",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "source", Usage: ""},
	},
	Action: takeoutAction,
}

func takeoutAction(c *cli.Context) {
	source := c.String("source")

	apiKey := os.Getenv("GOOGLE_API_KEY")

	awaySet, err := google.TakeAway(apiKey, source)
	if err != nil {
		log.Fatal().Err(err).Str("path", source).Msg("Error reading takeout directory")
	}

	awayOutput, _ := json.MarshalIndent(awaySet, "", "  ")
	//nolint:forbidigo
	fmt.Println(string(awayOutput))
}
