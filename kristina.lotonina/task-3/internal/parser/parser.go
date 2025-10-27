package parser

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"

	"github.com/kef1rch1k/task-3/internal/models"
	"golang.org/x/net/html/charset"
)

func ParseAndSortXML(path string) ([]models.Valute, error) {
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

	sortedValutes := sortValutesByValueDesc(valCurs.Valutes)

	return sortedValutes, nil
}

func sortValutesByValueDesc(valutes []models.Valute) []models.Valute {
	sorted := make([]models.Valute, len(valutes))
	copy(sorted, valutes)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	return sorted
}
