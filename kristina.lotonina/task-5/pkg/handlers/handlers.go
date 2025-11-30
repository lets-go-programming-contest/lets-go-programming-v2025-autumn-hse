package handlers

import (
	"context"
	"errors"
	"strings"
	"golang.org/x/sync/errgroup"
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
		}
	}
}

func MultiplexerFunc(ctx context.Context, ins []chan string, out chan string) error {
	merged := make(chan string)
	g, ctx := errgroup.WithContext(ctx)

	for _, inChan := range ins {
		ch := inChan
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case data, ok := <-ch:
					if !ok {
						return nil
					}
					if !strings.Contains(data, "no multiplexer") {
						select {
						case <-ctx.Done():
							return ctx.Err()
						case merged <- data:
						}
					}
				}
			}
		})
	}

	go func() {
		g.Wait()
		close(merged)
	}()

	for {
		select {
		case <-ctx.Done():
			close(out)
			return ctx.Err()
		case data, ok := <-merged:
			if !ok {
				close(out)
				return nil
			}
			select {
			case <-ctx.Done():
				close(out)
				return ctx.Err()
			case out <- data:
			}
		}
	}
}
