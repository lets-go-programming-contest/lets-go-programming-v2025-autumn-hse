package handlers

import (
	"context"
	"errors"
	"strings"
)

const (
	noDecorator      = "no decorator"
	noMultiplexer    = "no multiplexer"
	alreadyDecorated = "decorated: "
)

var (
	errCantBeDecorated = errors.New("can't be decorated")
	errSeparator       = errors.New("error in separator")
	errPrefixDecorator = errors.New("error in prefdecorator")
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil

		case data, ok := <-input:
			if !ok {

				return errPrefixDecorator
			}

			if strings.Contains(data, noDecorator) {
				return errCantBeDecorated
			}

			if !strings.HasPrefix(data, alreadyDecorated) {
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
	var counter int

	for {
		select {
		case <-ctx.Done():
			return nil

		case data, ok := <-input:
			if !ok {

				return errSeparator
			}

			index := counter % len(outputs)

			select {
			case outputs[index] <- data:
			case <-ctx.Done():
				return nil
			}

			counter++
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	channels := make([]<-chan string, len(inputs))
	for i, in := range inputs {
		channels[i] = in
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		default:
			for _, ch := range channels {
				select {
				case data, ok := <-ch:
					if ok {
						if strings.Contains(data, noMultiplexer) {
							continue
						}

						select {
						case output <- data:
						case <-ctx.Done():
							return nil
						}
					}
				default:
					continue
				}
			}
		}
	}
}
