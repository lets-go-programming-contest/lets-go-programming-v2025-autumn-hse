package xmlin

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/6ermvH/german.feskov/task-3/internal/valute"
	"golang.org/x/net/html/charset"
)

func Read(path string) (valute.ValCursXML, error) {
	var valCurs valute.ValCursXML

	file, err := os.Open(path)
	if err != nil {
		return valCurs, fmt.Errorf("open %q: %w", path, err)
	}

	defer func() {
		if ferr := file.Close(); ferr != nil {
			panic(fmt.Errorf("close file %q: %w", path, err))
		}
	}()

	dec := xml.NewDecoder(file)
	dec.CharsetReader = charset.NewReaderLabel

	if err := dec.Decode(&valCurs); err != nil {
		return valCurs, fmt.Errorf("decode to %q: %w", path, err)
	}

	return valCurs, nil
}
