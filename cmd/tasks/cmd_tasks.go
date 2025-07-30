package tasks

import "github.com/spf13/cobra"

var TasksCmd = &cobra.Command{
	Use:     "tasks",
	Short:   "Task management",
	Aliases: []string{"t"},
}

func init() {
	TasksCmd.AddCommand(
		newAddCmd(),
		newCompleteCmd(),
		newDeleteCmd(),
		newListCmd())
}
