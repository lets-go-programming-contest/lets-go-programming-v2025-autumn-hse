package xmlreader

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Anfisa111/task-3/internal/valutes"
	"golang.org/x/text/encoding/charmap"
)

var errCharset = errors.New("unknown charset")

func readFileXML(filepath string, value any) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(file))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil

		default:
			return nil, errCharset
		}
	}

	err = decoder.Decode(value)
	if err != nil {
		return fmt.Errorf("error decoding file: %w", err)
	}

	return nil
}

func DecodeValuteXML(filepath string) ([]valutes.Valute, error) {
	var valCurs valutes.ValCurs

	err := readFileXML(filepath, &valCurs)
	if err != nil {
		return nil, fmt.Errorf("error reading XML file: %w", err)
	}

	return valCurs.Valutes, nil
}
