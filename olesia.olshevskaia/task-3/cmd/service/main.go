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
		panic("Can not open XML file: " + path + " - " + err.Error())
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic("Error closing XML file: " + err.Error())
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var xmlData CurrenciesXML
	if err := decoder.Decode(&xmlData); err != nil {
		panic("Error parsing XML: " + err.Error())
	}

	result := make([]Currency, 0, len(xmlData.Currencies))

	for _, val := range xmlData.Currencies {
		num := 0
		value := 0.0

		if strings.TrimSpace(val.CodeNum) != "" {
			if n, err := strconv.Atoi(val.CodeNum); err == nil {
				num = n
			}
		}

		if strings.TrimSpace(val.RateValue) != "" {
			str := strings.ReplaceAll(val.RateValue, ",", ".")
			if v, err := strconv.ParseFloat(str, 64); err == nil {
				value = v
			}
		}

		result = append(result, Currency{
			CodeNum:   num,
			CodeChar:  val.CodeChar,
			RateValue: value,
			HasValue:  strings.TrimSpace(val.RateValue) != "",
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].RateValue > result[j].RateValue
	})

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
