package conveyer

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	ErrChanNotFound    = "chan not found"
	ErrUndefined       = "undefined"
	ErrConveyerStopped = "conveyer stopped"
)

type conveyerImpl struct {
	size     int
	channels map[string]chan string
	handlers []handler
	mutex    sync.RWMutex
	stopped  bool
}

type handler struct {
	handlerType string
	fn          interface{}
	inputs      []string
	outputs     []string
}

func New(size int) *conveyerImpl {
	return &conveyerImpl{
		size:     size,
		channels: make(map[string]chan string),
		handlers: make([]handler, 0),
		stopped:  false,
	}
}

func (c *conveyerImpl) getOrCreateChannel(id string) chan string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if ch, exists := c.channels[id]; exists {
		return ch
	}

	ch := make(chan string, c.size)
	c.channels[id] = ch
	return ch
}

func (c *conveyerImpl) getChannel(id string) (chan string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if c.stopped {
		return nil, errors.New(ErrConveyerStopped)
	}

	if ch, exists := c.channels[id]; exists {
		return ch, nil
	}

	return nil, errors.New(ErrChanNotFound)
}

func (c *conveyerImpl) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.getOrCreateChannel(input)
	c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handler{
		handlerType: "decorator",
		fn:          fn,
		inputs:      []string{input},
		outputs:     []string{output},
	})
}

func (c *conveyerImpl) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, input := range inputs {
		c.getOrCreateChannel(input)
	}
	c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handler{
		handlerType: "multiplexer",
		fn:          fn,
		inputs:      inputs,
		outputs:     []string{output},
	})
}

func (c *conveyerImpl) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.getOrCreateChannel(input)
	for _, output := range outputs {
		c.getOrCreateChannel(output)
	}

	c.handlers = append(c.handlers, handler{
		handlerType: "separator",
		fn:          fn,
		inputs:      []string{input},
		outputs:     outputs,
	})
}

func (c *conveyerImpl) Run(ctx context.Context) error {
	if c.stopped {
		return errors.New(ErrConveyerStopped)
	}

	g, ctx := errgroup.WithContext(ctx)

	for _, h := range c.handlers {
		h := h
		g.Go(func() error {
			return c.runHandler(ctx, h)
		})
	}

	return g.Wait()
}

func (c *conveyerImpl) runHandler(ctx context.Context, h handler) error {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	switch h.handlerType {
	case "decorator":
		if fn, ok := h.fn.(func(ctx context.Context, input chan string, output chan string) error); ok {
			inputChan, err := c.getChannel(h.inputs[0])
			if err != nil {
				return err
			}
			outputChan, err := c.getChannel(h.outputs[0])
			if err != nil {
				return err
			}
			return fn(ctx, inputChan, outputChan)
		}
	case "multiplexer":
		if fn, ok := h.fn.(func(ctx context.Context, inputs []chan string, output chan string) error); ok {
			inputChans := make([]chan string, len(h.inputs))
			for i, input := range h.inputs {
				ch, err := c.getChannel(input)
				if err != nil {
					return err
				}
				inputChans[i] = ch
			}
			outputChan, err := c.getChannel(h.outputs[0])
			if err != nil {
				return err
			}
			return fn(ctx, inputChans, outputChan)
		}
	case "separator":
		if fn, ok := h.fn.(func(ctx context.Context, input chan string, outputs []chan string) error); ok {
			inputChan, err := c.getChannel(h.inputs[0])
			if err != nil {
				return err
			}
			outputChans := make([]chan string, len(h.outputs))
			for i, output := range h.outputs {
				ch, err := c.getChannel(output)
				if err != nil {
					return err
				}
				outputChans[i] = ch
			}
			return fn(ctx, inputChan, outputChans)
		}
	}

	return errors.New("unknown handler type")
}

func (c *conveyerImpl) Send(input string, data string) error {
	ch, err := c.getChannel(input)
	if err != nil {
		return err
	}

	if c.stopped {
		return errors.New(ErrConveyerStopped)
	}

	select {
	case ch <- data:
		return nil
	default:
		return errors.New("channel is full")
	}
}

func (c *conveyerImpl) Recv(output string) (string, error) {
	ch, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	select {
	case data, ok := <-ch:
		if !ok {
			return "", errors.New(ErrUndefined)
		}
		return data, nil
	}
}

func (c *conveyerImpl) Stop() {
	c.mutex.Lock()
	c.stopped = true
	for id, ch := range c.channels {
		close(ch)
		delete(c.channels, id)
	}
	c.mutex.Unlock()
}