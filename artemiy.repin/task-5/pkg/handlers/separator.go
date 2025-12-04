package handlers

import (
	"context"
)

func SeparatorFunc(
	ctx context.Context,
	input chan string,
	outputs []chan string,
) error {
	if len(outputs) == 0 {
		for {
			select {
			case <-ctx.Done():
				return nil
			case _, ok := <-input:
				if !ok {
					return nil
				}
			}
		}
	}

	index := 0
	total := len(outputs)

	for {
		select {
		case <-ctx.Done():
			return nil

		case data, ok := <-input:
			if !ok {
				return nil
			}

			out := outputs[index]

			index++
			if index >= total {
				index = 0
			}

			select {
			case <-ctx.Done():
				return nil
			case out <- data:
			}
		}
	}
}
