package handlers

import (
	"context"
	"errors"
	"strings"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for val := range input {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if strings.Contains(val, "no decorator") {
				return errors.New("can't be decorated")
			}
			if !strings.HasPrefix(val, "decorated: ") {
				val = "decorated: " + val
			}
			output <- val
		}
	}
	return nil
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return errors.New("no output channels")
	}
	index := 0
	for val := range input {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			outputs[index] <- val
			index = (index + 1) % len(outputs)
		}
	}
	return nil
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return errors.New("no input channels")
	}
	for {
		allClosed := true
		for _, ch := range inputs {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case val, ok := <-ch:
				if ok {
					allClosed = false
					if strings.Contains(val, "no multiplexer") {
						continue
					}
					output <- val
				}
			default:
				allClosed = false
			}
		}
		if allClosed {
			return nil
		}
	}
}
