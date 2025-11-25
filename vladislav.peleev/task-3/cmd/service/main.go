package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/VlasfimosY/task-3/internal/config"
	"github.com/VlasfimosY/task-3/internal/jsonwriter"
	"github.com/VlasfimosY/task-3/internal/xmlparser"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to configuration file")
	flag.Parse()

	cfg, err := config.Load(configPath)
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}

	currencies, err := xmlparser.DecodeXML(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Error decoding XML: %v", err))
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	if err := jsonwriter.SaveJSON(cfg.OutputFile, currencies, 0o755); err != nil {
		panic(fmt.Sprintf("Error saving JSON: %v", err))
	}

	fmt.Println("Successfully processed currencies data")
}
