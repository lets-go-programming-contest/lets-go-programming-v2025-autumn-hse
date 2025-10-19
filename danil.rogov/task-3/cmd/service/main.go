package main

import (
	"cmp"
	"slices"

	"github.com/Tapochek2894/task-3/internal/config"
	"github.com/Tapochek2894/task-3/internal/jsonoutput"
	"github.com/Tapochek2894/task-3/internal/valute"
	"github.com/Tapochek2894/task-3/internal/xmlinput"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
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
