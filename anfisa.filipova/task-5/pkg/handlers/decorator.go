package handlers

import (
	"context"
	"errors"
	"strings"
)

const (
	noDecorator     = "no decorator"
	prefixDecorated = "decorated: "
)

var errCannotBeDecorated = errors.New("can't be decorated")

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, noDecorator) {
				return errCannotBeDecorated
			}

			if !strings.HasPrefix(data, prefixDecorated) {
				data = prefixDecorated + data
			}

			output <- data
		}
	}
}
