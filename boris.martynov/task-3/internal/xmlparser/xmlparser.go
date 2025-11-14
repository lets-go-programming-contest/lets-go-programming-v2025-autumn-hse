package xmlparser

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/JingolBong/task-3/internal/valuteinfo"
	"golang.org/x/net/html/charset"
)

func XMLParse(inputxml string) (valuteinfo.ValuteCurs, error) {
	var valuteCur valuteinfo.ValuteCurs

	xmlFile, err := os.Open(inputxml)
	if err != nil {
		return valuteinfo.ValuteCurs{}, fmt.Errorf("opening file: %w", err)
	}

	defer func() {
		if err := xmlFile.Close(); err != nil {
			panic(fmt.Errorf("failed to close file: %w", err))
		}
	}()

	reader, err := charset.NewReader(xmlFile, "")
	if err != nil {
		return valuteinfo.ValuteCurs{}, fmt.Errorf("charset reader: %w", err)
	}

	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&valuteCur); err != nil {
		return valuteinfo.ValuteCurs{}, fmt.Errorf("unmarshal xml: %w", err)
	}
	return valuteCur, nil
}
