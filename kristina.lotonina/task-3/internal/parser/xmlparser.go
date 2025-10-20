package parser

import (
	"github.com/kef1rch1k/task-3/internal/models"
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseAndSortXML(path string) ([]models.OutputValute, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var valCurs models.ValCurs
	err = xml.Unmarshal(data, &valCurs)
	if err != nil {
		return nil, err
	}

	for i, v := range valCurs.Valutes {
		val := strings.Replace(v.ValueStr, ",", ".", 1)
		floatVal, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid float in value: %s", v.ValueStr)
		}
		valCurs.Valutes[i].Value = floatVal
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	var output []models.OutputValute
	for _, v := range valCurs.Valutes {
		output = append(output, models.OutputValute{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    v.Value,
		})
	}

	return output, nil
}
