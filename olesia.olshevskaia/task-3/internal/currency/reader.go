package currency

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Olesia.Ol/task-3/internal/model"
	"golang.org/x/net/html/charset"
)

func Read(path string) ([]model.Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open XML file %q: %w", path, err)
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			panic(fmt.Sprintf("failed to close file: %v", cerr))
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var xmlData struct {
		Currencies []model.Currency `xml:"ValCurs>Valute"`
	}

	if err := decoder.Decode(&xmlData); err != nil {
		return nil, fmt.Errorf("failed to decode XML file %q: %w", path, err)
	}

	result := make([]model.Currency, 0, len(xmlData.Currencies))

	for _, val := range xmlData.Currencies {
		num := parseCodeNum(val.RawNum)
		value, hasValue := parseValue(val.RawValue)

		result = append(result, model.Currency{
			CodeNum:   num,
			CodeChar:  val.CodeChar,
			RateValue: value,
			HasValue:  hasValue,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].RateValue > result[j].RateValue
	})

	return result, nil
}

func parseCodeNum(raw string) int {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return 0
	}

	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0
	}

	return n
}

func parseValue(raw string) (float64, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return 0, false
	}

	raw = strings.ReplaceAll(raw, ",", ".")
	if value, err := strconv.ParseFloat(raw, 64); err != nil {
		return 0, false
	} else {
		return value, true
	}
}
