package handlers

import (
	"context"
	"fmt"
	"strings"
	"sync"
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
			return ctx.Err()
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

			output <- data
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	counter := 0
	numOutputs := len(outputs)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			outputIndex := counter % numOutputs
			counter++

			outputs[outputIndex] <- data
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)

		go func(input chan string) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-input:
					if !ok {
						return
					}

					if strings.Contains(data, noMultiplexerStr) {
						continue
					}

					output <- data
				}
			}

		}(input)

	}

	go func() {
		wg.Wait()
	}()

	return nil
}
