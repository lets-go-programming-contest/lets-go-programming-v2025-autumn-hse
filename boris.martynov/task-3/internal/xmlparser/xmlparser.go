package xmlparser

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/JingolBong/task-3/internal/valuteinfo"
	"golang.org/x/net/html/charset"
)

func Xmlparser(inputxml string) (valuteinfo.ValuteCurs, error) {
	var valuteCurs valuteinfo.ValuteCurs

	xmlFile, err := os.Open(inputxml)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := xmlFile.Close(); err != nil {
			panic(fmt.Errorf("failed to close file: %w", err))
		}
	}()

	decoder := xml.NewDecoder(xmlFile)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&valuteCurs); err != nil {

		return valuteinfo.ValuteCurs{}, fmt.Errorf("while decoding: %w", err)
	}

	return valuteCurs, nil
}
