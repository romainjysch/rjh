package tasks

import (
	"fmt"
	"rjh/internal/tasks"

	"github.com/spf13/cobra"
)

func newAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:     "add -d <description>",
		Short:   "Add a task",
		Example: "  rjh tasks add -d \"write a blog post\"",
		RunE: func(cmd *cobra.Command, args []string) error {
			description, err := cmd.Flags().GetString("description")
			if err != nil {
				return fmt.Errorf("parsing \"description\" flag: %w", err)
			}
			if description == "" {
				return fmt.Errorf("task description can't be empty")
			}

			if err := tasks.Add(description, tasks.FILENAME); err != nil {
				return err
			}

			fmt.Printf("Task \"%s\" added.\n", description)

			return nil
		},
	}
	addCmd.Flags().StringP("description", "d", "", "task description")
	_ = addCmd.MarkFlagRequired("description")

	return addCmd
}
