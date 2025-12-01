package handlers

import (
	"context"
	"errors"
	"strings"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, "no decorator") {
				return errors.New("can't be decorated")
			}

			if !strings.HasPrefix(data, "decorated: ") {
				data = "decorated: " + data
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case output <- data:
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	counter := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				for _, output := range outputs {
					close(output)
				}
				return nil
			}

			outputIndex := counter % len(outputs)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case outputs[outputIndex] <- data:
			}

			counter++
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	merged := make(chan string, 100)

	for _, input := range inputs {
		go func(in chan string) {
			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-in:
					if !ok {
						return
					}
					if !strings.Contains(data, "no multiplexer") {
						select {
						case <-ctx.Done():
							return
						case merged <- data:
						}
					}
				}
			}
		}(input)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-merged:
			if !ok {
				close(output)
				return nil
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case output <- data:
			}
		}
	}
}