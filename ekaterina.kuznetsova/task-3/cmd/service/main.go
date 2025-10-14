package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
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

type Valute struct {
	XMLName  xml.Name `xml:"Valute"`
	NumCode  string   `xml:"NumCode"`
	CharCode string   `xml:"CharCode"`
	Value    string   `xml:"Value"`
}

type ValuteCurs struct {
	Valutes []Valute `xml:"Valutes"`
}

type ValuteJSON struct {
	NumCode  int     `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Value    float64 `json:"Value"`
}

type ValuteCursJSON struct {
	Valutes []ValuteJSON `json:"Valutes"`
}

func LoadConfigYaml() (Config, error) {
	configPath := flag.String("config", "", "path of config file")
	flag.Parse()

	dataConfig, err := os.ReadFile(*configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error read config: %w", err)
	}

	var config Config

	err = yaml.Unmarshal(dataConfig, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error parsing yaml: %w", err)
	}

	return config, nil
}

func ParseValuteCursXML(inputFile string) (ValuteCurs, error) {
	var valCurs ValuteCurs

	file, err := os.Open(inputFile)
	if err != nil {
		return valCurs, fmt.Errorf("open input file: %w", err)
	}

	defer func() {
		file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&valCurs)
	if err != nil {
		return valCurs, fmt.Errorf("error parsing XML: %w", err)
	}

	return valCurs, nil
}

func CreateValuteCursJSON(valCurs ValuteCurs) ([]byte, error) {
	var valutesOutput []ValuteJSON

	for _, valute := range valCurs.Valutes {
		numCode, _ := strconv.Atoi(valute.NumCode)
		value, _ := strconv.ParseFloat(strings.ReplaceAll(valute.Value, ",", "."), 64)
		valutesOutput = append(valutesOutput, ValuteJSON{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	sort.Slice(valutesOutput, func(i, j int) bool {
		return valutesOutput[i].Value > valutesOutput[j].Value
	})

	outputJSON, err := json.MarshalIndent(valutesOutput, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	return outputJSON, nil
}

func WriteFileJSON(outputFile string, outputJSON []byte) error {
	err := os.MkdirAll(filepath.Dir(outputFile), 0766)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	err = os.WriteFile(outputFile, outputJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}

func main() {
	config, err := LoadConfigYaml()
	if err != nil {
		panic(err)
	}

	valCurs, err := ParseValuteCursXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	outputJSON, err := CreateValuteCursJSON(valCurs)
	if err != nil {
		panic(err)
	}

	err = WriteFileJSON(config.OutputFile, outputJSON)
	if err != nil {
		panic(err)
	}
}
