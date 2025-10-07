package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	errNoInputFile  error = errors.New("no input file")
	errNoOutputFile error = errors.New("no output file")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %q: %w", path, err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	if config.InputFile == "" {
		return nil, errNoInputFile
	}
	if config.OutputFile == "" {
		return nil, errNoOutputFile
	}

	return &config, nil
}
