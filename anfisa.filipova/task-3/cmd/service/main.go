package main

import (
	"bytes"
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

	"github.com/Anfisa111/task-3/internal/config"
	"golang.org/x/text/encoding/charmap"
)

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	ValueStr string `xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type ValuteOutput struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func DecodeValuteXML(filepath string) []ValuteOutput {
	file, err := os.ReadFile(filepath) //OpenFile(filepath, os.O_RDONLY, 0777)
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %v", err))
	}

	var valCurs ValCurs

	decoder := xml.NewDecoder(bytes.NewReader(file))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil

		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}
	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(fmt.Sprintf("Error decoding file: %v", err))
	}

	var currencies []ValuteOutput
	for _, valute := range valCurs.Valutes {
		valueStr := strings.Replace(valute.ValueStr, ",", ".", -1)
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			panic(fmt.Sprintf("Error type converting: %v", err))
		}
		currencies = append(currencies, ValuteOutput{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    value,
		})
	}

	return currencies
}

func SortCurrencies(currencies []ValuteOutput) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

func WriteJSON(currencies []ValuteOutput, filePath string) {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(fmt.Sprintf("Error creating directory: %v", err))
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error creating file: %v", err))
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(currencies)
	if err != nil {
		panic(fmt.Sprintf("Error encoding file: %v", err))
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	config := config.LoadConfig(*configPath)
	if _, err := os.Stat(config.InputFile); os.IsNotExist(err) {
		panic(fmt.Sprintf("Inputfile is not existing: %v", err))
	}
	currencies := DecodeValuteXML(config.InputFile)
	SortCurrencies(currencies)
	WriteJSON(currencies, config.OutputFile)
}
