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

			t, file, err := tasks.Load(filename)
			if err != nil {
				return err
			}
			defer file.Close()

			description := args[0]
			if description == "" {
				return fmt.Errorf("task description can't be empty")
			}

			if err := tasks.Add(description, t, file); err != nil {
				return err
			}

			fmt.Printf("Task \"%s\" added.\n", description)

			return nil
		},
	}

	return addCmd
}
