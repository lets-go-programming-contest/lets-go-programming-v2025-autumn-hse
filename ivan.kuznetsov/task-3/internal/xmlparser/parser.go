package xmlparser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/kuzid-17/task-3/internal/models"
	"golang.org/x/net/html/charset"
)

func ParseXML(filename string) (*models.ValCurs, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read input config file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs models.ValCurs

	err = decoder.Decode(&valCurs)
	if err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	return &valCurs, nil
}
