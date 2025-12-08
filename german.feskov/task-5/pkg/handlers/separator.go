package handlers

import (
	"context"
)

func SeparatorFunc(
	ctx context.Context,
	input chan string,
	outputs []chan string,
) error {
	ind := 0

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		select {
		case data, ok := <-input:
			if !ok {
				return nil
			}

			outputs[ind] <- data

			ind = (ind + 1) % len(outputs)
		default:
		}
	}
}
