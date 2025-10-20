package jsonwriter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kuzid-17/task-3/internal/xmlparser"
)

const mode = 0o755

func WriteJSON(filename string, valutes []xmlparser.OutputValute) {
	dir := filepath.Dir(filename)

	err := os.MkdirAll(dir, mode)
	if err != nil {
		panic("failed to create directory")
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
		panic("failed to create output file")
	}

	defer func() {
		err = file.Close()
		if err != nil {
			panic(err.Error() + "111")
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(output)
	if err != nil {
		fmt.Println("failed to encode JSON")
	}
}
