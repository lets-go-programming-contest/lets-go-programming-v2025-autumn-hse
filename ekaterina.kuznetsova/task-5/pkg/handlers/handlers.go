package handlers

import (
	"context"
	"errors"
	"strings"
)

func PrefixDecoratorFunc(ctx context.Context, in, out chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case val, ok := <-in:
			if !ok {
				return nil
			}
			if strings.Contains(val, "no decorator") {
				return errors.New("can't be decorated")
			}
			if !strings.HasPrefix(val, "decorated: ") {
				val = "decorated: " + val
			}
			select {
			case <-ctx.Done():
				return nil
			case out <- val:
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, in chan string, outs []chan string) error {
	i := 0
	for {
		select {
		case <-ctx.Done():
			return nil
		case val, ok := <-in:
			if !ok {
				return nil
			}
			select {
			case <-ctx.Done():
				return nil
			case outs[i] <- val:
			}
			i = (i + 1) % len(outs)
		}
	}
}

func MultiplexerFunc(ctx context.Context, ins []chan string, out chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		for _, ch := range ins {
			select {
			case <-ctx.Done():
				return nil
			case val, ok := <-ch:
				if !ok || strings.Contains(val, "no multiplexer") {
					continue
				}
				select {
				case <-ctx.Done():
					return nil
				case out <- val:
				}
			default:
			}
		}
	}
}
