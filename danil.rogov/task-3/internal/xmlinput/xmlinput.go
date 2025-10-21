package xmlinput

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/Tapochek2894/task-3/internal/valute"
	"golang.org/x/net/html/charset"
)

func ReadXML(path string) (valute.Valutes, error) {
	var valutes valute.Valutes

	data, err := os.ReadFile(path)
	if err != nil {
		return valute.Valutes{}, fmt.Errorf("error reading %s file: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&valutes)
	if err != nil {
		return valute.Valutes{}, fmt.Errorf("error decoding %s: %w", path, err)
	}

	return valutes, nil
}
