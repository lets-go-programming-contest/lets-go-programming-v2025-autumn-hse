package handlers

import (
	"context"
	"errors"
	"strings"
)

var ErrCantDecorate = errors.New("can't be decorated")

const (
	nonDecorate     = "no decorator"
	prefixDecorated = "decorated: "
)

func PrefixDecoratorFunc(
	ctx context.Context,
	input chan string,
	output chan string,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil

		case value, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(value, nonDecorate) {
				return ErrCantDecorate
			}

			if !strings.HasPrefix(value, prefixDecorated) {
				value = prefixDecorated + value
			}

			output <- value
		}
	}
}
