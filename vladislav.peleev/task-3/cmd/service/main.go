package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input_file"`
	OutputFile string `yaml:"output_file"`
}

type Currency struct {
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	Currencies []Currency `xml:"Valute"`
}

type CurrencyJSON struct {
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

var (
	errConfigNotFound = errors.New("no such config file")
	errFileNotFound   = errors.New("no such data file")
)

func loadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("open %s: %w", configPath, errConfigNotFound)
		}

		return nil, fmt.Errorf("open %s: %w", configPath, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &config, nil
}

func parseXML(filePath string) ([]CurrencyJSON, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("open %s: %w", filePath, errFileNotFound)
		}

		return nil, fmt.Errorf("open %s: %w", filePath, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read xml: %w", err)
	}

	var valCurs ValCurs
	if err := xml.Unmarshal(data, &valCurs); err != nil {
		return nil, fmt.Errorf("unmarshal xml: %w", err)
	}

	result := make([]CurrencyJSON, 0, len(valCurs.Currencies))

	for _, currency := range valCurs.Currencies {

		valStr := strings.Replace(currency.Value, ",", ".", -1)
		valStr = strings.TrimSpace(valStr)

		if valStr == "" {
			continue
		}

		var num float64
		_, err = fmt.Sscanf(valStr, "%f", &num)
		if err != nil {
			continue
		}

		result = append(result, CurrencyJSON{
			CharCode: currency.CharCode,
			Value:    num,
		})
	}

	return result, nil
}

func saveJSON(data []CurrencyJSON, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create json: %w", err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	config, err := loadConfig(*configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading config:", err)
		os.Exit(1)
	}

	xmlPath := filepath.Clean(config.InputFile)
	jsonPath := filepath.Clean(config.OutputFile)

	data, err := parseXML(xmlPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing XML:", err)
		os.Exit(1)
	}

	if err := saveJSON(data, jsonPath); err != nil {
		fmt.Fprintln(os.Stderr, "error saving JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Conversion completed successfully!")
}
