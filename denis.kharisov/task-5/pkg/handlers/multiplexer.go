package handlers

import (
	"context"
	"strings"
	"sync"
)

const subStrNoMultiplexer = "no multiplexer"

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		<-ctx.Done()

		return nil
	}

	var waitGroup sync.WaitGroup
	for _, channel := range inputs {
		waitGroup.Add(1)

		go func(inputCh chan string) {
			defer waitGroup.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-inputCh:
					if !ok {
						return
					}

					if strings.Contains(data, subStrNoMultiplexer) {
						continue
					}

					select {
					case output <- data:
					case <-ctx.Done():
						return
					}
				}
			}
		}(channel)
	}
	waitGroup.Wait()
	close(output)

	return nil
}
