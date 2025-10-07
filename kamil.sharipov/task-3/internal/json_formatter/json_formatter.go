package json

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	xml "github.com/kamilSharipov/task-3/internal/xml_parser"
)

type ValuteJSON struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func parseValue(s string) (float64, error) {
	s = strings.Replace(s, ",", ".", 1)
	return strconv.ParseFloat(s, 64)
}

func FormateJSON(valCurs *xml.ValCurs) ([]byte, error) {

	valutes := make([]ValuteJSON, len(valCurs.Valutes))
	for i, valute := range valCurs.Valutes {
		numCode, err := strconv.Atoi(valute.NumCode)
		if err != nil {
			return nil, fmt.Errorf("invalid numCode %q: %w", valute.NumCode, err)
		}

		value, err := parseValue(valute.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid Value %q: %w", valute.Value, err)
		}

		valutes[i].CharCode = valute.CharCode
		valutes[i].NumCode = numCode
		valutes[i].Value = value
	}

	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})

	bytes, err := json.Marshal(valutes)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal: %w", err)
	}
	return bytes, nil
}
