package xmlinput

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Tapochek2894/task-3/internal/valute"
	"golang.org/x/text/encoding/charmap"
)

func ReadXML(path string) ([]valute.ValuteInfo, error) {
	// Check if input file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch strings.ToLower(charset) {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return input, nil
		}
	}

	var valCurs valute.ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	result := make([]valute.ValuteInfo, 0, len(valCurs.Valutes))
	for _, v := range valCurs.Valutes {
		valueStr := strings.ReplaceAll(v.Value, ",", ".")
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("parse value %s: %w", v.Value, err)
		}
		numcode, _ := strconv.Atoi(v.NumCode)
		result = append(result, valute.ValuteInfo{
			NumCode:  numcode,
			CharCode: v.CharCode,
			Value:    value,
		})
	}

	return result, nil
}
