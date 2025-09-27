package xmlin

import (
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

const errMsg = "while read XML %q: %w"

func Read(path string, val any) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf(errMsg, path, err)
	}

	defer func() {
		if ferr := file.Close(); ferr != nil && err == nil {
			err = fmt.Errorf(errMsg, path, ferr)
		}
	}()

	dec := xml.NewDecoder(file)
	dec.CharsetReader = charset.NewReaderLabel

	if err := dec.Decode(val); err != nil {
		return fmt.Errorf(errMsg, path, err)
	}

	return nil
}
