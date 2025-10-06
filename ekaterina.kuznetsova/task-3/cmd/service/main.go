package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	XMLName  xml.Name `xml:"Valute"`
	NumCode  string   `xml:"NumCode"`
	CharCode string   `xml:"CharCode"`
	Value    string   `xml:"Value"`
}

type ValuteCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type ValuteJSON struct {
	NumCode  int     `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Value    float64 `json:"Value"`
}

type ValuteCursJSON struct {
	Valutes []ValuteJSON `json:"Valutes"`
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

	err = yaml.Unmarshal(dataConfig, &config)
	if err != nil {
		fmt.Printf("Error parsing yaml: %v\n", err)

		return
	}

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

	var valutesOutput ValuteCursJSON

	for _, valute := range valCurs.Valutes {
		numCode, _ := strconv.Atoi(valute.NumCode)
		value, _ := strconv.ParseFloat(strings.ReplaceAll(valute.Value, ",", "."), 64)
		valutesOutput.Valutes = append(valutesOutput.Valutes, ValuteJSON{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	sort.Slice(valutesOutput.Valutes, func(i, j int) bool {
		return valutesOutput.Valutes[i].Value > valutesOutput.Valutes[j].Value
	})

	outputJSON, err := json.MarshalIndent(valutesOutput, "", "    ")
	if err != nil {
		panic(fmt.Sprintf("Error marshaling JSON: %v", err))
	}

	err = os.MkdirAll(filepath.Dir(config.OutputFile), 0750)
	if err != nil {
		panic(fmt.Sprintf("Error creating directory: %v", err))
	}

	file, err := os.Create(config.OutputFile)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)

		return
	}
	defer file.Close()

	_, err = file.Write(outputJSON)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)

		return
	}

	// err = os.WriteFile(config.OutputFile, outputJSON, 0600)
	// if err != nil {
	// 	panic(fmt.Sprintf("Error writing output file: %v", err))
	// }
}
