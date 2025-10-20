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

func Load(configFlag *string) *Config {
	configFile, err := os.Open(*configFlag)
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		err = configFile.Close()
		if err != nil {
			panic(err.Error() + "AAA")
		}
	}()

	var configData Config

	decoder := yaml.NewDecoder(configFile)

	err = decoder.Decode(&configData)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &configData
}
