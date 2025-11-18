package xmlparser

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/VlasfimosY/task-3/internal/models"
)

func DecodeXML(filePath string) ([]models.CurrencyJSON, error) {
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
	decoder.CharsetReader = models.GetCharsetReader

	var valCurs models.ValCurs

	if err := decoder.Decode(&valCurs); err != nil {
		return nil, fmt.Errorf("cannot unmarshal XML: %w", err)
	}

	result := make([]models.CurrencyJSON, 0, len(valCurs.Currencies))

	for _, currency := range valCurs.Currencies {
		value, err := strconv.ParseFloat(strings.ReplaceAll(currency.Value, ",", "."), 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse value '%s': %w", currency.Value, err)
		}

		numStr := strings.TrimSpace(currency.NumCode)
		numCode := 0

		if numStr != "" {
			if n, err := strconv.Atoi(numStr); err == nil {
				numCode = n
			}
		}

		result = append(result, models.CurrencyJSON{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(currency.CharCode),
			Value:    value,
		})
	}

	return result, nil
}
