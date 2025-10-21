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

	cfg := loadConfig()

	xmlData := readXMLFile(cfg.InputFile)

	valCurs := parseXMLData(xmlData)
	sortValCurs(valCurs)

	jsonBytes := formatJSON(valCurs)
	writeOutput(cfg.OutputFile, jsonBytes)
}

func loadConfig() config.Config {
	configPath := flag.String("config", "", "config file path")
	flag.Parse()

	if *configPath == "" {
		panic("no config path")
	}

	return config.LoadConfig(*configPath)
}

func readXMLFile(path string) []byte {
	if _, err := os.Stat(path); err != nil {
		panic(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}

func parseXMLData(data []byte) []xml.Valute {
	valCurs, err := xml.ParseXML(data)
	if err != nil {
		panic(err)
	}

	return valCurs
}

func sortValCurs(valCurs []xml.Valute) {
	sort.Slice(valCurs, func(i, j int) bool {
		value1, err1 := valCurs[i].GetValue()
		value2, err2 := valCurs[j].GetValue()

		if err1 != nil || err2 != nil {
			panic(err1)
		}

		return value1 > value2
	})
}

func formatJSON(valCurs []xml.Valute) []byte {
	bytes, err := json.FormateJSON(valCurs)
	if err != nil {
		panic(err)
	}

	return bytes
}

func writeOutput(outputPath string, data []byte) {
	err := os.MkdirAll(filepath.Dir(outputPath), dirPermissions)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = outputFile.Close()
	}()

	_, err = outputFile.Write(data)
	if err != nil {
		panic(err)
	}
}
