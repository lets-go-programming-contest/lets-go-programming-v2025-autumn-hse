package handlers

import (
	"context"
	"strings"
)

func MultiplexerFunc(
	ctx context.Context,
	inputs []chan string,
	output chan string,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil

		default:
			received := false
			for _, ch := range inputs {
				select {
				case v, ok := <-ch:
					if !ok {
						continue
					}
					received = true

					if strings.Contains(v, "no multiplexer") {
						continue
					}

					output <- v
				default:
				}
			}
			if !received {
				continue
			}
		}
	}
}
