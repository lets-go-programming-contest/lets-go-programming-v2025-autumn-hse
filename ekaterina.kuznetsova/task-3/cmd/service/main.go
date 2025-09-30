package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"encoding/json"
	"encoding/xml"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValuteCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type ValuteJSON struct {
	NumCode  int     `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Value    float64 `json:"Value"`
}

type ValuteCursJSON struct {
	Valutes []ValuteJSON
}

func main() {
	configPath := flag.String("config", "", "path of config file")
	flag.Parse()

	dataConfig, err := os.ReadFile(*configPath)
	if err != nil {
		fmt.Printf("read config: %v\n", err)

		return
	}

	var config Config
	yaml.Unmarshal(dataConfig, &config)

	var valCurs ValuteCurs

	dataXML, err := os.ReadFile(config.InputFile)
	if err != nil {
		fmt.Printf("read config: %v\n", err)

		return
	}

	err = xml.Unmarshal(dataXML, &valCurs)
	if err != nil {
		fmt.Printf("Error parsing XML: %v\n", err)
		return
	}

	// sort.Slice(valCurs.Valutes, func(i, j int) bool {
	// 	value1 := strings.Replace(valCurs.Valutes[i].Value, ",", ".", -1)
	// 	value2 := strings.Replace(valCurs.Valutes[j].Value, ",", ".", -1)

	// 	float1, _ := strconv.ParseFloat(value1, 64)
	// 	float2, _ := strconv.ParseFloat(value2, 64)

	// 	return float1 > float2
	// })

	var valutesOutput []ValuteJSON

	for _, valute := range valCurs.Valutes {
		numCode, _ := strconv.Atoi(valute.NumCode)
		value, _ := strconv.ParseFloat(strings.Replace(valute.Value, ",", ".", -1), 64)
		valutesOutput = append(valutesOutput, ValuteJSON{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	sort.Slice(valutesOutput, func(i, j int) bool {
		return valutesOutput[i].Value > valutesOutput[j].Value
	})

	outputJSON, err := json.MarshalIndent(valutesOutput, "", "    ")
	if err != nil {
		panic(fmt.Sprintf("Error marshaling JSON: %v", err))
	}

	err = os.MkdirAll(filepath.Dir(config.OutputFile), 0700)
	if err != nil {
		panic(fmt.Sprintf("Error creating directory: %v", err))
	}

	err = os.WriteFile(config.OutputFile, outputJSON, 0700)
	if err != nil {
		panic(fmt.Sprintf("Error writing output file: %v", err))
	}
}
