package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/kamilSharipov/task-3/internal/config"
	json "github.com/kamilSharipov/task-3/internal/json_formatter"
	"github.com/kamilSharipov/task-3/internal/model"
	xml "github.com/kamilSharipov/task-3/internal/xml_parser"
)

const (
	dirPermissions = 0o755
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	configPath := parseFlags()
	cfg, err := loadConfig(configPath)
	panicOnErr(err)

	xmlData, err := readXMLFile(cfg.InputFile)
	panicOnErr(err)

	valCurs, err := parseXMLData(xmlData)
	panicOnErr(err)

	sortValCurs(valCurs)

	jsonBytes, err := formatJSON(valCurs)
	panicOnErr(err)

	err = writeOutput(cfg.OutputFile, jsonBytes)
	panicOnErr(err)
}

func parseFlags() string {
	configPath := flag.String("config", "", "config file path")
	flag.Parse()

	if *configPath == "" {
		panic("missing required flag: -config")
	}

	return *configPath
}

func loadConfig(path string) (config.Config, error) {
	cfg, err := config.LoadConfig(path)
	if err != nil {
		return config.Config{}, fmt.Errorf("failed to load config from %q: %w", path, err)
	}

	return cfg, nil
}

func readXMLFile(path string) ([]byte, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("failed to stat file %q: %w", path, err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %q: %w", path, err)
	}

	return data, nil
}

func parseXMLData(data []byte) ([]model.Valute, error) {
	valCurs, err := xml.ParseXML(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	return valCurs, nil
}

func sortValCurs(valCurs []model.Valute) {
	sort.Slice(valCurs, func(i, j int) bool {
		return float64(valCurs[i].Value) > float64(valCurs[j].Value)
	})
}

func formatJSON(valCurs []model.Valute) ([]byte, error) {
	bytes, err := json.FormateJSON(valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to format JSON: %w", err)
	}

	return bytes, nil
}

func writeOutput(outputPath string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(outputPath), dirPermissions)
	if err != nil {
		return fmt.Errorf("failed to create output directory for %q: %w", outputPath, err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file %q: %w", outputPath, err)
	}

	defer func() {
		_ = outputFile.Close()
	}()

	_, err = outputFile.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write to %q: %w", outputPath, err)
	}

	return nil
}
