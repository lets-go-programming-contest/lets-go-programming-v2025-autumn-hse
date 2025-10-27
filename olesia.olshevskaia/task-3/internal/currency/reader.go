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

	defer func() {
		if err := file.Close(); err != nil {
			panic("Error closing XML file: " + err.Error())
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var xmlData struct {
		Currencies []model.Currency `xml:"Valute"`
	}
	if err := decoder.Decode(&xmlData); err != nil {
		panic("Error parsing XML: " + err.Error())
	}

	for i := range xmlData.Currencies {
		cur := &xmlData.Currencies[i]

		if n, err := strconv.Atoi(strings.TrimSpace(cur.CodeChar)); err == nil {
			cur.CodeNum = n
		}

		str := strings.ReplaceAll(strings.TrimSpace(cur.CodeChar), ",", ".")
		if v, err := strconv.ParseFloat(str, 64); err == nil {
			cur.RateValue = v
			cur.HasValue = true
		}
	}

	sort.Slice(xmlData.Currencies, func(i, j int) bool {
		return xmlData.Currencies[i].RateValue > xmlData.Currencies[j].RateValue
	})

	return xmlData.Currencies
}
