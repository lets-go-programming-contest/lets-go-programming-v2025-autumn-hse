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
		default:
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case data, ok := <-in:
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
			return ctx.Err()
		default:
		}

		select {
		case <-ctx.Done():
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
				return ctx.Err()
			case outs[idx] <- data:
			}
		}
	}
}

func MultiplexerFunc(ctx context.Context, ins []chan string, out chan string) error {
	merged := make(chan string, 64)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, in := range ins {
		inChan := in
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-inChan:
					if !ok {
						return
					}
					if !strings.Contains(data, "no multiplexer") {
						select {
						case <-ctx.Done():
							return
						case merged <- data:
						}
					}
				}
			}
		}()
	}

	for {
		select {
		case <-ctx.Done():
			close(out)
			return ctx.Err()
		case data := <-merged:
			select {
			case <-ctx.Done():
				close(out)
				return ctx.Err()
			case out <- data:
			}
		}
	}
}
