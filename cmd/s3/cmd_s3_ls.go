package s3

import (
	"fmt"

	"rjh/internal/s3"

	"github.com/spf13/cobra"
)

func newLsCmd() *cobra.Command {
	lsCmd := &cobra.Command{
		Use:     "ls",
		Short:   "List objects in bucket",
		Example: "  rjh s3 ls backups/",
		Args:    cobra.MaximumNArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			var folder string
			if len(args) == 1 {
				folder = fmt.Sprintf("%s%s", s3.RJH, args[0])
			} else {
				folder = s3.RJH
			}

			argv := []string{
				"ls",
				folder,
			}

			s3.Run("mc", argv...)
		},
	}

	return lsCmd
}
