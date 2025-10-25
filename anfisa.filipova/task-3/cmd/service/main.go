package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Anfisa111/task-3/internal/config"
	"github.com/Anfisa111/task-3/internal/valutes"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Recovered from panic: %v\n", r)
		}
	}()

	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	config, err := config.LoadConfigYAML(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Cannot load config yaml: %v", err))
	}
	if _, err := os.Stat(config.InputFile); os.IsNotExist(err) {
		panic(fmt.Sprintf("Inputfile is not existing: %v", err))
	}

	currencies, err := valutes.DecodeValuteXML(config.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Cannot decode valute XML: %v", err))
	}

	valutes.SortCurrencies(currencies)

	err = valutes.WriteCurrenciesJSON(currencies, config.OutputFile)
	if err != nil {
		panic(fmt.Sprintf("Cannot write currencies to JSON: %v", err))
	}
}
