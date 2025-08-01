package tasks

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type Task struct {
	Description string `csv:"description"`
	Created     int64  `csv:"created"`
	Completed   int64  `csv:"completed"`
	Deleted     int64  `csv:"deleted"`
}

func Load(filename string) ([]*Task, *os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, nil, fmt.Errorf("opening csv file: %w", err)
	}

	var tasks []*Task

	if err := gocsv.UnmarshalFile(file, &tasks); err != nil {
		return nil, nil, fmt.Errorf("unmarshalling csv: %w", err)
	}

	return tasks, file, nil
}

func Add(description string, tasks []*Task, file *os.File) error {
	if description == "" {
		return fmt.Errorf("task description cannot be empty")
	}

	task := Task{
		Description: description,
		Created:     time.Now().Unix(),
		Completed:   0,
	}

	tasks = append(tasks, &task)

	if err := seekAndTruncate(file); err != nil {
		return err
	}

	if err := gocsv.MarshalFile(&tasks, file); err != nil {
		return fmt.Errorf("writing to csv: %w", err)
	}

	return nil
}

func Complete(id int, tasks []*Task, file *os.File) error {
	if id < 0 || id > len(tasks) {
		return fmt.Errorf("invalid task id: %d", id)
	}

	tasks[id].Completed = time.Now().Unix()

	if err := seekAndTruncate(file); err != nil {
		return err
	}

	if err := gocsv.MarshalFile(&tasks, file); err != nil {
		return fmt.Errorf("writing to csv: %w", err)
	}

	return nil
}

func Delete(id int, tasks []*Task, file *os.File) error {
	if id < 0 || id > len(tasks) {
		return fmt.Errorf("invalid task id: %d", id)
	}

	tasks[id].Deleted = time.Now().Unix()

	if err := seekAndTruncate(file); err != nil {
		return err
	}

	if err := gocsv.MarshalFile(&tasks, file); err != nil {
		return fmt.Errorf("writing to csv: %w", err)
	}

	return nil
}

func seekAndTruncate(f *os.File) error {
	if _, err := f.Seek(0, 0); err != nil {
		return fmt.Errorf("seeking to beginning of file: %w", err)
	}

	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("truncating file: %w", err)
	}

	return nil
}
