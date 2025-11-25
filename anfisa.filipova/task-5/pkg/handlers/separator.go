package handlers

import "context"

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	counter := 0
	for {
		select {
		case data, ok := <-input:
			if !ok {
				return nil
			}
			index := counter % len(outputs)
			select {
			case outputs[index] <- data:
				counter++
			case <-ctx.Done():
				return ctx.Err()
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
