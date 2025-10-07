package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var errInvalidArgs error = errors.New("invalid arguments")

type Config struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

func Load(filePath string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(filePath)
	if err != nil {
		return cfg, fmt.Errorf("read file %q: %w", filePath, err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("unmarshal from %q: %w", filePath, err)
	}

	if cfg.Input == "" || cfg.Output == "" {
		return cfg, fmt.Errorf("unmarshal from %q: %w", filePath, errInvalidArgs)
	}

	return cfg, nil
}
