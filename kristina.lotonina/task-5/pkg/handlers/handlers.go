package handlers

import (
	"context"
	"errors"
	"string"
)

func Decorator(ctx context.Context, in chan string, out chan string) error {
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
				return errors.New("can't decorate")
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

func Separator(ctx context.Context, in chan string, outs []chan string) error {
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

func Multiplexer(ctx context.Context, ins []chan string, out chan string) error {
	merged := make(chan string, 64)

	for _, in := range ins {
		go func(ch chan string) {
			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-ch:
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
		}(in)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case data := <-merged:
			select {
			case <-ctx.Done():
				return ctx.Err()
			case out <- data:
			}
		}
	}
}