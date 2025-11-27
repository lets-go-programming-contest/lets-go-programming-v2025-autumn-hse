package handlers

import (
	"context"
	"errors"
	"strings"
)

var ErrNoDecorator = errors.New("can't be decorated")

func PrefixDecoratorFunc(ctx context.Context,
	inputChan <-chan string,
	outputChan chan<- string) error {
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
				continue
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func SeparatorFunc(ctx context.Context,
	inputChan <-chan string,
	outputChans []chan<- string) error {
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
			case outputChans[chanIndex] <- value:
				chanIndex = (chanIndex + 1) % chanCount
				continue
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context,
	inputChans []<-chan string,
	outputChan chan<- string) error {
	for _, ch := range inputChans {
		go func(input <-chan string) {
			for {
				select {
				case <-ctx.Done():
					return

				case value, ok := <-input:
					if !ok {
						return
					}

					if strings.Contains(value, "no multiplexer") {
						continue
					}

					select {
					case outputChan <- value:
						continue
					case <-ctx.Done():
						return
					}
				}
			}
		}(ch)
	}

	<-ctx.Done()
	return ctx.Err()
}
