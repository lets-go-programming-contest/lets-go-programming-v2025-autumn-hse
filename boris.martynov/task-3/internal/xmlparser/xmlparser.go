package xmlparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"

	"github.com/JingolBong/task-3/internal/valuteinfo"
)

func Xmlparser(inputxml string) (valuteinfo.ValuteCurs, error) {
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

	data, err := io.ReadAll(xmlFile)
	if err != nil {

		return valuteinfo.ValuteCurs{}, fmt.Errorf("reading file: %w", err)
	}
	if bytes.Contains(data, []byte("windows-1251")) {
		decoder := charmap.Windows1251.NewDecoder()
		data, err = decoder.Bytes(data)
		if err != nil {

			return valuteinfo.ValuteCurs{}, fmt.Errorf("cannot decode windows-1251: %w", err)
		}
	}

	dataStr := strings.ReplaceAll(string(data), `encoding="windows-1251"`, `encoding="utf-8"`)

	data = []byte(dataStr)
	if err := xml.Unmarshal(data, &valuteCur); err != nil {

		return valuteinfo.ValuteCurs{}, fmt.Errorf("Unmarshal xml: %w", err)
	}

	return valuteCur, nil
}
