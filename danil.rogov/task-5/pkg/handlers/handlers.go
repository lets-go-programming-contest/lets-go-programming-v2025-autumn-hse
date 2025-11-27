package handlers

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/sync/errgroup"
)

var ErrNoDecorator = errors.New("can't be decorated")

func PrefixDecoratorFunc(ctx context.Context, inputChan, outputChan chan string) error {
	prefix := "decorated: "

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case value, ok := <-inputChan:
			if !ok {
				return nil
			}

			if strings.Contains(value, "no decorator") {
				return ErrNoDecorator
			}

			if !strings.HasPrefix(value, "decorated:") {
				value = prefix + value
			}

			select {
			case outputChan <- value:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, inputChan chan string, outputChans []chan string) error {
	var (
		chanIndex int
		chanCount int = len(outputChans)
	)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		case value, ok := <-inputChan:
			if !ok {
				return nil
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case outputChans[chanIndex] <- value:
				chanIndex = (chanIndex + 1) % chanCount
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputChans []chan string, outputChan chan string) error {
	errGroup, ctx := errgroup.WithContext(ctx)

	for _, ch := range inputChans {
		inputChan := ch
		errGroup.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case value, ok := <-inputChan:
					if !ok {
						return nil
					}

					if strings.Contains(value, "no multiplexer") {
						continue
					}

					select {
					case outputChan <- value:
					case <-ctx.Done():
						return ctx.Err()
					}
				}
			}
		})
	}

	return errGroup.Wait()
}
