package main

import (
	"flag"

	"github.com/Olesia.Ol/task-3/internal/config"
	"github.com/Olesia.Ol/task-3/internal/currency"
)

func main() {

	configPath := flag.String("config", "config.yaml", "Path to YAML config file")
	flag.Parse()

	cfg := config.Load(*configPath)

	currencies := currency.Read(cfg.InputFile)
	currency.Sort(currencies)
	currency.WriteJSON(currencies, cfg.OutputFile)
}
