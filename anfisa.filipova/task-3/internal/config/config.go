package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"ouput-file"`
}

func LoadConfig(configPath string) *Config {
	file, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshaling file: %v", err))
	}

	return &config
}
