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

	valCurs, err := currency.Read[model.ValCurs](cfg.InputFile)
	if err != nil {
		panic("Cannot read currencies: " + err.Error())
	}

	currency.Sort(valCurs.Currencies)

	if err := currency.WriteJSON(cfg.OutputFile, valCurs.Currencies); err != nil {
		panic("Cannot write JSON file: " + err.Error())
	}
}
