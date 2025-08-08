package tasks

import (
	"fmt"

	"rjh/config"
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
