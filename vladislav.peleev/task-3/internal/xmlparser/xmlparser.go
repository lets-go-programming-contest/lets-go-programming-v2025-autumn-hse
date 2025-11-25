package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/VlasfimosY/task-3/internal/models"
	"golang.org/x/text/encoding/charmap"
)

type xmlCurrency struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func GetCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(charset) {
	case "windows-1251", "cp1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	default:
		return input, nil
	}
}

func DecodeXML(filePath string) ([]models.Currency, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open XML file: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "warning: failed to close file %s: %v\n", filePath, cerr)
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = GetCharsetReader

	var valCurs struct {
		XMLName    xml.Name      `xml:"ValCurs"`
		Currencies []xmlCurrency `xml:"Valute"`
	}

	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("cannot unmarshal XML: %w", err)
	}

	result := make([]models.Currency, 0, len(valCurs.Currencies))

	for _, currency := range valCurs.Currencies {
		value, err := strconv.ParseFloat(strings.ReplaceAll(currency.Value, ",", "."), 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse value '%s': %w", currency.Value, err)
		}

		numCode := 0

		if currency.NumCode != "" {
			if n, err := strconv.Atoi(strings.TrimSpace(currency.NumCode)); err == nil {
				numCode = n
			}
		}

		result = append(result, models.Currency{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(currency.CharCode),
			Value:    value,
		})
	}

	return result, nil
}
