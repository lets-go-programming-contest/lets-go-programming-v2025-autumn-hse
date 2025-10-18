package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kamilSharipov/task-3/internal/config"
	json "github.com/kamilSharipov/task-3/internal/json_formatter"
	xml "github.com/kamilSharipov/task-3/internal/xml_parser"
)

const (
	dirPermissions = 0o755
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
			os.Exit(1)
		}
	}()

	configPath := flag.String("config", "", "config file path")
	flag.Parse()

	if *configPath == "" {
		panic("no config path")
	}

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(config.InputFile); err != nil {
		panic(err)
	}

	xmlData, err := os.ReadFile(config.InputFile)
	if err != nil {
		panic(err)
	}

	valCurs, err := xml.ParseXML(xmlData)
	if err != nil {
		panic(err)
	}

	bytes, err := json.FormateJSON(valCurs)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Dir(config.OutputFile), dirPermissions)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(config.OutputFile)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = outputFile.Close()
	}()

	_, err = outputFile.Write(bytes)
	if err != nil {
		panic(err)
	}
}
