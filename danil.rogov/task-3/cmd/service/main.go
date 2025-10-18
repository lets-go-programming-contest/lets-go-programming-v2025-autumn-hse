package main

import (
	"cmp"
	"flag"
	"fmt"
	"slices"

	"github.com/Tapochek2894/task-3/internal/config"
	"github.com/Tapochek2894/task-3/internal/jsonoutput"
	"github.com/Tapochek2894/task-3/internal/valute"
	"github.com/Tapochek2894/task-3/internal/xmlinput"
)

func main() {
	configPath := flag.String("config", "local.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	input, err := xmlinput.ReadXML(cfg.InputFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	slices.SortFunc(input, func(a, b valute.ValuteInfo) int {
		return -cmp.Compare(a.Value, b.Value)
	})

	err = jsonoutput.WriteJSON(cfg.OutputFile, input)
	if err != nil {
		fmt.Println(err)
		return
	}
}
