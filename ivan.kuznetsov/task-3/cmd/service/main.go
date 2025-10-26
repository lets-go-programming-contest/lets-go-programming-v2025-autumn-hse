package main

import (
	"flag"
	"sort"

	"github.com/kuzid-17/task-3/internal/config"
	"github.com/kuzid-17/task-3/internal/jsonwriter"
	"github.com/kuzid-17/task-3/internal/models"
	"github.com/kuzid-17/task-3/internal/xmlparser"
)

const mkdirMode = 0o755

func main() {
	configFlag := flag.String("config", "config.yaml", "Config file path")
	flag.Parse()

	config, err := config.Load(configFlag)
	if err != nil {
		panic(err)
	}

	valCurs, err := xmlparser.ParseXML[models.ValCurs](config.InputFile)
	if err != nil {
		panic(err)
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	err = jsonwriter.WriteJSON(config.OutputFile, mkdirMode, valCurs.Valutes)
	if err != nil {
		panic(err)
	}
}
