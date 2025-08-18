package s3

import (
	"fmt"

	"rjh/internal/s3"

	"github.com/spf13/cobra"
)

func newRmCmd() *cobra.Command {
	rmCmd := &cobra.Command{
		Use:     "rm <path/to/file>",
		Short:   "Remove an object",
		Example: "  rjh s3 rm backups/file",
		Args:    cobra.ExactArgs(1),
		Run: func(_ *cobra.Command, args []string) {
			argv := []string{
				"rm",
				fmt.Sprintf("%s%s", s3.RJH, args[0]),
			}

			s3.Run("mc", argv...)
		},
	}

	return rmCmd
}
