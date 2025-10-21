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

	"golang.org/x/net/html/charset"
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
	HasValue  bool    `json:"-"`
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
	file, err := os.Open(path)
	if err != nil {
		panic("Can not open XML file: " + path)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var xmlData CurrenciesXML
	if err := decoder.Decode(&xmlData); err != nil {
		panic("Error parsing XML: " + err.Error())
	}

	result := make([]Currency, 0, len(xmlData.Currencies))

	for _, val := range xmlData.Currencies {
		codeStr := strings.TrimSpace(val.CodeNum)
		if codeStr == "" {
			continue
		}

		num, err := strconv.Atoi(codeStr)
		if err != nil {
			continue
		}

		valueStr := strings.ReplaceAll(val.RateValue, ",", ".")
		valueStr = strings.TrimSpace(valueStr)

		c := Currency{
			CodeNum:  num,
			CodeChar: val.CodeChar,
			HasValue: false,
		}

		if valueStr != "" {
			value, err := strconv.ParseFloat(valueStr, 64)
			if err == nil {
				c.RateValue = value
				c.HasValue = true
			}
		}

		result = append(result, c)
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
