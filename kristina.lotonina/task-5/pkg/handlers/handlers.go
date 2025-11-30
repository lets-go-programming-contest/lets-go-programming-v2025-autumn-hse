package handlers

import (
	"context"
	"errors"
	"strings"
)

func PrefixDecoratorFunc(ctx context.Context, in chan string, out chan string) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-in:
			if !ok {
				close(out)
				return nil
			}
			if strings.Contains(data, "no decorator") {
				return errors.New("can't be decorated")
			}
			if !strings.HasPrefix(data, "decorated: ") {
				data = "decorated: " + data
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case out <- data:
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, in chan string, outs []chan string) error {
	i := 0
	for {
		select {
		case <-ctx.Done():
			for _, ch := range outs {
				close(ch)
			}
			return ctx.Err()
		case data, ok := <-in:
			if !ok {
				for _, ch := range outs {
					close(ch)
				}
				return nil
			}
			idx := i % len(outs)
			i++
			select {
			case <-ctx.Done():
				for _, ch := range outs {
					close(ch)
				}
				return ctx.Err()
			case outs[idx] <- data:
			}
		default:
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	openChannels := len(inputs)

	for openChannels > 0 {
		for i := 0; i < len(inputs); i++ {
			ch := inputs[i]
			if ch == nil {
				continue
			}

			select {
			case <-ctx.Done():
				close(output)
				return ctx.Err()
			case data, ok := <-ch:
				if !ok {
					openChannels--
					inputs[i] = nil
					continue
				}

				if strings.Contains(data, "no multiplexer") {
					continue
				}

				select {
				case <-ctx.Done():
					close(output)
					return ctx.Err()
				case output <- data:
				}
			default:
			}
		}
	}

	close(output)
	return nil
}
