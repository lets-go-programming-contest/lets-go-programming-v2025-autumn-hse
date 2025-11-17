package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/kef1rch1k/task-3/internal/config"
	"github.com/kef1rch1k/task-3/internal/jsonwriter"
	"github.com/kef1rch1k/task-3/internal/parser"
)

func main() {
	configPath := flag.String("config", "", "Path to config YAML file")
	flag.Parse()

	if *configPath == "" {
		panic("Missing --config flag with path to config file")
	}

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	valutes, err := parser.ParseXML(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse XML: %v", err))
	}

	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})

	err = jsonwriter.WriteToFile(valutes, cfg.OutputFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to write JSON: %v", err))
	}

	fmt.Println("Output written to", cfg.OutputFile)
}
