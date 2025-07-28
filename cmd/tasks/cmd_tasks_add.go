package tasks

import (
	"fmt"
	"os"

	"rjh/internal/tasks"

	"github.com/spf13/cobra"
)

func newAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:     "add \"<description>\"",
		Short:   "Add a task",
		Example: "  rjh tasks add \"write a blog post\"",
		Aliases: []string{"a"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filename, ok := os.LookupEnv("TASKS_FILEPATH")
			if !ok {
				return fmt.Errorf("no tasks filepath variable found")
			}

			description := args[0]
			if description == "" {
				return fmt.Errorf("task description can't be empty")
			}

			if err := tasks.Add(description, filename); err != nil {
				return err
			}

			fmt.Printf("Task \"%s\" added.\n", description)

			return nil
		},
	}

	return addCmd
}
