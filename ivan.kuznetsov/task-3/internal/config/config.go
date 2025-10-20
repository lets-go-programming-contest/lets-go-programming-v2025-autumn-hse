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

func Load(configFlag *string) (*Config, error) {
	configFile, err := os.Open(*configFlag)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}

	defer func() {
		err = configFile.Close()
		if err != nil {
			panic(err.Error())
		}
	}()

	var configData Config

	decoder := yaml.NewDecoder(configFile)

	err = decoder.Decode(&configData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode YAML %w", err)
	}

	return &configData, nil
}
