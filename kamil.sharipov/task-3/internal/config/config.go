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

func (c *Config) Validate() error {
	if c.InputFile == "" {
		return errNoInputFile
	}

	if c.OutputFile == "" {
		return errNoOutputFile
	}

	return nil
}

func LoadConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open config file %q: %w", path, err)
	}

	defer func() {
		if ferr := file.Close(); ferr != nil {
			panic(fmt.Errorf("close file %q: %w", path, err))
		}
	}()

	var config Config

	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to decode config: %w", err)
	}

	err = config.Validate()
	if err != nil {
		return Config{}, fmt.Errorf("invalid config: %w", err)
	}

	return config, nil
}
