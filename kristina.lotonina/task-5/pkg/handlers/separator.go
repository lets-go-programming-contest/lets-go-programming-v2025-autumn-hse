package handlers

import (
	"context"
)

func SeparatorFunc(
	ctx context.Context,
	input chan string,
	outputs []chan string,
) error {
	i := 0
	for {
		select {
		case <-ctx.Done():
			return nil
		case v, ok := <-input:
			if !ok {
				return nil
			}

			out := outputs[i%len(outputs)]
			out <- v
			i++
		}
	}
}
