package s3

import (
	"fmt"

	"rjh/config"
	"rjh/internal/s3"

	"github.com/spf13/cobra"
)

func newCpCmd() *cobra.Command {
	cpCmd := &cobra.Command{
		Use:     "cp <file> <path/to/file>",
		Short:   "Copy object to bucket",
		Example: "  rjh s3 cp file backups/file",
		Args:    cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			cfg, err := config.Load(config.PATH)
			if err != nil {
				return err
			}

			argv := []string{
				"cp",
				"--enc-c", fmt.Sprintf("%s=%s", s3.RJH, cfg.S3.Key),
				args[0],
				fmt.Sprintf("%s%s", s3.RJH, args[1]),
			}

			s3.Run("mc", argv...)

			return nil
		},
	}

	return cpCmd
}
