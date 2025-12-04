package conveyer

import (
	"context"
	"errors"
	"sync"
)

var ErrChanNotFound = errors.New("chan not found")

type handlerFunc func(ctx context.Context) error

type Conveyer struct {
	size int

	chans map[string]chan string

	mu sync.RWMutex

	handlers []handlerFunc
}

func New(size int) *Conveyer {
	return &Conveyer{
		size:     size,
		chans:    make(map[string]chan string),
		handlers: make([]handlerFunc, 0),
	}
}

func (c *Conveyer) getOrCreate(id string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	ch, ok := c.chans[id]
	if !ok {
		ch = make(chan string, c.size)
		c.chans[id] = ch
	}
	return ch
}

func (c *Conveyer) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	in := c.getOrCreate(input)
	out := c.getOrCreate(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, in, out)
	})
}

func (c *Conveyer) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	in := make([]chan string, 0, len(inputs))
	for _, id := range inputs {
		in = append(in, c.getOrCreate(id))
	}
	out := c.getOrCreate(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, in, out)
	})
}

func (c *Conveyer) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	in := c.getOrCreate(input)
	outs := make([]chan string, 0, len(outputs))

	for _, id := range outputs {
		outs = append(outs, c.getOrCreate(id))
	}

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, in, outs)
	})
}

func (c *Conveyer) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	for _, h := range c.handlers {
		wg.Add(1)
		go func(h handlerFunc) {
			defer wg.Done()
			if err := h(ctx); err != nil {
				errCh <- err
			}
		}(h)
	}

	go func() {
		wg.Wait()
		errCh <- nil
	}()

	err := <-errCh

	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, ch := range c.chans {
		close(ch)
	}

	return err
}

func (c *Conveyer) Send(id string, data string) error {
	c.mu.RLock()
	ch, ok := c.chans[id]
	c.mu.RUnlock()

	if !ok {
		return ErrChanNotFound
	}

	ch <- data
	return nil
}

func (c *Conveyer) Recv(id string) (string, error) {
	c.mu.RLock()
	ch, ok := c.chans[id]
	c.mu.RUnlock()

	if !ok {
		return "", ErrChanNotFound
	}

	v, ok := <-ch
	if !ok {
		return "undefined", nil
	}
	return v, nil
}
