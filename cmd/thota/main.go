package main

import (
	"os"

	"github.com/urfave/cli"
	"go.farcloser.world/core/log"
)

func main() {
	envLogLevel := os.Getenv("LOG_LEVEL")

	logLevel := log.WarnLevel
	switch envLogLevel {
	case "debug":
		logLevel = log.DebugLevel
	case "info":
		logLevel = log.InfoLevel
	case "error":
		logLevel = log.ErrorLevel
	}

	log.Init(&log.Config{
		Level: logLevel,
	})

	app := cli.NewApp()
	app.Name = "thota"
	app.Usage = "a tool to process maps lists"
	app.Version = "0.1"
	app.Commands = []cli.Command{
		takeoutCommand,
		wegoCommand,
	}
	_ = app.Run(os.Args)
}
