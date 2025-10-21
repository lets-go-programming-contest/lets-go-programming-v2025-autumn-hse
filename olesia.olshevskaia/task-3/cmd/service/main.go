package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
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
		panic("Can not read config file: " + path)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic("Can not parse YAML: " + err.Error())
	}

	return cfg
}

func ReadCurrencies(path string) []Currency {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("Can not parse XML: " + path)
	}

	var xmlData CurrenciesXML
	if err := xml.Unmarshal(data, &xmlData); err != nil {
		panic("Error of parsing XML: " + err.Error())
	}

	result := make([]Currency, 0, len(xmlData.Currencies))

	for _, val := range xmlData.Currencies {
		if strings.TrimSpace(val.CodeNum) == "" {
			continue
		}

		num, err := strconv.Atoi(val.CodeNum)
		if err != nil {
			panic("NodeNum err: " + val.CodeNum)
		}

		valueString := strings.ReplaceAll(val.RateValue, ",", ".")
		valueString = strings.TrimSpace(valueString)
		if valueString == "" {
			continue
		}

		value, err := strconv.ParseFloat(valueString, 64)
		if err != nil {
			panic("Incorrect exchange rate value: " + val.RateValue)
		}

		result = append(result, Currency{
			CodeNum:   num,
			CodeChar:  val.CodeChar,
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
		panic("I can't create a directory: " + err.Error())
	}

	file, err := os.Create(outputPath)
	if err != nil {
		panic("I can't create a JSON file:  " + err.Error())
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic("Error closing file:" + err.Error())
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		panic("Write error JSON: " + err.Error())
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
