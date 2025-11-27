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
		case <-ctx.Done():
			return nil
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

			output <- data
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	i := 0
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			outputs[i] <- data
			i = (i + 1) % len(outputs)
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	for _, in := range inputs {
		go func(input chan string) {
			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-input:
					if !ok {
						return
					}
					if !strings.Contains(data, skipMultiplexer) {
						output <- data
					}
				}
			}
		}(in)
	}

	<-ctx.Done()
	return nil
}
