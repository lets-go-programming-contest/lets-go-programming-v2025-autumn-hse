package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var errInvalidFileStruct = errors.New("invalid file struct")

const errMsg = "while load config %q: %w"

type Config struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

func Load(filePath string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(filePath)
	if err != nil {
		return cfg, fmt.Errorf(errMsg, filePath, err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf(errMsg, filePath, err)
	}

	if cfg.Input == "" || cfg.Output == "" {
		return cfg, fmt.Errorf(errMsg, filePath, errInvalidFileStruct)
	}

	return cfg, nil
}
