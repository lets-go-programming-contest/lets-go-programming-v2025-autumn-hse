package handlers

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrCantDecorate = errors.New("can't be decorated")
)

func PrefixDecoratorFunc(ctx context.Context, input, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil

		case value, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(value, "no decorator") {
				return ErrCantDecorate
			}

			if !strings.HasPrefix(value, "decorated: ") {
				value = "decorated: " + value
			}

			output <- value
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	index := 0

	for {
		select {
		case <-ctx.Done():
			return nil

		case value, ok := <-input:
			if !ok {
				return nil
			}

			outputs[index] <- value
			index = (index + 1) % len(outputs)
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	for {
		for _, channel := range inputs {
			select {
			case <-ctx.Done():
				return nil

			case value, ok := <-channel:
				if !ok {
					continue
				}

				if strings.Contains(value, "no multiplexer") {
					continue
				}

				output <- value

			}
		}
	}
}
