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
				return ctx.Err()
			case output <- val:
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return errors.New("no output channels")
	}
	index := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case val, ok := <-input:
			if !ok {
				return nil
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case outputs[index] <- val:
			}
			index = (index + 1) % len(outputs)
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return errors.New("no input channels")
	}

	active := make([]bool, len(inputs))
	for i := range active {
		active[i] = true
	}

	for {
		for i, ch := range inputs {
			if !active[i] {
				continue
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case val, ok := <-ch:
				if !ok {
					active[i] = false
					continue
				}
				if strings.Contains(val, "no multiplexer") {
					continue
				}
				select {
				case <-ctx.Done():
					return ctx.Err()
				case output <- val:
				}
			default:
			}
		}
		allClosed := true
		for _, a := range active {
			if a {
				allClosed = false
				break
			}
		}
		if allClosed {
			return nil
		}
	}
}
