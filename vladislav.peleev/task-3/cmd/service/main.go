package main

import (
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input_file"`
	OutputFile string `yaml:"output_file"`
}

type CurrencyXML struct {
	CharCode string `xml:"CharCode"`
	Nominal  string `xml:"Nominal"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	Currencies []CurrencyXML `xml:"Currency"`
}

type CurrencyJSON struct {
	Code   string  `json:"code"`
	Amount int     `json:"amount"`
	Value  float64 `json:"value"`
}

func readConfig(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, errors.New("config path is empty")
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("open %s: no such file or directory", configPath)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("cannot parse config: %w", err)
	}

	return &config, nil
}

func parseXML(filePath string) ([]CurrencyJSON, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("open %s: no such file or directory", filePath)
	}

	var valCurs ValCurs
	if err := xml.Unmarshal(data, &valCurs); err != nil {
		return nil, fmt.Errorf("error parsing XML: %w", err)
	}

	result := make([]CurrencyJSON, 0, len(valCurs.Currencies))
	for _, currency := range valCurs.Currencies {
		amount := 1
		if currency.Nominal != "" {
			fmt.Sscanf(currency.Nominal, "%d", &amount)
		}

		var value float64
		if currency.Value != "" {
			fmt.Sscanf(currency.Value, "%f", &value)
		}

		result = append(result, CurrencyJSON{
			Code:   currency.CharCode,
			Amount: amount,
			Value:  value,
		})
	}

	return result, nil
}

func writeJSON(outputPath string, data []CurrencyJSON) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("cannot create file %s: %w", outputPath, err)
	}
	defer file.Close()

	for _, c := range data {
		_, err := fmt.Fprintf(file, "%s %d %.2f\n", c.Code, c.Amount, c.Value)
		if err != nil {
			return fmt.Errorf("cannot write to file: %w", err)
		}
	}

	return nil
}

func run() error {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	config, err := readConfig(*configPath)
	if err != nil {
		return err
	}

	data, err := parseXML(config.InputFile)
	if err != nil {
		return err
	}

	if err := writeJSON(config.OutputFile, data); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
