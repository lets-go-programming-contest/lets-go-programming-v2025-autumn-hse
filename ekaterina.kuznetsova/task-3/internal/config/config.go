package config

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfigYaml() (Config, error) {
	configPath := flag.String("config", "", "path of config file")
	flag.Parse()

	dataConfig, err := os.ReadFile(*configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error read config: %w", err)
	}

	var config Config

	err = yaml.Unmarshal(dataConfig, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing yaml: %w", err)
	}

	return config, nil
}
