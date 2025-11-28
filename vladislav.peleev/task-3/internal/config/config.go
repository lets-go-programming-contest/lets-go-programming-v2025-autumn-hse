package config

import (
	"fmt"
	"os"

	"github.com/VlasfimosY/task-3/internal/models"
	"gopkg.in/yaml.v3"
)

func Load(configPath *string) (*models.Config, error) {
	data, err := os.ReadFile(*configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	var cfg models.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("cannot unmarshal YAML: %w", err)
	}

	if cfg.DirPerm == 0 {
		cfg.DirPerm = 0o755
	}

	if cfg.FilePerm == 0 {
		cfg.FilePerm = 0o644
	}

	return &cfg, nil
}
