package config

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var (
	errFailedToOpenFile     error = errors.New("failed to open file")
	errFailedToDecodeConfig error = errors.New("failed to decode config")
	errNoInputFile          error = errors.New("no input file")
	errNoOutputFile         error = errors.New("no output file")
	errInputFileDoesntExist error = errors.New("input file does not exist")
	errFailedToCreateDir    error = errors.New("failed to create output directory")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errFailedToOpenFile
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, errFailedToDecodeConfig
	}

	if config.InputFile == "" {
		return nil, errNoInputFile
	}
	if config.OutputFile == "" {
		return nil, errNoOutputFile
	}

	if _, err := os.Stat(config.InputFile); err != nil {
		return nil, errInputFileDoesntExist
	}

	if err := os.MkdirAll(filepath.Dir(config.OutputFile), 0755); err != nil {
		return nil, errFailedToCreateDir
	}

	return &config, nil
}
