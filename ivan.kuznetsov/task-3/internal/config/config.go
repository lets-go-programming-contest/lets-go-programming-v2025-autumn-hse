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

func Load() *Config {
	configFlag := flag.String("config", "", "Config file path")
	flag.Parse()

	if *configFlag == "" {
		panic("invalid config file path")
	}

	configFile, err := os.Open(*configFlag)
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		err = configFile.Close()
		if err != nil {
			fmt.Println("error closing file", err)
		}
	}()

	var configData Config

	decoder := yaml.NewDecoder(configFile)

	err = decoder.Decode(&configData)
	if err != nil {
		panic(err.Error())
	}

	return &configData
}
