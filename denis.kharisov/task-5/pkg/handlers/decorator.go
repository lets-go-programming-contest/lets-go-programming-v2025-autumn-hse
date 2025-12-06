package handlers

import (
	"context"
	"errors"
	"strings"
)

var ErrCantBeDecorated = errors.New("can`t be decorated")

const (
	SubStrNoDecorator = "no decorator"
	PrefixDecorated = "decorated: "
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
			if strings.Contains(data, SubStrNoDecorator) {
				return ErrCantBeDecorated
			}
			if !strings.HasPrefix(data, PrefixDecorated) {
				data = PrefixDecorated + data
			}
			select {
			case <-ctx.Done():
				return nil
			case output <- data:
			}
		}
	}
}