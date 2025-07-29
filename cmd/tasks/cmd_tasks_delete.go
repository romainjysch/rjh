package tasks

import (
	"fmt"
	"os"
	"strconv"

	"rjh/internal/tasks"

	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:     "delete <id>",
		Short:   "Delete a task",
		Example: "  rjh tasks delete 10",
		Aliases: []string{"d"},
		Args:    cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			filename, ok := os.LookupEnv("TASKS_FILEPATH")
			if !ok {
				return fmt.Errorf("no tasks filepath variable found")
			}

			t, file, err := tasks.Load(filename)
			if err != nil {
				return err
			}
			defer file.Close()

			id, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("task id must be an integer")
			}

			if err := tasks.Delete(int(id), t, file); err != nil {
				return err
			}

			fmt.Printf("Task %d deleted.\n", id)

			return nil
		},
	}

	return deleteCmd
}
