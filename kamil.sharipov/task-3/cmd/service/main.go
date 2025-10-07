package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/kamilSharipov/task-3/internal/config"
	json "github.com/kamilSharipov/task-3/internal/json_formatter"
	xml "github.com/kamilSharipov/task-3/internal/xml_parser"
)

var (
	errNoConfigPath error = errors.New("no config path")
)

func main() {
	configPath := flag.String("config", "", "config file path")
	flag.Parse()

	if *configPath == "" {
		fmt.Println(errNoConfigPath)

		return
	}

	config, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Println(err)

		return
	}

	xmlData, err := os.ReadFile(config.InputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)

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
