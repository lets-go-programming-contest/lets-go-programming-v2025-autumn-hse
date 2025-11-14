package handlers

import (
	"context"
	"fmt"
	"strings"
)

const noMultiplexerStr = "no multiplexer"

func MultiplexerFunc(
	ctx context.Context,
	inputs []chan string,
	output chan string,
) error {
	for {
		closedChan := 0
		for _, ch := range inputs {
			select {
			case <-ctx.Done():
				return nil
			case data, ok := <-ch:
				if !ok {
					closedChan++
					continue
				}

				if strings.Contains(data, noMultiplexerStr) {
					return fmt.Errorf("cant be multiplex %q", data)
				}

				output <- data

			}
		}
		if closedChan == len(inputs) {
			return nil
		}
	}
}
