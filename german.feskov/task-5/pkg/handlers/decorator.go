package handlers

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

const (
	cantDecoratedStr = "no decorator"
	decoratedStr     = "decorated: "
)

var ErrCantBeDecorated = errors.New("can't be decorated")

func PrefixDecoratorFunc(
	ctx context.Context,
	input chan string,
	output chan string,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, cantDecoratedStr) {
				return fmt.Errorf("std %q: %w", data, ErrCantBeDecorated)
			}

			if !strings.HasPrefix(data, decoratedStr) {
				data = decoratedStr + data
			}

			output <- data
		}
	}
}
