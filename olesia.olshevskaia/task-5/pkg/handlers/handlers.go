package handlers

import (
	"context"
	"errors"
	"strings"
)

var errCantBeDecorated = errors.New("can't be decorated")

const (
	skipDecorator   = "no decorator"
	decoratedPrefix = "decorated: "
	skipMultiplexer = "no multiplexer"
)

func PrefixDecoratorFunc(ctx context.Context, input, output chan string) error {
	for {
		select {
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, skipDecorator) {
				return errCantBeDecorated
			}

			if !strings.HasPrefix(data, decoratedPrefix) {
				data = decoratedPrefix + data
			}

			select {
			case output <- data:
			case <-ctx.Done():
				return nil
			}

		case <-ctx.Done():
			for range input {
			}
			return nil
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	idn := 0

	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}
			outputs[idn] <- data
			idn = (idn + 1) % len(outputs)
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	for {
		for _, ch := range inputs {
			select {
			case <-ctx.Done():
				return nil
			case value, ok := <-ch:
				if !ok {
					continue
				}

				if strings.Contains(value, skipMultiplexer) {
					continue
				}
				output <- value
			default:
			}
		}
	}
}
