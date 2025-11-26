package handlers

import (
	"context"
	"errors"
	"strings"
)

func PrefixDecoratorFunc(ctx context.Context, in, out chan string) error {
	for {
		if ctx.Err() != nil {
			return nil
		}
		val, ok := <-in
		if !ok {
			return nil
		}
		if strings.Contains(val, "no decorator") {
			return errors.New("can't be decorated")
		}
		if !strings.HasPrefix(val, "decorated: ") {
			val = "decorated: " + val
		}
		out <- val
	}
}

func SeparatorFunc(ctx context.Context, in chan string, outs []chan string) error {
	i := 0
	for {
		if ctx.Err() != nil {
			return nil
		}
		val, ok := <-in
		if !ok {
			return nil
		}
		outs[i] <- val
		i = (i + 1) % len(outs)
	}
}

func MultiplexerFunc(ctx context.Context, ins []chan string, out chan string) error {
	for {
		if ctx.Err() != nil {
			return nil
		}
		for _, ch := range ins {
			val, ok := <-ch
			if !ok || strings.Contains(val, "no multiplexer") {
				continue
			}
			out <- val
		}
	}
}
