package tasks

import "github.com/spf13/cobra"

var TasksCmd = &cobra.Command{
	Use:     "tasks",
	Short:   "Manage tasks",
	Aliases: []string{"t"},
}

func init() {
	TasksCmd.AddCommand(newListCmd())
}
