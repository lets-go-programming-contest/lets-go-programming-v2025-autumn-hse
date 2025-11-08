package xml

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/kamilSharipov/task-3/internal/model"
	"golang.org/x/net/html/charset"
)

var errNoValutes = errors.New("XML contains no valutes")

func ParseXML(xmlData []byte) ([]model.Valute, error) {
	xmlDecoder := xml.NewDecoder(bytes.NewReader(xmlData))
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	var temp struct {
		Valutes []model.Valute `xml:"Valute"`
	}

	if err := xmlDecoder.Decode(&temp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML: %w", err)
	}

	if len(temp.Valutes) == 0 {
		return nil, errNoValutes
	}

	return temp.Valutes, nil
}
