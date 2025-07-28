package tasks

import (
	"fmt"
	"os"
	"strconv"

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
		RunE: func(cmd *cobra.Command, args []string) error {
			filename, ok := os.LookupEnv("TASKS_FILEPATH")
			if !ok {
				return fmt.Errorf("no tasks filepath variable found")
			}

			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("task id must be an integer")
			}

			if err := tasks.Complete(int(id), filename); err != nil {
				return err
			}

			fmt.Printf("Task %d completed.\n", id)

			return nil
		},
	}

	return completeCmd
}
