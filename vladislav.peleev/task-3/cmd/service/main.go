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
	InputFile  string `yaml:"input"`
	OutputFile string `yaml:"output"`
}

type ValCurs struct {
	Currencies []Valute `xml:"Valute"`
}

type Valute struct {
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type CurrencyJSON struct {
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

var (
	errFileNotFound = errors.New("no such file or directory")
	errParseXML     = errors.New("error parsing XML")
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	config, err := readConfig(*configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := parseXML(config.InputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := writeJSON(config.OutputFile, data); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readConfig(configPath string) (*Config, error) {
	f, err := os.Open(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("open %s: %w", configPath, errFileNotFound)
		}
		return nil, fmt.Errorf("open %s: %w", configPath, err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}

func parseXML(filePath string) ([]CurrencyJSON, error) {
	cleanPath := filepath.Clean(filePath)
	if strings.TrimSpace(cleanPath) == "" || cleanPath == "." {
		return nil, fmt.Errorf("open %s: %w", cleanPath, errFileNotFound)
	}

	file, err := os.Open(cleanPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("open %s: %w", cleanPath, errFileNotFound)
		}
		return nil, fmt.Errorf("open %s: %w", cleanPath, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read xml: %w", err)
	}

	var valCurs ValCurs
	if err := xml.Unmarshal(data, &valCurs); err != nil {
		return nil, fmt.Errorf("%w: %v", errParseXML, err)
	}

	result := make([]CurrencyJSON, 0, len(valCurs.Currencies))

	for _, currency := range valCurs.Currencies {
		valStr := strings.TrimSpace(strings.ReplaceAll(currency.Value, ",", "."))
		if valStr == "" {
			continue
		}

		var num float64
		if _, err := fmt.Sscanf(valStr, "%f", &num); err != nil {
			continue
		}

		result = append(result, CurrencyJSON{
			CharCode: currency.CharCode,
			Value:    num,
		})
	}

	return result, nil
}

func writeJSON(filePath string, data []CurrencyJSON) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("create %s: %w", filePath, err)
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	if err := enc.Encode(data); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}
