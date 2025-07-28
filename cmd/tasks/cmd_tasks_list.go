package tasks

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"rjh/internal/tasks"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List tasks to do",
		Example: "  rjh tasks list",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			tasks, err := todo.FetchTasks("internal/tasks/data/tasks.csv")
			if err != nil {
				return err
			}

			all, err := cmd.Flags().GetBool("all")
			if err != nil {
				return fmt.Errorf("getting flag \"all\": %w", err)
			}

			if all {
				printAllTasks(tasks)
			} else {
				printTasks(tasks)
			}

			return nil
		},
	}
	listCmd.Flags().BoolP("all", "a", false, "list all tasks")

	return listCmd
}

func printTasks(tasks []*todo.Task) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	defer tw.Flush()

	fmt.Fprintf(tw, "%s\t%s\t%s\n", "Id", "Task", "Created")

	for _, task := range tasks {
		if task.Completed != 0 {
			continue
		}

		createdSince := getTimeDiff(task.Created)

		fmt.Fprintf(tw, "%d\t%s\t%s\n", task.Id, task.Description, createdSince)
	}
}

func printAllTasks(tasks []*todo.Task) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	defer tw.Flush()

	fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", "Id", "Task", "Created", "Completed")

	for _, task := range tasks {
		createdSince := getTimeDiff(task.Created)

		var completedSince string
		if task.Completed != 0 {
			completedSince = getTimeDiff(task.Completed)
		} else {
			completedSince = ""
		}

		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\n", task.Id, task.Description, createdSince, completedSince)
	}
}

func getTimeDiff(creation int64) string {
	return timediff.TimeDiff(
		time.Unix(creation, 0),
		timediff.WithStartTime(time.Now()))
}
