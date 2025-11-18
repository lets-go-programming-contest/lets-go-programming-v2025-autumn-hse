package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Anfisa111/task-3/internal/config"
	"github.com/Anfisa111/task-3/internal/jsonwriter"
	"github.com/Anfisa111/task-3/internal/valutes"
	"github.com/Anfisa111/task-3/internal/xmlreader"
)

const fileOpenPermission = 0o755

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Recovered from panic: %v\n", r)
		}
	}()

	configPath := flag.String("config", "config.yaml", "Path to configuration file")
	flag.Parse()

	config, err := config.LoadConfigYAML(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Cannot load config yaml: %v", err))
	}

	var valCurs valutes.ValCurs

	err = xmlreader.ReadFileXML(config.InputFile, &valCurs)
	if err != nil {
		panic(fmt.Sprintf("Cannot read Inputfile: %v", err))
	}

	currencies := valCurs.Valutes
	valutes.SortValutes(currencies)

	err = jsonwriter.WriteFileJSON(currencies, config.OutputFile, fileOpenPermission)
	if err != nil {
		panic(fmt.Sprintf("Cannot write currencies to JSON: %v", err))
	}
}
