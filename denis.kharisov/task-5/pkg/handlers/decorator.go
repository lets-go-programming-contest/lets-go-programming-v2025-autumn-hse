package handlers

import (
	"context"
	"errors"
	"strings"
)

var errCantBeDecorated = errors.New("can`t be decorated")

const (
	subStrNoDecorator = "no decorator"
	prefixDecorated   = "decorated: "
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, subStrNoDecorator) {
				return errCantBeDecorated
			}

			if !strings.HasPrefix(data, prefixDecorated) {
				data = prefixDecorated + data
			}
			select {
			case <-ctx.Done():
				return nil
			case output <- data:
			}
		}
	}
}
