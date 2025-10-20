package parser

import (
	"encoding/xml"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/kef1rch1k/task-3/internal/models"

	"golang.org/x/net/html/charset"
)


func ParseAndSortXML(path string) ([]models.OutputValute, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs models.ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return nil, err
	}

	for i, v := range valCurs.Valutes {
		val := strings.Replace(v.ValueStr, ",", ".", 1)
		floatVal, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err != nil {
			return nil, err
		}
		valCurs.Valutes[i].Value = floatVal
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	var output []models.OutputValute
	for _, v := range valCurs.Valutes {
		output = append(output, models.OutputValute{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    v.Value,
		})
	}

	return output, nil
}
