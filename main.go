package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "rjh",
		Usage:   "Personal CLI tool",
		Version: "0.1.2",
		Commands: []*cli.Command{
			{
				Name:      "weather",
				Aliases:   []string{"w"},
				Usage:     "Print current weather for a specified city",
				ArgsUsage: "<city>",
				Action:    cmdWeather,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
