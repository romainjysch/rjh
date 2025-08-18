package s3

import (
	"fmt"

	"rjh/config"
	"rjh/internal/s3"

	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	getCmd := &cobra.Command{
		Use:     "get <path/to/file> <file>",
		Short:   "Get object to local",
		Example: "  rjh s3 get backups/file file",
		Args:    cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			cfg, err := config.Load(config.PATH)
			if err != nil {
				return err
			}

			argv := []string{
				"get",
				"--enc-c", fmt.Sprintf("%s=%s", s3.RJH, cfg.S3.Key),
				fmt.Sprintf("%s%s", s3.RJH, args[0]),
				args[1],
			}

			s3.Run("mc", argv...)

			return nil
		},
	}

	return getCmd
}
