package valutes

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const (
	maxValuteCount = 70
	filePermission = 0o755
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

func readValCurs(filepath string) ValCurs {
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

	return valCurs
}

func DecodeValuteXML(filepath string) []ValuteOutput {
	valCurs := readValCurs(filepath)
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

func WriteCurrenciesJSON(currencies []ValuteOutput, filePath string) {
	dir := filepath.Dir(filePath)

	err := os.MkdirAll(dir, filePermission)
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
