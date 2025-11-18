package parser

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"github.com/yourusername/projectname/model"
)

func (v *model.CurrencyValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("error decode xml: %w", err)
	}

	str = strings.ReplaceAll(str, ",", ".")

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("error parse float: %w", err)
	}

	*v = model.CurrencyValue(value)

	return nil
}
