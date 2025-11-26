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
			return nil
		case val, ok := <-input:
			if !ok {
				return nil
			}
			if strings.Contains(val, "no decorator") {
				return errors.New("can't be decorated")
			}
			if !strings.HasPrefix(val, "decorated: ") {
				val = "decorated: " + val
			}
			select {
			case <-ctx.Done():
				return nil
			case output <- val:
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	i := 0
	for {
		select {
		case <-ctx.Done():
			return nil
		case val, ok := <-input:
			if !ok {
				return nil
			}
			select {
			case <-ctx.Done():
				return nil
			case outputs[i] <- val:
			}
			i = (i + 1) % len(outputs)
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
			case val, ok := <-ch:
				if !ok || strings.Contains(val, "no multiplexer") {
					continue
				}
				select {
				case <-ctx.Done():
					return nil
				case output <- val:
				}
			default:
			}
		}
	}
}
