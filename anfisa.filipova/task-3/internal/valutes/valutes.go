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

const filePermission = 0o755

type FloatValue struct {
	Value float64
}

func (value *FloatValue) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var strValue string
	err := decoder.DecodeElement(&strValue, &start)
	if err != nil {
		return fmt.Errorf("error decoding string: %w", err)
	}

	strValue = strings.ReplaceAll(strValue, ",", ".")

	floatValue, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return fmt.Errorf("error converting string to float: %w", err)
	}

	value.Value = floatValue
	return nil
}

type Valute struct {
	NumCode  int        `xml:"NumCode" json:"num_code"`
	CharCode string     `xml:"CharCode" json:"char_code"`
	Value    FloatValue `xml:"Value" json:"value"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

var errCharset = errors.New("unknown charset")

func readFileXML(filepath string, value any) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(file))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil

		default:
			return nil, errCharset
		}
	}

	err = decoder.Decode(value)
	if err != nil {
		return fmt.Errorf("error decoding file: %w", err)
	}

	return nil
}

func DecodeValuteXML(filepath string) ([]Valute, error) {
	var valCurs ValCurs
	err := readFileXML(filepath, &valCurs)
	if err != nil {
		return nil, fmt.Errorf("error reading XML file: %w", err)
	}

	return valCurs.Valutes, nil
}

func SortCurrencies(currencies []Valute) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value.Value > currencies[j].Value.Value
	})
}

func WriteCurrenciesJSON(currencies []Valute, filePath string) error {
	dir := filepath.Dir(filePath)

	err := os.MkdirAll(dir, filePermission)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
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
		return fmt.Errorf("error encoding file: %w", err)
	}

	return nil
}
