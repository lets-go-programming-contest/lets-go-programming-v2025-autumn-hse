package config

import (
	"fmt"
	"os"

	"github.com/Olesia.Ol/task-3/internal/model"
	"gopkg.in/yaml.v3"
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

	return cfg, nil
}
