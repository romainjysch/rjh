package tasks

import (
	"fmt"
	"rjh/internal/tasks"
	"strconv"

	"github.com/spf13/cobra"
)

func newCompleteCmd() *cobra.Command {
	var completeCmd = &cobra.Command{
		Use:     "complete <id>",
		Short:   "Complete a task",
		Example: "  rjh tasks complete 10",
		Aliases: []string{"c"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("task id must be an integer")
			}

			if err := tasks.Complete(int(id), tasks.FILENAME); err != nil {
				return err
			}

			fmt.Printf("Task %d completed.\n", id)

			return tasks.Complete(int(id), tasks.FILENAME)
		},
	}

	return completeCmd
}
