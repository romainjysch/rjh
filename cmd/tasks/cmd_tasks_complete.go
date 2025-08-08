package tasks

import (
	"fmt"
	"strconv"

	"rjh/config"
	"rjh/internal/tasks"

	"github.com/spf13/cobra"
)

func newCompleteCmd() *cobra.Command {
	var completeCmd = &cobra.Command{
		Use:     "complete <id>",
		Short:   "Complete a task",
		Example: "  rjh tasks complete 10",
		Aliases: []string{"c"},
		Args:    cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			cfg, err := config.Load(config.PATH)
			if err != nil {
				return err
			}

			t, file, err := tasks.Load(cfg.Tasks.Path)
			if err != nil {
				return err
			}
			defer file.Close()

			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("task id must be an integer")
			}

			if err := tasks.Complete(int(id), t, file); err != nil {
				return err
			}

			fmt.Printf("Task %d completed.\n", id)

			return nil
		},
	}

	return completeCmd
}
