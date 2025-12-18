package cmd

import (
	"rjh/cmd/deepl"
	"rjh/cmd/s3"
	"rjh/cmd/tasks"
	"rjh/cmd/weather"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rjh",
	Short: "Personal command-line tool",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	rootCmd.AddCommand(
		deepl.DeeplCmd,
		s3.S3Cmd,
		tasks.TasksCmd,
		weather.WeatherCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
