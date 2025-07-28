package cmd

import (
	"rjh/cmd/network"
	"rjh/cmd/tasks"
	"rjh/cmd/weather"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "rjh",
	Short:   "Personal command-line tool",
	Version: "0.2.1",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	rootCmd.AddCommand(
		network.NetworkCmd,
		tasks.TasksCmd,
		weather.WeatherCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
