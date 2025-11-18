package main

import (
	"flag"

	"github.com/Tapochek2894/task-3/internal/config"
	"github.com/Tapochek2894/task-3/internal/jsonoutput"
	"github.com/Tapochek2894/task-3/internal/valute"
	"github.com/Tapochek2894/task-3/internal/xmlinput"
)

const outputMode = 0o777

func main() {
	configPath := flag.String("config", "local.yaml", "path to config file")
	flag.Parse()

	config, err := config.Load(*configPath)
	if err != nil {
		panic("Config read error: " + err.Error())
	}

	var data valute.Valutes

	err = xmlinput.ReadXML(config.InputFile, &data)
	if err != nil {
		panic("XML read error: " + err.Error())
	}

	data.Sort(true)

	err = jsonoutput.WriteJSON(config.OutputFile, data.Valutes, outputMode)
	if err != nil {
		panic("JSON write error: " + err.Error())
	}
}
