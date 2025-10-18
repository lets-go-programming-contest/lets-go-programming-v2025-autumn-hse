package main

import (
	"github.com/Ekaterina-101/task-3/internal/config"
	"github.com/Ekaterina-101/task-3/internal/parsexml"
	"github.com/Ekaterina-101/task-3/internal/writerjson"
)

func main() {
	config, err := config.LoadConfigYaml()
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
