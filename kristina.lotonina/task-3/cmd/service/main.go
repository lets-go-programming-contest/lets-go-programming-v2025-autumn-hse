package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/kef1rch1k/task-3/internal/config"
	"github.com/kef1rch1k/task-3/internal/parser"
	"github.com/kef1rch1k/task-3/internal/utils"
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

	valutes, err := parser.ParseAndSortXML(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse XML: %v", err))
	}

	err = utils.EnsureDir(cfg.OutputFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to create directory: %v", err))
	}

	file, err := os.Create(cfg.OutputFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to create output file: %v", err))
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("Failed to close file: %v", err))
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(valutes)

	if err != nil {
		panic(fmt.Sprintf("Failed to encode JSON: %v", err))
	}

	fmt.Println("Output written to", cfg.OutputFile)
}
