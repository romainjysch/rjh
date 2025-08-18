package s3

import (
	"fmt"

	"rjh/config"
	"rjh/internal/s3"

	"github.com/spf13/cobra"
)

func newStatCmd() *cobra.Command {
	statCmd := &cobra.Command{
		Use:     "stat <path/to/object>",
		Short:   "Show object metadata",
		Example: "  rjh s3 stat backups/file",
		Args:    cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			cfg, err := config.Load(config.PATH)
			if err != nil {
				return err
			}

			argv := []string{
				"stat",
				"--enc-c", fmt.Sprintf("%s=%s", s3.RJH, cfg.S3.Key),
				fmt.Sprintf("%s%s", s3.RJH, args[0]),
			}

			s3.Run("mc", argv...)

			return nil
		},
	}

	return statCmd
}
