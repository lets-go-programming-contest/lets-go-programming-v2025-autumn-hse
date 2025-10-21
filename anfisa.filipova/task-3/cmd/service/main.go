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

	config := config.LoadConfigYAML(*configPath)
	if _, err := os.Stat(config.InputFile); os.IsNotExist(err) {
		panic(fmt.Sprintf("Inputfile is not existing: %v", err))
	}

	currencies := valutes.DecodeValuteXML(config.InputFile)
	valutes.SortCurrencies(currencies)
	valutes.WriteCurrenciesJSON(currencies, config.OutputFile)
}
