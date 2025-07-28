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
		Short:   "List tasks",
		Example: "  rjh tasks list",
		Aliases: []string{"ls"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			filename, ok := os.LookupEnv("TASKS_FILEPATH")
			if !ok {
				return fmt.Errorf("no tasks filepath variable found")
			}

			t, file, err := tasks.Load(filename)
			if err != nil {
				return err
			}
			defer file.Close()

			all, err := cmd.Flags().GetBool("all")
			if err != nil {
				return fmt.Errorf("parsing \"all\" flag: %w", err)
			}

			if all {
				printAllTasks(t)
			} else {
				printTasks(t)
			}

			return nil
		},
	}
	listCmd.Flags().BoolP("all", "a", false, "list all tasks")

	return listCmd
}

func printTasks(tasks []*tasks.Task) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	defer tw.Flush()

	fmt.Fprintf(tw, "%s\t%s\t%s\n", "Id", "Task", "Created")

	for i, task := range tasks {
		if task.Completed != 0 || task.Deleted != 0 {
			continue
		}

		fmt.Fprintf(tw, "%d\t%s\t%s\n", i, task.Description, getTimeDiff(task.Created))
	}
}

func printAllTasks(tasks []*tasks.Task) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	defer tw.Flush()

	fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", "Id", "Task", "Created", "Completed")

	for i, task := range tasks {
		if task.Deleted != 0 {
			continue
		}

		var completedSince string
		if task.Completed != 0 {
			completedSince = getTimeDiff(task.Completed)
		} else {
			completedSince = ""
		}

		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\n", i, task.Description, getTimeDiff(task.Created), completedSince)
	}
}

func getTimeDiff(creation int64) string {
	return timediff.TimeDiff(
		time.Unix(creation, 0),
		timediff.WithStartTime(time.Now()))
}
