package network

import "github.com/spf13/cobra"

var NetworkCmd = &cobra.Command{
	Use:     "network",
	Short:   "Network statistics",
	Aliases: []string{"n"},
}

func init() {
	NetworkCmd.AddCommand(
		newPingCmd())
}
