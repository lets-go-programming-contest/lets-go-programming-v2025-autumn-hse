package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Tapochek2894/task-3/internal/jsonoutput"
	"github.com/Tapochek2894/task-3/internal/xmlinput"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input_file"`
	OutputFile string `yaml:"output_file"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file %s: %w", path, err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error unmarshalling from %s: %w", path, err)
	}

	return config, nil
}

func main() {
	configPath := flag.String("config", "local.yaml", "path to config file")
	flag.Parse()

	cfg, err := Load(*configPath)
	if err != nil {
		panic("Config read error: " + err.Error())
	}

	input, err := xmlinput.ReadXML(cfg.InputFile)
	if err != nil {
		panic("XML read error: " + err.Error())
	}

	input.Sort(true)

	err = jsonoutput.CreateValuteCursJSON(cfg.OutputFile, input)
	if err != nil {
		panic("JSON write error: " + err.Error())
	}
}
