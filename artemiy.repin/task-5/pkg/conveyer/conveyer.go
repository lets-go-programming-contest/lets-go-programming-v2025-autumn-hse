package conveyer

import (
	"context"
	"errors"
	"sync"
)

const Undefined = "undefined"

var (
	ErrChannelNotFound = errors.New("chan not found")
)

type Conveyer struct {
	mu       sync.Mutex
	channels map[string]chan string
	bufSize  int

	handlers []func(context.Context) error
}

func New(size int) *Conveyer {
	if size < 0 {
		size = 0
	}

	return &Conveyer{
		channels: make(map[string]chan string),
		bufSize:  size,
		handlers: make([]func(context.Context) error, 0),
	}
}

func (c *Conveyer) getChannel(name string) (chan string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ch, ok := c.channels[name]
	return ch, ok
}

func (c *Conveyer) ensureChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	ch, ok := c.channels[name]
	if ok {
		return ch
	}

	ch = make(chan string, c.bufSize)
	c.channels[name] = ch
	return ch
}

func (c *Conveyer) closeAllChannels() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, ch := range c.channels {
		if ch != nil {
			close(ch)
		}
	}
}

func (c *Conveyer) Run(ctx context.Context) error {
	defer c.closeAllChannels()

	if len(c.handlers) == 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error, len(c.handlers))
	var wg sync.WaitGroup

	for _, h := range c.handlers {
		handler := h

		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := handler(ctx); err != nil {
				select {
				case errCh <- err:
					cancel()
				default:
				}
			}
		}()
	}

	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
	}

	return nil
}

func (c *Conveyer) Send(name, data string) error {
	ch, ok := c.getChannel(name)
	if !ok || ch == nil {
		return ErrChannelNotFound
	}

	ch <- data
	return nil
}

func (c *Conveyer) Recv(name string) (string, error) {
	ch, ok := c.getChannel(name)
	if !ok || ch == nil {
		return Undefined, ErrChannelNotFound
	}

	val, ok := <-ch
	if !ok {
		return Undefined, nil
	}

	return val, nil
}

func (c *Conveyer) RegisterDecorator(
	fn func(context.Context, chan string, chan string) error,
	input string,
	output string,
) {
	inCh := c.ensureChannel(input)
	outCh := c.ensureChannel(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, inCh, outCh)
	})
}

func (c *Conveyer) RegisterMultiplexer(
	fn func(context.Context, []chan string, chan string) error,
	inputs []string,
	output string,
) {
	inChans := make([]chan string, len(inputs))
	for i, name := range inputs {
		inChans[i] = c.ensureChannel(name)
	}
	outCh := c.ensureChannel(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, inChans, outCh)
	})
}

func (c *Conveyer) RegisterSeparator(
	fn func(context.Context, chan string, []chan string) error,
	input string,
	outputs []string,
) {
	inCh := c.ensureChannel(input)
	outChans := make([]chan string, len(outputs))
	for i, name := range outputs {
		outChans[i] = c.ensureChannel(name)
	}

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, inCh, outChans)
	})
}
