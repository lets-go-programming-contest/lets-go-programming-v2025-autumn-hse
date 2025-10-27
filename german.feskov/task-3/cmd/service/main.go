package main

import (
	"cmp"
	"flag"
	"slices"

	"github.com/6ermvH/german.feskov/task-3/internal/config"
	"github.com/6ermvH/german.feskov/task-3/internal/jsonout"
	"github.com/6ermvH/german.feskov/task-3/internal/valute"
	"github.com/6ermvH/german.feskov/task-3/internal/xmlin"
)

func main() {
	configPath := flag.String("config", "configs/local.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	valCurs, err := xmlin.Read(cfg.Input)
	if err != nil {
		panic(err)
	}

	slices.SortFunc(valCurs.Valutes, func(a, b valute.Valute) int {
		return -cmp.Compare(a.Value, b.Value)
	})

	if err := jsonout.Write(cfg.Output, valCurs.Valutes); err != nil {
		panic(err)
	}
}
