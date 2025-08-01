package tasks

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	TESTDIR  = "testdata"
	TESTFILE = "testtemplate.csv"
)

func copyTestFile(t *testing.T) ([]*Task, *os.File) {
	tasks, fsrc, err := Load(filepath.Join(TESTDIR, TESTFILE))
	require.NoError(t, err)
	require.NotNil(t, fsrc)
	require.NotNil(t, tasks)
	defer fsrc.Close()

	fdst, err := os.CreateTemp(t.TempDir(), "tmptest.csv")
	require.NoError(t, err)

	_, err = fsrc.Seek(0, io.SeekStart)
	require.NoError(t, err)

	_, err = io.Copy(fdst, fsrc)
	require.NoError(t, err)

	_, err = fdst.Seek(0, io.SeekStart)
	require.NoError(t, err)

	return tasks, fdst
}

func TestLoad(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expected := []*Task{
			{Description: "task 1", Created: 131131, Completed: 131131, Deleted: 0},
			{Description: "task 2", Created: 131131, Completed: 0, Deleted: 0},
		}

		tasks, file, err := Load(filepath.Join(TESTDIR, TESTFILE))
		defer file.Close()

		require.NoError(t, err)
		require.NotNil(t, file)
		require.Equal(t, expected, tasks)
	})

	t.Run("File Not Found", func(t *testing.T) {
		_, _, err := Load("")
		require.Contains(t, err.Error(), "opening csv file")
		require.Contains(t, err.Error(), "no such file")
	})

	t.Run("Bad CSV File", func(t *testing.T) {
		ftmp, err := os.CreateTemp(t.TempDir(), "tmpbad.csv")
		require.NoError(t, err)
		defer ftmp.Close()

		var b strings.Builder
		b.WriteString("description,created,completed,deleted\n")
		b.WriteString("0,test,will,faill\n")
		content := b.String()

		err = os.WriteFile(ftmp.Name(), []byte(content), os.ModePerm)
		require.NoError(t, err)

		_, _, err = Load(ftmp.Name())
		require.Contains(t, err.Error(), "unmarshalling csv")
	})
}

func TestAdd(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expected := []*Task{
			{Description: "task 1", Created: 131131, Completed: 131131, Deleted: 0},
			{Description: "task 2", Created: 131131, Completed: 0, Deleted: 0},
			{Description: "task 3", Created: time.Now().Unix(), Completed: 0, Deleted: 0},
		}

		tasks, ftest := copyTestFile(t)
		defer ftest.Close()

		err := Add("task 3", tasks, ftest)
		require.NoError(t, err)

		_, err = ftest.Seek(0, io.SeekStart)
		require.NoError(t, err)

		updated, fcheck, err := Load(ftest.Name())
		defer fcheck.Close()
		require.NoError(t, err)
		require.Equal(t, len(expected), len(updated))

		// if flakiness
		require.InDelta(t, expected[2].Created, updated[2].Created, 1)

		// other fields
		require.Equal(t, expected[2].Description, updated[2].Description)
		require.Equal(t, expected[2].Completed, updated[2].Completed)
		require.Equal(t, expected[2].Deleted, updated[2].Deleted)

		// no change on i0 and i1
		require.Equal(t, expected[:2], updated[:2])
	})

	t.Run("Emtpy Description", func(t *testing.T) {
		tasks, ftest := copyTestFile(t)
		defer ftest.Close()

		err := Add("", tasks, ftest)
		require.Contains(t, err.Error(), "description cannot be empty")
	})
}

func TestComplete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expected := []*Task{
			{Description: "task 1", Created: 131131, Completed: 131131, Deleted: 0},
			{Description: "task 2", Created: 131131, Completed: time.Now().Unix(), Deleted: 0},
		}

		tasks, ftest := copyTestFile(t)
		defer ftest.Close()

		err := Complete(1, tasks, ftest)
		require.NoError(t, err)

		// if flakiness
		require.InDelta(t, expected[1].Completed, tasks[1].Completed, 1)

		// no change on other fields
		require.Equal(t, expected[1].Description, tasks[1].Description)
		require.Equal(t, expected[1].Created, tasks[1].Created)
		require.Equal(t, expected[1].Deleted, tasks[1].Deleted)

		// no change on i0
		require.Equal(t, expected[0], tasks[0])
	})

	t.Run("Invalid Task ID", func(t *testing.T) {
		tasks, ftest := copyTestFile(t)
		defer ftest.Close()

		err := Complete(-1, tasks, ftest)
		require.Contains(t, err.Error(), "invalid task id")

		err = Complete(len(tasks)+1, tasks, ftest)
		require.Contains(t, err.Error(), "invalid task id")
	})
}

func TestDelete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expected := []*Task{
			{Description: "task 1", Created: 131131, Completed: 131131, Deleted: time.Now().Unix()},
			{Description: "task 2", Created: 131131, Completed: 0, Deleted: 0},
		}

		tasks, ftest := copyTestFile(t)
		defer ftest.Close()

		err := Delete(0, tasks, ftest)
		require.NoError(t, err)

		// if flakiness
		require.InDelta(t, expected[0].Deleted, tasks[0].Deleted, 1)

		// no change on other fields
		require.Equal(t, expected[0].Description, tasks[0].Description)
		require.Equal(t, expected[0].Created, tasks[0].Created)
		require.Equal(t, expected[0].Completed, tasks[0].Completed)

		// no change on i1
		require.Equal(t, expected[1], tasks[1])
	})

	t.Run("Invalid Task ID", func(t *testing.T) {
		tasks, ftest := copyTestFile(t)
		defer ftest.Close()

		err := Delete(-1, tasks, ftest)
		require.Contains(t, err.Error(), "invalid task id")

		err = Delete(len(tasks)+1, tasks, ftest)
		require.Contains(t, err.Error(), "invalid task id")
	})
}
