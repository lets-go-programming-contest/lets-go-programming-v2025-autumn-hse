package handlers

import (
	"context"
	"strings"
	"sync"
)

const (
	noMultiplexer = "no multiplexer"
)

func MultiplexerFunc(
	ctx context.Context,
	inputChannels []chan string,
	outputChannel chan string,
) error {
	if len(inputChannels) == 0 {
		close(outputChannel)

		return nil
	}

	return multiplexerHelper(ctx, inputChannels, outputChannel)
}

func multiplexerHelper(
	ctx context.Context,
	inputChannels []chan string,
	outputChannel chan string,
) error {
	var waitGroup sync.WaitGroup

	for _, inputChannel := range inputChannels {
		if inputChannel == nil {
			continue
		}

		waitGroup.Add(1)

		worker := func(channel chan string) {
			defer waitGroup.Done()

			for {
				select {
				case <-ctx.Done():
					return

				case data, receivedOk := <-channel:
					if !receivedOk {
						return
					}

					if strings.Contains(data, noMultiplexer) {
						continue
					}

					select {
					case <-ctx.Done():
						return
					case outputChannel <- data:
					}
				}
			}
		}

		go worker(inputChannel)
	}

	waitGroup.Wait()
	close(outputChannel)

	return nil
}
