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

type CurrencyXML struct {
	CodeNum   string `xml:"NumCode"`
	CodeChar  string `xml:"CharCode"`
	RateValue string `xml:"Value"`
}

type CurrenciesXML struct {
	Currencies []CurrencyXML `xml:"Valute"`
}

type Currency struct {
	CodeNum   int     `json:"num_code"`
	CodeChar  string  `json:"char_code"`
	RateValue float64 `json:"value"`
}

func loadConfig(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Can not read config file: %s", path))
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Sprintf("Can not parse YAML: %v", err))
	}

	return cfg
}

func ReadCurrencies(path string) []Currency {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Can not parse XML: %v", path))
	}

	var xmlData CurrenciesXML
	if err := xml.Unmarshal(data, &xmlData); err != nil {
		panic(fmt.Sprintf("Error of parsing XML: %v", err))
	}

	result := make([]Currency, 0, len(xmlData.Currencies))
	for _, c := range xmlData.Currencies {
		num, err := strconv.Atoi(c.CodeNum)
		if err != nil {
			panic(fmt.Sprintf("NodeNum err: %s", c.CodeNum))
		}
		valueString := strings.ReplaceAll(c.RateValue, ",", ".")
		value, err := strconv.ParseFloat(valueString, 64)

		if err != nil {
			panic(fmt.Sprintf("Incorrect exchange rate value: %s", c.RateValue))
		}

		result = append(result, Currency{
			CodeNum:   num,
			CodeChar:  c.CodeChar,
			RateValue: value,
		})
	}
	return result
}

func SortCurrencies(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].RateValue > currencies[j].RateValue
	})
}

func WriteJSON(currencies []Currency, outputPath string) {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Sprintf("I can't create a directory: %v", err))
	}

	file, err := os.Create(outputPath)
	if err != nil {
		panic(fmt.Sprintf("I can't create a JSON file: %v", err))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(currencies); err != nil {
		panic(fmt.Sprintf("Write error JSON: %v", err))
	}
}

func main() {
	configPath := flag.String("config", "config.yaml", "Path to YAML config file")
	flag.Parse()

	cfg := loadConfig(*configPath)

	currencies := ReadCurrencies(cfg.InputFile)

	SortCurrencies(currencies)

	WriteJSON(currencies, cfg.OutputFile)
}
