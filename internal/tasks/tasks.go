package todo

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Task struct {
	Id          int16  `csv:"id"`
	Description string `csv:"description"`
	Created     int64  `csv:"created"`
	Completed   int64  `csv:"completed"`
}

func FetchTasks(filename string) ([]*Task, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening csv file: %w", err)
	}
	defer f.Close()

	var tasks []*Task

	if err := gocsv.UnmarshalFile(f, &tasks); err != nil {
		return nil, fmt.Errorf("unmarshalling csv: %w", err)
	}

	return tasks, nil
}
