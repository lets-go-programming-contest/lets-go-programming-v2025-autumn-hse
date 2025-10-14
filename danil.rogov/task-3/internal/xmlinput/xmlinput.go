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
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unknown charset: %s", charset)
	}

	var valCurs valute.ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	result := make([]valute.ValuteInfo, 0, len(valCurs.Valutes))
	for _, v := range valCurs.Valutes {
		value, _ := strconv.ParseFloat(strings.ReplaceAll(v.Value, ",", "."), 64)

		result = append(result, valute.ValuteInfo{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    value,
		})
	}

	return result, nil
}
