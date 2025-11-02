package xmlreader

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

var errCharset = errors.New("unknown charset")

func ReadFileXML(filepath string, value any) error {
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
