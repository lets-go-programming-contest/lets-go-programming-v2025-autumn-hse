package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kuzid-17/task-3/internal/xmlparser"
)

type OutputValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func WriteJSON(filename string, valutes []xmlparser.Valute) {
	dir := filepath.Dir(filename)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic("failed to create directory")
	}

	output := make([]OutputValute, 0, len(valutes))

	for _, valute := range valutes {
		numCode, err := strconv.Atoi(valute.NumCode)
		if err != nil {
			continue
		}

		output = append(output, OutputValute{
			NumCode:  numCode,
			CharCode: valute.CharCode,
			Value:    float64(valute.Value),
		})
	}

	file, err := os.Create(filename)
	if err != nil {
		panic("failed to create output file")
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("error closing file", err)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(output)
	if err != nil {
		panic("failed to encode JSON")
	}
}
