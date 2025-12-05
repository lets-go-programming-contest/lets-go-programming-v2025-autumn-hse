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

			for _, channelRef := range inputs {
				select {
				case value, ok := <-channelRef:
					if !ok {
						continue
					}

					received = true

					if strings.Contains(value, "no multiplexer") {
						continue
					}

					output <- value
				default:
				}
			}

			if !received {
				continue
			}
		}
	}
}
