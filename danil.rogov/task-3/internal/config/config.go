package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input_file"`
	OutputFile string `yaml:"output_file"`
}

func LoadConfig(path string) (*Config, error) {
	if path == "" {
		return nil, &os.PathError{Op: "open", Path: "", Err: os.ErrNotExist}
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	if config.InputFile == "" {
		return nil, &os.PathError{Op: "open", Path: "input_file", Err: os.ErrInvalid}
	}
	if config.OutputFile == "" {
		return nil, &os.PathError{Op: "open", Path: "output_file", Err: os.ErrInvalid}
	}

	return &config, nil
}
