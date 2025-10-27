package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/JingolBong/task-3/internal/valuteinfo"
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

	data, err := io.ReadAll(xmlFile)

	if err != nil {
		panic(err)
	}
	if err := xml.Unmarshal(data, &valuteCurs); err != nil {
		panic(err)
	}

	return valuteCurs, nil
}
