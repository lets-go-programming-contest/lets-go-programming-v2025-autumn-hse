package handlers

import (
	"context"
	"strings"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, "no decorator") {
				continue
			}

			if !strings.HasPrefix(data, "decorated: ") {
				data = "decorated: " + data
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case output <- data:
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	counter := 0
	outputCount := len(outputs)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			for attempt := 0; attempt < outputCount; attempt++ {
				outputIndex := (counter + attempt) % outputCount
				select {
				case <-ctx.Done():
					return ctx.Err()
				case outputs[outputIndex] <- data:
					counter = (outputIndex + 1) % outputCount
					break
				default:
					continue
				}
				break
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	merged := make(chan string, len(inputs)*2)

	done := make(chan struct{}, len(inputs))

	for _, in := range inputs {
		go func(inputChan chan string) {
			defer func() {
				done <- struct{}{}
			}()

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-inputChan:
					if !ok {
						return
					}

					if strings.Contains(data, "no multiplexer") {
						continue
					}

					select {
					case <-ctx.Done():
						return
					case merged <- data:
					}
				}
			}
		}(in)
	}

	go func() {
		completed := 0
		for range done {
			completed++
			if completed == len(inputs) {
				close(merged)
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-merged:
			if !ok {
				return nil
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case output <- data:
			}
		}
	}
}
