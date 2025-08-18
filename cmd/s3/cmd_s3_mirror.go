package s3

import (
	"fmt"

	"rjh/config"
	"rjh/internal/s3"

	"github.com/spf13/cobra"
)

func newMirrorCmd() *cobra.Command {
	mirrorCmd := &cobra.Command{
		Use:     "mirror",
		Short:   "Synchronize object(s) to bucket",
		Example: "  rjh s3 mirror backups/ backups/",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(config.PATH)
			if err != nil {
				return err
			}

			argv := []string{
				"mirror",
				"--remove",
				"--enc-c", fmt.Sprintf("%s=%s", s3.RJH, cfg.S3.Key),
				args[0],
				fmt.Sprintf("%s%s", s3.RJH, args[1]),
			}

			s3.Run("mc", argv...)

			return nil
		},
	}

	return mirrorCmd
}
