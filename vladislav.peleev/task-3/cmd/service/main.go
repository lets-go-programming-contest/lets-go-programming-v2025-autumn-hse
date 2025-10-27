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
	NumCode  string  `json:"num_code"`
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

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read XML file: %w", err)
	}

	utf8Data := convertWindows1251ToUTF8(data)

	var valCurs ValCurs
	err = xml.Unmarshal(utf8Data, &valCurs)
	if err != nil {
		err2 := xml.Unmarshal(data, &valCurs)
		if err2 != nil {
			return nil, fmt.Errorf("cannot unmarshal XML: %w", err)
		}
	}

	result := make([]CurrencyJSON, len(valCurs.Currencies))
	for i, currency := range valCurs.Currencies {
		valStr := strings.Replace(currency.Value, ",", ".", -1)
		value, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse value '%s': %w", currency.Value, err)
		}

		result[i] = CurrencyJSON{
			NumCode:  strings.TrimSpace(currency.NumCode),
			CharCode: strings.TrimSpace(currency.CharCode),
			Value:    value,
		}
	}

	return result, nil
}

func convertWindows1251ToUTF8(data []byte) []byte {
	windows1251ToUTF8 := map[byte]string{
		0xC0: "А", 0xC1: "Б", 0xC2: "В", 0xC3: "Г", 0xC4: "Д", 0xC5: "Е", 0xC6: "Ж", 0xC7: "З",
		0xC8: "И", 0xC9: "Й", 0xCA: "К", 0xCB: "Л", 0xCC: "М", 0xCD: "Н", 0xCE: "О", 0xCF: "П",
		0xD0: "Р", 0xD1: "С", 0xD2: "Т", 0xD3: "У", 0xD4: "Ф", 0xD5: "Х", 0xD6: "Ц", 0xD7: "Ч",
		0xD8: "Ш", 0xD9: "Щ", 0xDA: "Ъ", 0xDB: "Ы", 0xDC: "Ь", 0xDD: "Э", 0xDE: "Ю", 0xDF: "Я",
		0xE0: "а", 0xE1: "б", 0xE2: "в", 0xE3: "г", 0xE4: "д", 0xE5: "е", 0xE6: "ж", 0xE7: "з",
		0xE8: "и", 0xE9: "й", 0xEA: "к", 0xEB: "л", 0xEC: "м", 0xED: "н", 0xEE: "о", 0xEF: "п",
		0xF0: "р", 0xF1: "с", 0xF2: "т", 0xF3: "у", 0xF4: "ф", 0xF5: "х", 0xF6: "ц", 0xF7: "ч",
		0xF8: "ш", 0xF9: "щ", 0xFA: "ъ", 0xFB: "ы", 0xFC: "ь", 0xFD: "э", 0xFE: "ю", 0xFF: "я",
	}

	var result strings.Builder
	for _, b := range data {
		if utf8Char, exists := windows1251ToUTF8[b]; exists {
			result.WriteString(utf8Char)
		} else {
			result.WriteByte(b)
		}
	}

	return []byte(result.String())
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
