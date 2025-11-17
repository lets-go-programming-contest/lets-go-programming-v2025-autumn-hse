package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/VlasfimosY/task-3/internal/models"
)

func Load(configPath *string) (*models.Config, error) {
	file, err := os.Open(*configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to close file %s: %v\n", *configPath, cerr)
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	var config models.Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("cannot unmarshal YAML: %w", err)
	}

	return &config, nil
}
