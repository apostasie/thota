package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli"
	"go.farcloser.world/core/filesystem"
	"go.farcloser.world/core/log"

	"github.com/apostasie/thota/pkg/away"
	"github.com/apostasie/thota/pkg/wego"
)

var wegoCommand = cli.Command{
	Name:  "wego",
	Usage: "convert a takeaway file into a wego collection",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "source", Usage: ""},
	},
	Action: wegoAction,
}

func wegoAction(c *cli.Context) {
	source := c.String("source")

	apiKey := os.Getenv("WEGO_API_KEY")

	f, err := os.ReadFile(source)
	if err != nil {
		log.Fatal().Err(err).Str("file", source).Msg("failed to read file")
	}

	lists := []*away.List{}
	_ = json.Unmarshal(f, &lists)
	result, err := wego.AwayToWeGo(apiKey, lists)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to convert to WeGo format")
	}

	for index, list := range result {
		wegoOutput, _ := json.MarshalIndent(list, "", "  ")
		_ = os.WriteFile(
			fmt.Sprintf("wego-%d.json", index+1),
			wegoOutput,
			filesystem.FilePermissionsDefault,
		)
	}
}
