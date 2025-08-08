package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const PATH = "/Users/rjh/code/go/rjh/config/data/config.yml"

type OpenWeatherMap struct {
	Key string `yaml:"key"`
}

type Tasks struct {
	Path string `yaml:"filepath"`
}

type Config struct {
	OpenWeatherMap OpenWeatherMap `yaml:"openweathermap"`
	Tasks          Tasks          `yaml:"tasks"`
}

func Load(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening config file: %w", err)
	}
	defer f.Close()

	var config *Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("decoding config file: %w", err)
	}

	return config, nil
}
