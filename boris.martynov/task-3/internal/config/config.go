package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ConfigLoader(configPath string) (Config, error) {
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("while reading config file: %w", err)
	}

	var config Config

	if err := yaml.Unmarshal(configData, &config); err != nil {
		return Config{}, fmt.Errorf("while parsing yaml: %w", err)
	}

	return config, nil
}
