package main

import (
	"flag"

	"github.com/Olesia.Ol/task-3/internal/config"
	"github.com/Olesia.Ol/task-3/internal/currency"
	"github.com/Olesia.Ol/task-3/internal/model"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to YAML config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic("Cannot load config: " + err.Error())
	}

	currencies, err := currency.Read[model.Currency](cfg.InputFile, "Valute")
	if err != nil {
		panic("Cannot read currencies: " + err.Error())
	}

	currency.Sort(currencies)

	if err := currency.WriteJSON(cfg.OutputFile, currencies); err != nil {
		panic("Cannot write JSON file: " + err.Error())
	}
}
