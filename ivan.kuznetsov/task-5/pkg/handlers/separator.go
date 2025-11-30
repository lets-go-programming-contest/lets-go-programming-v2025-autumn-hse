package handlers

import "context"

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	index := 0

	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			select {
			case <-ctx.Done():
				return nil
			case outputs[index] <- data:
				index = (index + 1) % len(outputs)
			}
		}
	}
}
