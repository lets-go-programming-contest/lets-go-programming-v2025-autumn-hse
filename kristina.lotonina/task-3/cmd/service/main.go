package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/kef1rch1k/task-3/internal/config"
	"github.com/kef1rch1k/task-3/internal/jsonwriter"
	"github.com/kef1rch1k/task-3/internal/models"
	"github.com/kef1rch1k/task-3/internal/parser"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config YAML file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	var valCurs models.ValCurs

	err = parser.ParseXML(cfg.InputFile, &valCurs)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse XML: %v", err))
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	err = jsonwriter.WriteToFile(valCurs.Valutes, cfg.OutputFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to write JSON: %v", err))
	}

	fmt.Println("Output written to", cfg.OutputFile)
}
