package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TESTDIR  = "testdata"
	TESTFILE = "testconfig.yml"
)

func TestLoad(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expected := Config{
			OpenWeatherMap: OpenWeatherMap{
				Key: "fake_api_key",
			},
			Tasks: Tasks{
				Path: "/fake/path/to/tasks.csv",
			},
		}

		config, err := Load(filepath.Join(TESTDIR, TESTFILE))
		require.NoError(t, err)
		require.NotNil(t, config)
		require.Equal(t, expected.OpenWeatherMap.Key, config.OpenWeatherMap.Key)
		require.Equal(t, expected.Tasks.Path, config.Tasks.Path)
	})

	t.Run("File Not Found", func(t *testing.T) {
		_, err := Load("")
		require.Contains(t, err.Error(), "opening config file")
		require.Contains(t, err.Error(), "no such file")
	})

	t.Run("Bad YAML File", func(t *testing.T) {
		f, err := os.CreateTemp(t.TempDir(), "tmpbad.yml")
		require.NoError(t, err)
		defer f.Close()

		err = os.WriteFile(f.Name(), []byte("test will fail"), os.ModePerm)
		require.NoError(t, err)

		_, err = Load(f.Name())
		require.Contains(t, err.Error(), "decoding config file")
	})
}
