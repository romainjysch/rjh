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
			tasks, file, err := tasks.FetchTasks(tasks.FILENAME)
			if err != nil {
				return err
			}
			defer file.Close()

			all, err := cmd.Flags().GetBool("all")
			if err != nil {
				return fmt.Errorf("parsing \"all\" flag: %w", err)
			}

			completed, err := cmd.Flags().GetBool("completed")
			if err != nil {
				return fmt.Errorf("parsing \"completed\" flag: %w", err)
			}

			if all {
				printAllTasks(tasks)
			} else if completed {
				printCompletedTasks(tasks)
			} else {
				printTasks(tasks)
			}

			return nil
		},
	}
	listCmd.Flags().BoolP("all", "a", false, "list all tasks")
	listCmd.Flags().BoolP("completed", "c", false, "list completed tasks")

	return listCmd
}

func printTasks(tasks []*tasks.Task) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	defer tw.Flush()

	fmt.Fprintf(tw, "%s\t%s\t%s\n", "Id", "Task", "Created")

	for i, task := range tasks {
		if task.Completed != 0 {
			continue
		}

		createdSince := getTimeDiff(task.Created)

		fmt.Fprintf(tw, "%d\t%s\t%s\n", i, task.Description, createdSince)
	}
}

func printCompletedTasks(tasks []*tasks.Task) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	defer tw.Flush()

	fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", "Id", "Task", "Created", "Completed")

	for i, task := range tasks {
		if task.Completed == 0 {
			continue
		}

		createdSince := getTimeDiff(task.Created)
		completedSince := getTimeDiff(task.Completed)

		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\n", i, task.Description, createdSince, completedSince)
	}
}

func printAllTasks(tasks []*tasks.Task) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	defer tw.Flush()

	fmt.Fprintf(tw, "%s\t%s\t%s\t%s\n", "Id", "Task", "Created", "Completed")

	for i, task := range tasks {
		createdSince := getTimeDiff(task.Created)

		var completedSince string
		if task.Completed != 0 {
			completedSince = getTimeDiff(task.Completed)
		} else {
			completedSince = ""
		}

		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\n", i, task.Description, createdSince, completedSince)
	}
}

func getTimeDiff(creation int64) string {
	return timediff.TimeDiff(
		time.Unix(creation, 0),
		timediff.WithStartTime(time.Now()))
}
