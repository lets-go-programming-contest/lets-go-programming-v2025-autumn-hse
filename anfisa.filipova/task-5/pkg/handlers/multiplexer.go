package handlers

import (
	"context"
	"strings"
)

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	for {
		for _, channel := range inputs {
			select {
			case data, ok := <-channel:
				if !ok {
					continue
				}
				if strings.Contains(data, "no multiplexer") {
					continue
				}
				select {
				case output <- data:
				case <-ctx.Done():
					return ctx.Err()
				}
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
		}
	}
}
