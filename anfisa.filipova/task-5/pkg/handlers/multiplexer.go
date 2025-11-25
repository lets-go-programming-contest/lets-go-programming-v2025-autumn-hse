package handlers

import (
	"context"
	"strings"
	"sync"
)

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	var wg sync.WaitGroup

	for i, input := range inputs {
		wg.Add(1)
		go func(index int, in chan string) {
			defer wg.Done()
			for {
				select {
				case data, ok := <-in:
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
				case <-ctx.Done():
					return
				}
			}
		}(i, input)
	}

	wg.Wait()
	return nil
}
