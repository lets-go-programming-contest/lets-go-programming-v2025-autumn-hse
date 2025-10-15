package writerjson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/Ekaterina-101/task-3/internal/parsexml"
)

type Valute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func CreateValuteCursJSON(valCurs parsexml.ValuteCurs) ([]byte, error) {
	valutesOutput := make([]Valute, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		numCode, err := strconv.Atoi(valute.NumCode)
		if err != nil {
			numCode = 0
		}
		valutesOutput = append(valutesOutput, Valute{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    float64(valute.Value),
		})
	}

	sort.Slice(valutesOutput, func(i, j int) bool {
		return valutesOutput[i].Value > valutesOutput[j].Value
	})

	outputJSON, err := json.MarshalIndent(valutesOutput, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	return outputJSON, nil
}

func WriteFileJSON(outputFile string, outputJSON []byte) error {
	const dirMode = 0o755

	const fileMode = 0o600

	err := os.MkdirAll(filepath.Dir(outputFile), dirMode)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	err = os.WriteFile(outputFile, outputJSON, fileMode)
	if err != nil {
		return fmt.Errorf("error writing output file: %w", err)
	}

	return nil
}
