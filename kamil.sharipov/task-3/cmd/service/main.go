package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/kamilSharipov/task-3/internal/config"
	jsonutil "github.com/kamilSharipov/task-3/internal/json_util"
	"github.com/kamilSharipov/task-3/internal/model"
	xmlutil "github.com/kamilSharipov/task-3/internal/xml_util"
)

type ValCurs struct {
	Valutes []model.Valute `xml:"Valute"`
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	configPath := parseFlags()
	cfg, err := loadConfig(configPath)
	panicOnErr(err)

	valCurs, err := xmlutil.ReadInput[ValCurs](cfg.InputFile)
	panicOnErr(err)

	sortValCurs(valCurs.Valutes)

	err = jsonutil.WriteOutput(cfg.OutputFile, valCurs.Valutes)
	panicOnErr(err)
}

func parseFlags() string {
	configPath := flag.String("config", "./config.yaml", "config file path")
	flag.Parse()

	return *configPath
}

func loadConfig(path string) (config.Config, error) {
	cfg, err := config.LoadConfig(path)
	if err != nil {
		return config.Config{}, fmt.Errorf("failed to load config from %q: %w", path, err)
	}

	return cfg, nil
}

func sortValCurs(valCurs []model.Valute) {
	sort.Slice(valCurs, func(i, j int) bool {
		return float64(valCurs[i].Value) > float64(valCurs[j].Value)
	})
}
