package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var ErrCantBeDecorated = errors.New("can't be decorated")

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				close(output)
				return nil
			}
			if strings.Contains(data, "no decorator") {
				return ErrCantBeDecorated
			}
			if !strings.HasPrefix(data, "decorated: ") {
				data = "decorated: " + data
			}
			select {
			case output <- data:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	counter := 0
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				for _, ch := range outputs {
					close(ch)
				}
				return nil
			}
			idx := counter % len(outputs)
			select {
			case outputs[idx] <- data:
			case <-ctx.Done():
				return nil
			}
			counter++
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	var wg sync.WaitGroup

	for _, ch := range inputs {
		wg.Add(1)
		go func(rc chan string) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-rc:
					if !ok {
						return
					}
					if strings.Contains(data, "no multiplexer") {
						continue
					}
					select {
					case output <- data:
					case <-ctx.Done():
						return
					}
				}
			}
		}(ch)
	}

	wg.Wait()
	close(output)

	return nil
}
