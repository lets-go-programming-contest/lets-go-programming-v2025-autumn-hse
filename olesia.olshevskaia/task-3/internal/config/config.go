package config

import (
	"os"

	"github.com/Olesia.Ol/task-3/internal/model"
	"gopkg.in/yaml.v3"
)

func Load(path string) (model.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return model.Config{}, err
	}

	var cfg model.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return model.Config{}, err
	}

	return cfg, nil
}
