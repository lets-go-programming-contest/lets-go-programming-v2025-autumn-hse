package conveyer

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

const Undefined = "undefined"

type Conveyer interface {
	RegisterDecorator(fn func(context.Context, chan string, chan string) error, input, output string) error
	RegisterMultiplexer(fn func(context.Context, []chan string, chan string) error, inputs []string, output string) error
	RegisterSeparator(fn func(context.Context, chan string, []chan string) error, input string, outputs []string) error
	Run(ctx context.Context) error
	Send(input, data string) error
	Recv(output string) (string, error)
}

type conveyerImpl struct {
	mu       sync.RWMutex
	channels map[string]chan string
	handlers []func(context.Context) error
	started  bool
	bufSize  int
}

func New(size int) Conveyer {
	return &conveyerImpl{
		channels: make(map[string]chan string),
		bufSize:  size,
	}
}

func (c *conveyerImpl) ensureChannel(name string) chan string {
	if ch, ok := c.channels[name]; ok {
		return ch
	}

	ch := make(chan string, c.bufSize)
	c.channels[name] = ch

	return ch
}

func (c *conveyerImpl) RegisterDecorator(
	fn func(context.Context, chan string, chan string) error,
	input, output string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.started {
		return errors.New("cannot register after Run")
	}

	inCh := c.ensureChannel(input)
	outCh := c.ensureChannel(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, inCh, outCh)
	})

	return nil
}

func (c *conveyerImpl) RegisterSeparator(
	fn func(context.Context, chan string, []chan string) error,
	input string, outputs []string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.started {
		return errors.New("cannot register after Run")
	}

	inCh := c.ensureChannel(input)

	outChs := make([]chan string, len(outputs))
	for i, name := range outputs {
		outChs[i] = c.ensureChannel(name)
	}

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, inCh, outChs)
	})

	return nil
}

func (c *conveyerImpl) RegisterMultiplexer(
	fn func(context.Context, []chan string, chan string) error,
	inputs []string, output string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.started {
		return errors.New("cannot register after Run")
	}

	inChs := make([]chan string, len(inputs))
	for i, name := range inputs {
		inChs[i] = c.ensureChannel(name)
	}

	outCh := c.ensureChannel(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, inChs, outCh)
	})

	return nil
}

func (c *conveyerImpl) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.started {
		c.mu.Unlock()
		return errors.New("already started")
	}

	c.started = true
	c.mu.Unlock()

	g, gCtx := errgroup.WithContext(ctx)

	for _, h := range c.handlers {

		g.Go(func() error { return h(gCtx) })
	}

	err := g.Wait()

	c.close()

	return err
}

func (c *conveyerImpl) close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, ch := range c.channels {
		func(cch chan string) {
			defer func() { recover() }()
			close(cch)
		}(ch)
	}
}

func (c *conveyerImpl) Send(input, data string) error {
	c.mu.RLock()
	ch, ok := c.channels[input]
	c.mu.RUnlock()

	if !ok {
		return errors.New("chan not found")
	}

	ch <- data

	return nil
}

func (c *conveyerImpl) Recv(output string) (string, error) {
	c.mu.RLock()
	ch, ok := c.channels[output]
	c.mu.RUnlock()

	if !ok {
		return "", errors.New("chan not found")
	}

	data, ok := <-ch
	if !ok {
		return Undefined, nil
	}

	return data, nil
}
