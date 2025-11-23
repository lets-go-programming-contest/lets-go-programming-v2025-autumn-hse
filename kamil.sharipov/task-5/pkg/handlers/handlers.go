package handlers

import (
	"context"
	"fmt"
	"strings"
)

const (
	noMultiplexerStr = "no multiplexer"
	noDecoratorStr   = "no decorator"
	decoratedStr     = "decorated: "
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, noDecoratorStr) {
				return fmt.Errorf("can't be decorated")
			}

			if !strings.HasPrefix(data, decoratedStr) {
				data = decoratedStr + data
			}

			select {
			case <-ctx.Done():
				return nil
			case output <- data:
			}
		}
	}
}

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
			}

			index = (index + 1) % len(outputs)
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		for _, ch := range inputs {
			select {
			case <-ctx.Done():
				return nil
			case data, ok := <-ch:
				if !ok {
					continue
				}

				if strings.Contains(data, noMultiplexerStr) {
					continue
				}

				select {
				case <-ctx.Done():
					return nil
				case output <- data:
				}
			default:
			}
		}
	}
}
