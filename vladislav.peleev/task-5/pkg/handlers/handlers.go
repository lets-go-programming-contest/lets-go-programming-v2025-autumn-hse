package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var errCantBeDecorated = errors.New("can't be decorated")

const (
	decoratedPrefix   = "decorated: "
	noDecoratorSubstr = "no decorator"
	noMultiplexSubstr = "no multiplexer"
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

			if strings.Contains(data, noDecoratorSubstr) {
				return ErrCantBeDecorated
			}

			if !strings.HasPrefix(data, decoratedPrefix) {
				data = decoratedPrefix + data
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
	var waitGroup sync.WaitGroup

	for _, ch := range inputs {
		inCh := ch
		waitGroup.Add(1)

		go func(inputChan chan string) {
			defer waitGroup.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-inputChan:
					if !ok {
						return
					}

					if strings.Contains(data, noMultiplexSubstr) {
						continue
					}

					select {
					case output <- data:
					case <-ctx.Done():
						return
					}
				}
			}
		}(inCh)
	}

	waitGroup.Wait()

	return nil
}
