package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/Olesia.Ol/task-3/internal/model"
)

func Load(path string) model.Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("Can not read config file: " + path)
	}

	var cfg model.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic("Can not parse YAML: " + err.Error())
	}

	return cfg
}
