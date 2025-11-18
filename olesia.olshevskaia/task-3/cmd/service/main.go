package main

import (
	"flag"
	"fmt"

	"github.com/Olesia.Ol/task-3/internal/config"
	"github.com/Olesia.Ol/task-3/internal/currency"
	"github.com/Olesia.Ol/task-3/internal/model"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to YAML config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Println("Cannot load config:", err)
		return
	}

	currencies, err := currency.Read[model.Currency](cfg.InputFile, "Valute")
	if err != nil {
		fmt.Println("Cannot read currencies:", err)
		return
	}

	currency.Sort(currencies)

	if err := currency.WriteJSON(cfg.OutputFile, currencies); err != nil {
		fmt.Println("Cannot write JSON file:", err)
		return
	}
}
