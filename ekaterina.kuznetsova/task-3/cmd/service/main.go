package main

import (
	"flag"
	"sort"

	"github.com/Ekaterina-101/task-3/internal/config"
	"github.com/Ekaterina-101/task-3/internal/model"
	"github.com/Ekaterina-101/task-3/internal/parsexml"
	"github.com/Ekaterina-101/task-3/internal/writerjson"
)

const (
	dirMode  = 0o755
	fileMode = 0o600
)

func main() {
	configPath := flag.String("config", "config.yaml", "path of config file")
	flag.Parse()

	config, err := config.LoadConfigYaml(*configPath)
	if err != nil {
		panic(err)
	}

	valCurs, err := parsexml.ParseXML[model.ValuteCurs](config.InputFile)
	if err != nil {
		panic(err)
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	err = writerjson.ParseJSON(config.OutputFile, valCurs, dirMode, fileMode)
	if err != nil {
		panic(err)
	}
}
