package handlers

import (
	"context"
	"errors"
	"strings"
)

var ErrNoDecorator = errors.New("can't be decorated")

func PrefixDecoratorFunc(ctx context.Context, inputChan, outputChan chan string) error {
	prefix := "decorated: "

	for {
		select {
		case <-ctx.Done():
			return nil

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
				continue
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, inputChan chan string, outputChans []chan string) error {
	var (
		chanIndex int
		chanCount = len(outputChans)
	)

	for {
		select {
		case <-ctx.Done():
			return nil
		case value, ok := <-inputChan:
			if !ok {
				return nil
			}

			select {
			case outputChans[chanIndex] <- value:
				chanIndex = (chanIndex + 1) % chanCount

				continue
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputChans []chan string, outputChan chan string) error {
	for {
		for _, channel := range inputChans {
			select {
			case <-ctx.Done():
				return nil

			case value, ok := <-channel:
				if !ok {
					continue
				}

				if strings.Contains(value, "no multiplexer") {
					continue
				}

				select {
				case outputChan <- value:
					continue
				case <-ctx.Done():
					return nil
				}
			default:
				continue
			}
		}
	}
}
