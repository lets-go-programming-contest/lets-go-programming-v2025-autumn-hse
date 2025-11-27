package handlers

import (
	"context"
	"strings"
	"sync"
)

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	var inputsWg sync.WaitGroup

	for _, input := range inputs {
		inputsWg.Add(1)
		go func(in chan string) {
			defer inputsWg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-in:
					if !ok {
						return
					}

					if strings.Contains(data, "no multiplexer") {
						continue
					}

					select {
					case <-ctx.Done():
						return
					case output <- data:
					}
				}
			}
		}(input)
	}

	inputsWg.Wait()

	return nil
}
