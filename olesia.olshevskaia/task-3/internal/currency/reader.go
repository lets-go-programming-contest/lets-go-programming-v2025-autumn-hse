package currency

import (
	"encoding/xml"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Olesia.Ol/task-3/internal/model"
	"golang.org/x/net/html/charset"
)

func Read(path string) []model.Currency {
	file, err := os.Open(path)
	if err != nil {
		panic("Cannot open XML file: " + err.Error())
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var xmlData struct {
		Currencies []model.Currency `xml:"Valute"`
	}
	if err := decoder.Decode(&xmlData); err != nil {
		panic("Error parsing XML: " + err.Error())
	}

	result := make([]model.Currency, 0, len(xmlData.Currencies))

	for _, val := range xmlData.Currencies {
		num := 0
		value := 0.0

		if strings.TrimSpace(val.RawNum) != "" {
			if n, err := strconv.Atoi(val.RawNum); err == nil {
				num = n
			}
		}

		hasValue := false
		if strings.TrimSpace(val.RawValue) != "" {
			str := strings.ReplaceAll(val.RawValue, ",", ".")
			if v, err := strconv.ParseFloat(str, 64); err == nil {
				value = v
				hasValue = true
			}
		}

		result = append(result, model.Currency{
			CodeNum:   num,
			CodeChar:  val.CodeChar,
			RateValue: value,
			HasValue:  hasValue,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].RateValue > result[j].RateValue
	})

	return result
}
