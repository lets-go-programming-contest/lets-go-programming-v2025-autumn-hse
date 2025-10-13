package main

import (
	"sort"

	"github.com/kuzid-17/task-3/internal/config"
	"github.com/kuzid-17/task-3/internal/jsonwriter"
	"github.com/kuzid-17/task-3/internal/xmlparser"
)

func main() {
	config := config.Load()
	valCurs := xmlparser.ParseXML(config.InputFile)

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	jsonwriter.WriteJSON(config.OutputFile, valCurs.Valutes)
}
