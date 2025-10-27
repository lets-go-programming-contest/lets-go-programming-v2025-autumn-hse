package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const dirPerm = 0o755

type Currency struct {
	NumCode string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	Currencies []Currency `xml:"Valute"`
}

type OutputCurrency struct {
	NumCode int     `json:"num_code"`
	CharCode string `json:"char_code"`
	Value   float64 `json:"value"`
}

func readXML(filePath string) ([]Currency, error) {
	file, err := os.Open(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("open %s: no such file or directory", filePath)
	}
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "failed to close file %s: %v\n", filePath, cerr)
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	var valCurs ValCurs
	if err := xml.Unmarshal(data, &valCurs); err != nil {
		return nil, fmt.Errorf("error decoding XML: %w", err)
	}

	return valCurs.Currencies, nil
}

func convertCurrencies(valCurs ValCurs) ([]OutputCurrency, error) {
	var output []OutputCurrency

	for _, currency := range valCurs.Currencies {
		valStr := strings.ReplaceAll(currency.Value, ",", ".")
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value %q: %w", currency.Value, err)
		}

		numStr := strings.TrimSpace(currency.NumCode)
		if numStr == "" {
			continue
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("invalid NumCode %q: %w", currency.NumCode, err)
		}

		output = append(output, OutputCurrency{
			NumCode:  num,
			CharCode: currency.CharCode,
			Value:    val,
		})
	}

	return output, nil
}

func writeJSON(currencies []OutputCurrency, outputPath string) error {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("cannot create directory: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "failed to close output file: %v\n", cerr)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("cannot encode JSON: %w", err)
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: service <input.xml> <output.json>")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	currencies, err := readXML(inputPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	valCurs := ValCurs{Currencies: currencies}
	output, err := convertCurrencies(valCurs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := writeJSON(output, outputPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Conversion completed successfully.")
}

