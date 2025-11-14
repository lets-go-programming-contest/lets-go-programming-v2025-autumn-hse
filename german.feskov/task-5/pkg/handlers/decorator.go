package handlers

import (
	"context"
	"fmt"
	"strings"
)

const (
	cantDecoratedStr = "no decorator"
	decoratedStr     = "decorated: "
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
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, cantDecoratedStr) {
				return fmt.Errorf("can't be decorated %q", data)
			}

			if !strings.HasPrefix(data, decoratedStr) {
				data = decoratedStr + data
			}

			output <- data

		}
	}
}
