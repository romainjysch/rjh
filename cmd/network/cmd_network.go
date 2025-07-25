package network

import (
	"github.com/spf13/cobra"
)

var NetworkCmd = &cobra.Command{
	Use:     "network",
	Short:   "Networks commands",
	Aliases: []string{"n"},
}

func init() {
	NetworkCmd.AddCommand(generatePingCmd())
}
