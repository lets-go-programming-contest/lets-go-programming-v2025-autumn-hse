package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
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

var errCharset = errors.New("unknown charset")

const maxValuteCount = 70

func DecodeValuteXML(filepath string) []ValuteOutput {
	file, err := os.ReadFile(filepath)
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
			return nil, errCharset
		}
	}

	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(fmt.Sprintf("Error decoding file: %v", err))
	}

	currencies := make([]ValuteOutput, 0, maxValuteCount)

	for _, valute := range valCurs.Valutes {
		valueStr := strings.ReplaceAll(valute.ValueStr, ",", ".")

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
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		panic(fmt.Sprintf("Error creating directory: %v", err))
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error creating file: %v", err))
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("Error close file: %v", err))
		}
	}()

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
			fmt.Fprintf(os.Stderr, "Recovered from panic: %v\n", r)
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
