package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/Olesia.Ol/task-3/internal/model"
	"gopkg.in/yaml.v3"
)

var (
	ErrInputFileEmpty  = errors.New("input_file is empty")
	ErrOutputFileEmpty = errors.New("output_file is empty")
)

func Load(path string) (model.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return model.Config{}, fmt.Errorf("failed to read config file %q: %w", path, err)
	}

	var cfg model.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return model.Config{}, fmt.Errorf("failed to parse yaml config %q: %w", path, err)
	}

	if cfg.InputFile == "" {
		return model.Config{}, ErrInputFileEmpty
	}

	if cfg.OutputFile == "" {
		return model.Config{}, ErrOutputFileEmpty
	}

	return cfg, nil
}
