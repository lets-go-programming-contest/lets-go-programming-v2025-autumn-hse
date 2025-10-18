package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	errNoInputFile  error = errors.New("no input file")
	errNoOutputFile error = errors.New("no output file")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *Config) Validate() error {
	if c.InputFile == "" {
		return errNoInputFile
	}
	if c.OutputFile == "" {
		return errNoOutputFile
	}
	return nil
}

func LoadConfig(path string) Config {
	file, err := os.Open(path)
	panicOnErr(err)
	defer file.Close()

	var config Config
	err = yaml.NewDecoder(file).Decode(&config)
	panicOnErr(err)

	config.Validate()
	panicOnErr(config.Validate())

	return config
}
