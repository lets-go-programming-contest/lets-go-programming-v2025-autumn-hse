package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Currency struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

type CurrencyJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	if *configPath == "" {
		panic("Config file path is required")
	}

	config, err := loadConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}

	if _, err := os.Stat(config.InputFile); os.IsNotExist(err) {
		panic(fmt.Sprintf("open %s: no such file or directory", config.InputFile))
	}

	currencies, err := decodeXML(config.InputFile)
	if err != nil {
		panic(fmt.Sprintf("Error decoding XML: %v", err))
	}

	sortCurrencies(currencies)

	err = saveJSON(config.OutputFile, currencies)
	if err != nil {
		panic(fmt.Sprintf("Error saving JSON: %v", err))
	}

	fmt.Println("Successfully processed currencies data")
}

func loadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("open %s: no such file or directory", configPath)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal YAML: %w", err)
	}

	return &config, nil
}

func decodeXML(filePath string) ([]CurrencyJSON, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open %s: no such file or directory", filePath)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return input, nil
	}

	var valCurs ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal XML: %w", err)
	}

	result := make([]CurrencyJSON, 0, len(valCurs.Currencies))
	for _, currency := range valCurs.Currencies {
		valStr := strings.Replace(currency.Value, ",", ".", -1)
		value, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse value '%s': %w", currency.Value, err)
		}

		numStr := strings.TrimSpace(currency.NumCode)
		numCode := 0
		if numStr != "" {
			if n, err := strconv.Atoi(numStr); err == nil {
				numCode = n
			}
		}

		result = append(result, CurrencyJSON{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(currency.CharCode),
			Value:    value,
		})
	}

	return result, nil
}

func sortCurrencies(currencies []CurrencyJSON) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

func saveJSON(outputPath string, currencies []CurrencyJSON) error {
	dir := filepath.Dir(outputPath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("cannot create directory: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("cannot create output file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(currencies)
	if err != nil {
		return fmt.Errorf("cannot encode JSON: %w", err)
	}

	return nil
}
