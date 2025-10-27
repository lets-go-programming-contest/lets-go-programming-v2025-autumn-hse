package currency

import (
	"encoding/xml"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"

	"github.com/Olesia.Ol/task-3/internal/model"
)

func Read(path string) []model.Currency {
	file, err := os.Open(path)
	if err != nil {
		panic("Can not open XML file: " + path + " - " + err.Error())
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var xmlData model.CurrenciesXML
	if err := decoder.Decode(&xmlData); err != nil {
		panic("Error parsing XML: " + err.Error())
	}

	result := make([]model.Currency, 0, len(xmlData.Currencies))

	for _, val := range xmlData.Currencies {
		num := 0
		value := 0.0

		if n, err := strconv.Atoi(strings.TrimSpace(val.CodeNum)); err == nil {
			num = n
		}

		if strings.TrimSpace(val.RateValue) != "" {
			str := strings.ReplaceAll(val.RateValue, ",", ".")
			if v, err := strconv.ParseFloat(str, 64); err == nil {
				value = v
			}
		}

		result = append(result, model.Currency{
			CodeNum:   num,
			CodeChar:  val.CodeChar,
			RateValue: value,
			HasValue:  strings.TrimSpace(val.RateValue) != "",
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].RateValue > result[j].RateValue
	})

	return result
}
