package handlers

import (
	"context"
)

func SeparatorFunc(
	ctx context.Context,
	input chan string,
	outputs []chan string,
) error {
	index := 0

	for {
		select {
		case <-ctx.Done():
			return nil
		case value, ok := <-input:
			if !ok {
				return nil
			}

			out := outputs[index%len(outputs)]
			out <- value

			index++
		}
	}
}
