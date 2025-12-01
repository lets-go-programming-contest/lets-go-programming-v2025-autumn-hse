package handlers

import (
	"context"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	defer func() {
		close(output)
	}()

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
	defer func() {
		for _, output := range outputs {
			close(output)
		}
	}()

	counter := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}

			outputIndex := counter % len(outputs)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case outputs[outputIndex] <- data:
			}

			counter++
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	defer close(output)

	g, ctx := errgroup.WithContext(ctx)
	merged := make(chan string, len(inputs)*10)

	for _, input := range inputs {
		input := input
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case data, ok := <-input:
					if !ok {
						return nil
					}
					if strings.Contains(data, "no multiplexer") {
						continue
					}
					select {
					case <-ctx.Done():
						return ctx.Err()
					case merged <- data:
					}
				}
			}
		})
	}

	go func() {
		g.Wait()
		close(merged)
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