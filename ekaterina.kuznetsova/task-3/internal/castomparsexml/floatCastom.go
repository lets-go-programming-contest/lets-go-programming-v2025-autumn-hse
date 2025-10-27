package castomparsexml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ValueFloat float64

func (v *ValueFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := d.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("error decode xml: %w", err)
	}

	str = strings.ReplaceAll(str, ",", ".")

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return fmt.Errorf("error parse float  %w", err)
	}

	*v = ValueFloat(value)

	return nil
}
