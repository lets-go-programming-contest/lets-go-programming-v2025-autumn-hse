package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ekaterina-101/task-3/internal/parsexml"
	"github.com/Ekaterina-101/task-3/internal/writerjson"
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

func main() {
	config, err := LoadConfigYaml()
	if err != nil {
		panic(err)
	}

	valCurs, err := parsexml.ParseValuteCursXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	outputJSON, err := writerjson.CreateValuteCursJSON(valCurs)
	if err != nil {
		panic(err)
	}

	err = writerjson.WriteFileJSON(config.OutputFile, outputJSON)
	if err != nil {
		panic(err)
	}
}
