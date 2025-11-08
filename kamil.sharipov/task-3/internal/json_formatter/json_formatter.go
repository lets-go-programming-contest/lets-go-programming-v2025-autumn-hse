package json

import (
	"encoding/json"
	"fmt"

	"github.com/kamilSharipov/task-3/internal/model"
)

func FormateJSON(valutes []model.Valute) ([]byte, error) {
	data, err := json.Marshal(valutes)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal valutes to JSON: %w", err)
	}

	return data, nil
}
