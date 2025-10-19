package main

import (
	"cmp"
	"flag"
	"slices"

	"github.com/Tapochek2894/task-3/internal/config"
	"github.com/Tapochek2894/task-3/internal/jsonoutput"
	"github.com/Tapochek2894/task-3/internal/valute"
	"github.com/Tapochek2894/task-3/internal/xmlinput"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic("Config error: " + err.Error())
	}

	input, err := xmlinput.ReadXML(cfg.InputFile)
	if err != nil {
		panic("XML read error: " + err.Error())
	}

	slices.SortFunc(input, func(a, b valute.ValuteInfo) int {
		return -cmp.Compare(a.Value, b.Value)
	})

	err = jsonoutput.WriteJSON(cfg.OutputFile, input)
	if err != nil {
		panic("JSON write error: " + err.Error())
	}
}
