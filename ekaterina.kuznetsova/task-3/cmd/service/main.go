package main

import (
	"bytes"
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

type Value float64

type ValuteXML struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    Value  `xml:"Value"`
}

type ValuteCurs struct {
	Valutes []ValuteXML `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
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

func (v *Value) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return err
	}

	str = strings.Replace(str, ",", ".", -1)

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("error parse float  %w", err)
	}

	*v = Value(value)

	return nil
}

func ParseValuteCursXML(inputFile string) (ValuteCurs, error) {
	var valCurs ValuteCurs

	dataXML, err := os.ReadFile(inputFile)
	if err != nil {
		return valCurs, fmt.Errorf("error read input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(dataXML))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&valCurs)
	if err != nil {
		return valCurs, fmt.Errorf("error parsing XML: %w", err)
	}

	return valCurs, nil
}

func CreateValuteCursJSON(valCurs ValuteCurs) ([]byte, error) {
	var valutesOutput []Valute

	for _, valute := range valCurs.Valutes {
		numCode, err := strconv.Atoi(valute.NumCode)
		if err != nil {
			numCode = 0
		}
		valutesOutput = append(valutesOutput, Valute{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    float64(valute.Value),
		})
	}

	sort.Slice(valutesOutput, func(i, j int) bool {
		return valutesOutput[i].Value > valutesOutput[j].Value
	})

	outputJSON, err := json.MarshalIndent(valutesOutput, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	return outputJSON, nil
}

func WriteFileJSON(outputFile string, outputJSON []byte) error {
	err := os.MkdirAll(filepath.Dir(outputFile), 0755)
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
