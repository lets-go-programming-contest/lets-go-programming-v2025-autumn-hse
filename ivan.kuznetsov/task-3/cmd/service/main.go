package main

import (
	"flag"
	"sort"

	"github.com/kuzid-17/task-3/internal/config"
	"github.com/kuzid-17/task-3/internal/jsonwriter"
	"github.com/kuzid-17/task-3/internal/xmlparser"
)

func main() {
	configFlag := flag.String("config", "", "Config file path")
	flag.Parse()

	config := config.Load(configFlag)

	valCurs := xmlparser.ParseXML(config.InputFile)

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	jsonwriter.WriteJSON(config.OutputFile, valCurs.Valutes)
}
