package writerjson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Ekaterina-101/task-3/internal/parsexml"
)

const (
	dirMode  = 0o755
	fileMode = 0o600
)

func CreateValuteCursJSON(valCurs parsexml.ValuteCurs) ([]byte, error) {
	// valutesOutput := make([]parsexml.Valute, 0, len(valCurs.Valutes))

	// for _, valute := range valCurs.Valutes {
	// 	valutesOutput = append(valutesOutput, parsexml.Valute{
	// 		NumCode:  valute.NumCode,
	// 		CharCode: valute.CharCode,
	// 		Value:    parsexml.ValueFloat(valute.Value),
	// 	})
	// }

	outputJSON, err := json.MarshalIndent(valCurs.Valutes, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	return outputJSON, nil
}

func WriteFileJSON(outputFile string, outputJSON []byte) error {
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
