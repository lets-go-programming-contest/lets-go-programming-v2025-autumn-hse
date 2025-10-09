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

func main() {
	configPath := flag.String("config", "", "config file path")
	flag.Parse()

	if *configPath == "" {
		fmt.Println("no config path")

		return
	}

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Println(err)

		return
	}

	if _, err := os.Stat(config.InputFile); err != nil {
		fmt.Println("input file does not exist:", err)

		return
	}

	xmlData, err := os.ReadFile(config.InputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)

		return
	}

	valCurs, err := xml.ParseXML(xmlData)
	if err != nil {
		fmt.Println(err)

		return
	}

	bytes, err := json.FormateJSON(valCurs)
	if err != nil {
		fmt.Println(err)

		return
	}

	if err := os.MkdirAll(filepath.Dir(config.OutputFile), 0755); err != nil {
		fmt.Println("Error creating output directory:", err)

		return
	}

	outputFile, err := os.Create(config.OutputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)

		return
	}
	defer outputFile.Close()

	_, err = outputFile.Write(bytes)
	if err != nil {
		fmt.Println("Error writing to output file:", err)

		return
	}
}
