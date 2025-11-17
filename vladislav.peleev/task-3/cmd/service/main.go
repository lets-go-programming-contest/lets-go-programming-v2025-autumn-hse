package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/VlasfimosY/task-3/internal/config"
	"github.com/VlasfimosY/task-3/internal/jsonwriter"
	"github.com/VlasfimosY/task-3/internal/models"
	"github.com/VlasfimosY/task-3/internal/xmlparser"
)

func main() {
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	if *configPath == "" {
		panic("Config file path is required")
	}

	config, err := config.Load(configPath)
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}

	if _, err := os.Stat(config.InputFile); os.IsNotExist(err) {
		panic(fmt.Sprintf("open %s: no such file or directory", config.InputFile))
	}

	currencies, err := xmlparser.DecodeXML(config.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Error decoding XML: %v", err))
	}

	sortCurrencies(currencies)

	err = jsonwriter.SaveJSON(config.OutputFile, currencies)
	if err != nil {
		panic(fmt.Sprintf("Error saving JSON: %v", err))
	}

	fmt.Println("Successfully processed currencies data")
}

func sortCurrencies(currencies []models.CurrencyJSON) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}
