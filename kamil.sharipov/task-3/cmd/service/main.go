package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

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

	config := config.LoadConfig(*configPath)

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

	sort.Slice(valCurs, func(i, j int) bool {
		value1, err1 := valCurs[i].GetValue()
		value2, err2 := valCurs[j].GetValue()

		if err1 != nil || err2 != nil {
			panic(err1)
		}

		return value1 > value2
	})

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
