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
	InputPath  string `yaml:"input_path"`
	OutputPath string `yaml:"output_path"`
}

type ValCurs struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

type Currency struct {
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrencyJSON struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}

var (
	errNoConfigFile = errors.New("no such file or directory")
	errNoXMLFile    = errors.New("no such XML file or directory")
)

func loadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", configPath, errNoConfigFile)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}

	return &cfg, nil
}

func parseXML(filePath string) ([]CurrencyJSON, error) {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", filePath, errNoXMLFile)
	}
	defer xmlFile.Close()

	data, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, fmt.Errorf("read xml file: %w", err)
	}

	var valCurs ValCurs
	if err := xml.Unmarshal(data, &valCurs); err != nil {
		return nil, fmt.Errorf("unmarshal xml: %w", err)
	}

	result := make([]CurrencyJSON, 0, len(valCurs.Currencies))

	for _, currency := range valCurs.Currencies {
		valStr := strings.Replace(currency.Value, ",", ".", -1)
		numStr := strings.TrimSpace(valStr)

		if numStr == "" {
			continue
		}

		var value float64
		if _, err := fmt.Sscanf(numStr, "%f", &value); err != nil {
			continue
		}

		result = append(result, CurrencyJSON{
			Code:  currency.CharCode,
			Value: value,
		})
	}

	return result, nil
}

func saveJSON(path string, data []CurrencyJSON) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	const dirPerm = 0o755

	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("create dir: %w", err)
	}

	if err := os.WriteFile(path, jsonData, 0o644); err != nil {
		return fmt.Errorf("write json: %w", err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	if *configPath == "" {
		fmt.Println("Config path is required")
		os.Exit(1)
	}

	cfg, err := loadConfig(*configPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	data, err := parseXML(cfg.InputPath)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		os.Exit(1)
	}

	if err := saveJSON(cfg.OutputPath, data); err != nil {
		fmt.Println("Error saving JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Conversion completed successfully.")
}
