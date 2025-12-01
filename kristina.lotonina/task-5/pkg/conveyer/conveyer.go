package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	ErrChanNotFound    = "chan not found"
	ErrUndefined       = "undefined"
	ErrConveyerStopped = "conveyer stopped"
	ErrChannelFull     = "channel is full"
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
		mutex:    sync.RWMutex{},
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

	ch, exists := c.channels[id]
	if !exists {
		return nil, errors.New(ErrChanNotFound)
	}

	return ch, nil
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
	c.mutex.Lock()
	if c.stopped {
		c.mutex.Unlock()
		return errors.New(ErrConveyerStopped)
	}
	c.mutex.Unlock()

	g, gctx := errgroup.WithContext(ctx)

	for _, h := range c.handlers {
		h := h
		g.Go(func() error {
			return c.executeHandler(gctx, h)
		})
	}

	if err := g.Wait(); err != nil {
		c.Stop()
		return fmt.Errorf("conveyer failed: %w", err)
	}

	return nil
}

func (c *conveyerImpl) executeHandler(ctx context.Context, h handler) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if c.stopped {
		return nil
	}

	switch h.handlerType {
	case "decorator":
		fn, ok := h.fn.(func(ctx context.Context, input chan string, output chan string) error)
		if !ok {
			return errors.New("invalid decorator handler type")
		}

		inputChan, exists := c.channels[h.inputs[0]]
		if !exists {
			return errors.New(ErrChanNotFound)
		}
		outputChan, exists := c.channels[h.outputs[0]]
		if !exists {
			return errors.New(ErrChanNotFound)
		}

		return fn(ctx, inputChan, outputChan)

	case "multiplexer":
		fn, ok := h.fn.(func(ctx context.Context, inputs []chan string, output chan string) error)
		if !ok {
			return errors.New("invalid multiplexer handler type")
		}

		inputChans := make([]chan string, len(h.inputs))
		for i, input := range h.inputs {
			ch, exists := c.channels[input]
			if !exists {
				return errors.New(ErrChanNotFound)
			}
			inputChans[i] = ch
		}

		outputChan, exists := c.channels[h.outputs[0]]
		if !exists {
			return errors.New(ErrChanNotFound)
		}

		return fn(ctx, inputChans, outputChan)

	case "separator":
		fn, ok := h.fn.(func(ctx context.Context, input chan string, outputs []chan string) error)
		if !ok {
			return errors.New("invalid separator handler type")
		}

		inputChan, exists := c.channels[h.inputs[0]]
		if !exists {
			return errors.New(ErrChanNotFound)
		}

		outputChans := make([]chan string, len(h.outputs))
		for i, output := range h.outputs {
			ch, exists := c.channels[output]
			if !exists {
				return errors.New(ErrChanNotFound)
			}
			outputChans[i] = ch
		}

		return fn(ctx, inputChan, outputChans)

	default:
		return errors.New("unknown handler type")
	}
}

func (c *conveyerImpl) Send(input string, data string) error {
	ch, err := c.getChannel(input)
	if err != nil {
		return err
	}

	c.mutex.RLock()
	stopped := c.stopped
	full := len(ch) == cap(ch)
	c.mutex.RUnlock()

	if stopped {
		return errors.New(ErrConveyerStopped)
	}

	if full {
		return errors.New(ErrChannelFull)
	}

	ch <- data
	return nil
}

func (c *conveyerImpl) Recv(output string) (string, error) {
	ch, err := c.getChannel(output)
	if err != nil {
		return ErrUndefined, err
	}

	data, ok := <-ch
	if !ok {
		return ErrUndefined, nil
	}

	return data, nil
}

func (c *conveyerImpl) Stop() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.stopped {
		return
	}

	c.stopped = true

	for _, ch := range c.channels {
		close(ch)
	}
}