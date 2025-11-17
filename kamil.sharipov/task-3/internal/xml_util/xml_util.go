package xmlutil

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

var (
	ErrEmpty     = errors.New("XML contains no elements")
	ErrWrongRoot = errors.New("XML root element name mismatch")
)

func ReadInput[T any](path string) (*T, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("failed to stat file %q: %w", path, err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %q: %w", path, err)
	}

	xmlDecoder := xml.NewDecoder(bytes.NewReader(data))
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	var result T
	if err := xmlDecoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	return &result, nil
}
