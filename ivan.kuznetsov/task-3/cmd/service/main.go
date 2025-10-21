package main

import (
	"flag"
	"sort"

	"github.com/kuzid-17/task-3/internal/config"
	"github.com/kuzid-17/task-3/internal/jsonwriter"
	"github.com/kuzid-17/task-3/internal/xmlparser"
)

func main() {
	configFlag := flag.String("config", "config.yaml", "Config file path")
	flag.Parse()

	config, err := config.Load(configFlag)
	if err != nil {
		panic(err)
	}

	valCurs, err := xmlparser.ParseXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	err = jsonwriter.WriteJSON(config.OutputFile, valCurs.Valutes)
	if err != nil {
		panic(err)
	}
}
