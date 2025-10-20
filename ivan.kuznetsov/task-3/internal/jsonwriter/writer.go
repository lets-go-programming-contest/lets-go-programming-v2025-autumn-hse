package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kuzid-17/task-3/internal/xmlparser"
)

const mode = 0o755

func WriteJSON(filename string, valutes []xmlparser.OutputValute) int {
	dir := filepath.Dir(filename)

	err := os.MkdirAll(dir, mode)
	if err != nil {
		return -1
	}

	output := make([]xmlparser.OutputValute, 0, len(valutes))

	for _, valute := range valutes {
		output = append(output, xmlparser.OutputValute{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    xmlparser.CommaFloat(valute.Value),
		})
	}

	file, err := os.Create(filename)
	if err != nil {
		return -1
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err.Error())
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(output)
	if err != nil {
		fmt.Println("failed to encode JSON")
	}

	return 0
}
