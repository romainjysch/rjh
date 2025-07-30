package weather

import "github.com/spf13/cobra"

var WeatherCmd = &cobra.Command{
	Use:     "weather",
	Short:   "Weather information",
	Aliases: []string{"w"},
}

func init() {
	WeatherCmd.AddCommand(newCurrentCmd())
}
