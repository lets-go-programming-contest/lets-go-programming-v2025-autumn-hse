package xmlutil

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"

	"golang.org/x/net/html/charset"
)

var ErrEmpty = errors.New("XML contains no elements")

func ParseElements[T any](xmlData []byte, rootName, elementName string) ([]T, error) {
	xmlDecoder := xml.NewDecoder(bytes.NewReader(xmlData))
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	var temp struct {
		XMLName  xml.Name `xml:""`
		Elements []T      `xml:",any"`
	}

	if err := xmlDecoder.Decode(&temp); err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	if temp.XMLName.Local != rootName {
		return nil, fmt.Errorf("expected root element %q, got %q", rootName, temp.XMLName.Local)
	}

	if len(temp.Elements) == 0 {
		return nil, ErrEmpty
	}

	return temp.Elements, nil
}
