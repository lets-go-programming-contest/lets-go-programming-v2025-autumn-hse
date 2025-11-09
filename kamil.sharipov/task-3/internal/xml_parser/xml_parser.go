package xml

import (
	"fmt"

	"github.com/kamilSharipov/task-3/internal/model"
	xmlutil "github.com/kamilSharipov/task-3/internal/xml_util"
)

func ParseXML(xmlData []byte) ([]model.Valute, error) {
	valutes, err := xmlutil.ParseElements[model.Valute](xmlData, "ValCurs", "Valute")
	if err != nil {
		return nil, fmt.Errorf("failed to parse XML into Valute: %w", err)
	}

	return valutes, nil
}
