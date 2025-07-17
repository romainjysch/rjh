package cmd

import (
	"github.com/spf13/cobra"
)

var rjhCmd = &cobra.Command{
	Use:     "rjh",
	Short:   "Personal command-line tool",
	Version: "0.1.3",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	rjhCmd.AddCommand(weatherCmd)
}

func Execute() error {
	return rjhCmd.Execute()
}
