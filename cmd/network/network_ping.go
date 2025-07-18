package network

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/prometheus-community/pro-bing"
	"github.com/spf13/cobra"
)

func generatePingCmd() *cobra.Command {
	var pingCmd = &cobra.Command{
		Use:     "ping <IP address or domain>",
		Short:   "Ping an IP address or domain",
		Example: "  rjh network ping 1.1.1.1",
		Aliases: []string{"p"},
		Args:    cobra.ExactArgs(1),
		RunE:    runPing,
	}
	pingCmd.Flags().IntP("count", "c", 5, "Stop after sending (and receiving) count")

	return pingCmd
}

func runPing(cmd *cobra.Command, args []string) error {
	fmt.Println("Running ping command...\n")

	address := args[0]

	pinger, err := probing.NewPinger(address)
	if err != nil {
		return fmt.Errorf("could not create pinger: %w", err)
	}

	pinger.Count, err = cmd.Flags().GetInt("count")
	if err != nil {
		return fmt.Errorf("could not get count flag: %w", err)
	}

	err = pinger.Run()
	if err != nil {
		return fmt.Errorf("could not ping %s: %w", address, err)
	}
	s := pinger.Statistics()

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer tw.Flush()

	fmt.Fprintf(tw, "IP address:\t%s\n", s.IPAddr.IP.String())
	fmt.Fprintf(tw, "Send:\t%d\n", s.PacketsSent)
	fmt.Fprintf(tw, "Received:\t%d\n", s.PacketsRecv)
	fmt.Fprintf(tw, "Loss:\t%.2f%%\n", s.PacketLoss)
	fmt.Fprintf(tw, "Max rtt:\t%d ms\n", s.MaxRtt.Milliseconds())
	fmt.Fprintf(tw, "Min rtt:\t%d ms\n", s.MinRtt.Milliseconds())
	fmt.Fprintf(tw, "Avg rtt:\t%d ms\n", s.AvgRtt.Milliseconds())
	fmt.Fprintf(tw, "SD rtt:\t%d ms\n", s.StdDevRtt.Milliseconds())

	return nil
}
