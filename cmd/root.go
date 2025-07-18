package cmd

import (
	"rjh/cmd/network"
	"rjh/cmd/weather"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "rjh",
	Short:   "Personal command-line tool",
	Version: "0.1.3",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	rootCmd.AddCommand(
		network.NetworkCmd,
		weather.WeatherCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
