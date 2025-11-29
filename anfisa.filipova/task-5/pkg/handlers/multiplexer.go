package handlers

import (
	"context"
	"strings"
)

const noMultiplexer = "no multiplexer"

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	for {
		for _, channel := range inputs {
			select {
			case <-ctx.Done():
				return nil

			case data, ok := <-channel:
				if !ok {
					continue
				}

				if strings.Contains(data, noMultiplexer) {
					continue
				}

				output <- data
			default:
			}
		}
	}
}
