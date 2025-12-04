package handlers

import (
	"context"
	"errors"
	"strings"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}
			if strings.Contains(data, "no decorator") {
				return errors.New("can't be decorated")
			}
			if !strings.HasPrefix(data, "decorated: ") {
				data = "decorated: " + data
			}
			select {
			case output <- data:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	counter := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-input:
			if !ok {
				return nil
			}
			idx := counter % len(outputs)
			select {
			case outputs[idx] <- data:
			case <-ctx.Done():
				return ctx.Err()
			}
			counter++
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-inputs[0]:
			if !ok {
				return nil
			}
			if strings.Contains(data, "no multiplexer") {
				continue
			}
			select {
			case output <- data:
			case <-ctx.Done():
				return ctx.Err()
			}
		case data, ok := <-inputs[1]:
			if !ok {
				return nil
			}
			if strings.Contains(data, "no multiplexer") {
				continue
			}
			select {
			case output <- data:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}
