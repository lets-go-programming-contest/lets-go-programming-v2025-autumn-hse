package handlers

import "context"

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
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

	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}
			select {
			case outputs[index] <- data:
			case <-ctx.Done():
				return nil
			}

			index = (index + 1) % len(outputs)
		}
	}
}
