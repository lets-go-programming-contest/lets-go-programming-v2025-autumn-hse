package handlers

import (
	"context"
	"errors"
	"strings"
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

			output <- data
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
				for _, out := range outputs {
					close(out)
				}
				return nil
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
						if strings.Contains(data, "no multiplexer") {
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

