package parser

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"

	"github.com/kef1rch1k/task-3/internal/models"
)


func ParseAndSortXML(path string) ([]models.OutputValute, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(fmt.Sprintf("Failed to close file: %v", err))
		}
	}()

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
