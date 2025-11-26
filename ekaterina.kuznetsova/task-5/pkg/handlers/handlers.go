package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

func PrefixDecoratorFunc(ctx context.Context, input, output chan string) error {
	defer close(output)
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
			output <- val
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return errors.New("no outputs")
	}
	defer func() { for _, ch := range outputs { close(ch) } }()

	i := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case val, ok := <-input:
			if !ok {
				return nil
			}
			outputs[i] <- val
			i = (i + 1) % len(outputs)
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return errors.New("no inputs")
	}
	defer close(output)

	var wg sync.WaitGroup
	for _, ch := range inputs {
		wg.Add(1)
		go func(c chan string) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-c:
					if !ok {
						return
					}
					if strings.Contains(v, "no multiplexer") {
						continue
					}
					output <- v
				}
			}
		}(ch)
	}
	wg.Wait()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
