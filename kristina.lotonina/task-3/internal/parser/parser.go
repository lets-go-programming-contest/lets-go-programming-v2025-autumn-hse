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
		return nil, fmt.Errorf("opening XML file: %w", err)
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
		return nil, fmt.Errorf("decoding XML: %w", err)
	}

	for index, v := range valCurs.Valutes {
		val := strings.Replace(v.ValueStr, ",", ".", 1)

		floatVal, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err != nil {
			return nil, fmt.Errorf("parsing float: %w", err)
		}

		valCurs.Valutes[index].Value = floatVal
	}

	sorted := valCurs.Valutes
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	output := make([]models.OutputValute, 0, len(sorted))

	for _, v := range sorted {
		output = append(output, models.OutputValute{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    v.Value,
		})
	}

	return output, nil
}
