package handlers

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

const noMultiplexerStr = "no multiplexer"

var ErrCantBeMultiplex = errors.New("cant be multiplex")

func MultiplexerFunc(
	ctx context.Context,
	inputs []chan string,
	output chan string,
) error {
	for closedChan := 0; closedChan != len(inputs); {

		for _, ch := range inputs {
			select {
			case <-ctx.Done():
				return nil
			case data, ok := <-ch:
				fmt.Println(data)
				if !ok {
					closedChan++

					continue
				}

				if strings.Contains(data, noMultiplexerStr) {
					continue
				}

				output <- data
			default:
				continue
			}
		}
	}
	return nil
}
