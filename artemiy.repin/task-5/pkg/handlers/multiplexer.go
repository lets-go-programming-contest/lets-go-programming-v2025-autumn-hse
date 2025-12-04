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

	var wgroup sync.WaitGroup

	for _, channel := range inputs {
		if channel == nil {
			continue
		}

		wgroup.Add(1)

		go func(inChan chan string) {
			defer wgroup.Done()

			for {
				select {
				case <-ctx.Done():
					return

				case data, ok := <-inChan:
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
		}(channel)
	}

	done := make(chan struct{})

	go func() {
		wgroup.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		return nil
	case <-done:
		return nil
	}
}
