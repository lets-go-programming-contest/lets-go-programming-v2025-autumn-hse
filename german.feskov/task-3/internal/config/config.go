package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyInput  = errors.New("input-file is required")
	ErrEmptyOutput = errors.New("output-file is required")
	ErrSamePath    = errors.New("input-file and output-file must differ")
)

type Config struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

func (c Config) Validate() error {
	if c.Input == "" {
		return ErrEmptyInput
	}

	if c.Output == "" {
		return ErrEmptyOutput
	}

	inAbs, err := filepath.Abs(c.Input)
	if err != nil {
		return fmt.Errorf("could not get the absolute path for the input-file: %w", err)
	}

	outAbs, err := filepath.Abs(c.Output)
	if err != nil {
		return fmt.Errorf("could not get the absolute path for the output-file: %w", err)
	}

	if inAbs == outAbs {
		return ErrSamePath
	}

	return nil
}

func Load(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("read file %q: %w", filePath, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("unmarshal from %q: %w", filePath, err)
	}

	if err := cfg.Validate(); err != nil {
		return cfg, fmt.Errorf("unmarshal from %q: %w", filePath, err)
	}

	return cfg, nil
}
