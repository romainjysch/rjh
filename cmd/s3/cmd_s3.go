package s3

import "github.com/spf13/cobra"

var S3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Object storage management",
}

func init() {
	S3Cmd.AddCommand(
		newCpCmd(),
		newGetCmd(),
		newLsCmd(),
		newMirrorCmd(),
		newRmCmd(),
		newStatCmd(),
	)
}
