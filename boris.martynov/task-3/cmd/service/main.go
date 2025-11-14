package main

import (
	"flag"
	"sort"

	"github.com/JingolBong/task-3/internal/config"
	"github.com/JingolBong/task-3/internal/jsonwriter"
	"github.com/JingolBong/task-3/internal/valuteinfo"
	"github.com/JingolBong/task-3/internal/xmlparser"
)

const dirPerm = 0o755

func main() {
	configPath := flag.String("config", "config.yaml", "config file path")
	flag.Parse()

	cfg, err := config.ConfigLoader(*configPath)
	if err != nil {
		panic(err)
	}

	var valCurs valuteinfo.ValuteCurs
	err = xmlparser.XMLParse(cfg.InputFile, &valCurs)

	if err != nil {
		panic(err)
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	if err := jsonwriter.JSONWrite(valCurs.Valutes, cfg.OutputFile, dirPerm); err != nil {
		panic(err)
	}
}
