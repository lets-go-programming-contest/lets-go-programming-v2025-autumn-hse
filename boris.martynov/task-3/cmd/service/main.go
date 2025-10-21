package main

import (
	"github.com/JingolBong/task-3/internal/config"
	"github.com/JingolBong/task-3/internal/jsonwriter"
	"github.com/JingolBong/task-3/internal/xmlparser"
)

func main() {
	cfg, err := config.ConfigLoader()
	if err != nil {
		panic(err)
	}

	valCurs, err := xmlparser.Xmlparser(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	if err := jsonwriter.Jsonwrite(valCurs, cfg.OutputFile); err != nil {
		panic(err)
	}
}
