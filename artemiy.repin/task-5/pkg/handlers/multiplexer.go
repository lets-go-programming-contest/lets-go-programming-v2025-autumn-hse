package handlers

import (
	"context"
	"strings"
	"sync"
)

func MultiplexerFunc(
	ctx context.Context,
	inputs []chan string,
	output chan string,
) error {
	if len(inputs) == 0 {
		<-ctx.Done()
		return nil
	}

	var wg sync.WaitGroup

	for _, ch := range inputs {
		if ch == nil {
			continue
		}

		wg.Add(1)

		go func(in chan string) {
			defer wg.Done()

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
		}(ch)
	}

	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		return nil
	case <-done:
		return nil
	}
}
