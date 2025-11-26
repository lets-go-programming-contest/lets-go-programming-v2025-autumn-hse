package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type conveyer struct {
	channels map[string]chan string
	size     int
	handlers []handler
	mutex    sync.Mutex
	running  bool
}

type handlerType string

const (
	handlerTypeDecorator   handlerType = "decorator"
	handlerTypeMultiplexer handlerType = "multiplexer"
	handlerTypeSeparator   handlerType = "separator"
)

const (
	undefined = "undefined"
)

var (
	ErrChanNotFound           = errors.New("chan not found")
	ErrConveyerAlreadyRunning = errors.New("conveyer is already running")
	ErrChannelFull            = errors.New("channel is full")
	ErrUnknownHandlerType     = errors.New("unknown handler type")
)

type handler struct {
	typ     handlerType
	fn      interface{}
	inputs  []string
	outputs []string
}

func (c *conveyer) Run(ctx context.Context) error {
	c.mutex.Lock()
	if c.running {
		c.mutex.Unlock()
		return ErrConveyerAlreadyRunning
	}

	c.running = true
	c.mutex.Unlock()

	defer func() {
		c.mutex.Lock()
		c.running = false
		c.closeAllChannels()
		c.mutex.Unlock()
	}()

	errGroup, ctx := errgroup.WithContext(ctx)

	for _, handler := range c.handlers {
		errGroup.Go(func() error {
			return c.runHandler(ctx, handler)
		})
	}

	return fmt.Errorf("conveyer execution failed: %w", errGroup.Wait())
}

func (c *conveyer) runHandler(ctx context.Context, h handler) error {
	switch h.typ {
	case handlerTypeDecorator:
		fn := h.fn.(func(ctx context.Context, input chan string, output chan string) error)

		inputCh, exists := c.getChannel(h.inputs[0])
		if !exists {
			return ErrChanNotFound
		}

		outputCh, exists := c.getChannel(h.outputs[0])
		if !exists {
			return ErrChanNotFound
		}

		return fn(ctx, inputCh, outputCh)

	case handlerTypeMultiplexer:
		fn := h.fn.(func(ctx context.Context, inputs []chan string, output chan string) error)
		inputChs := make([]chan string, len(h.inputs))

		for index, input := range h.inputs {
			ch, exists := c.getChannel(input)
			if !exists {
				return ErrChanNotFound
			}

			inputChs[index] = ch
		}

		outputCh, exists := c.getChannel(h.outputs[0])
		if !exists {
			return ErrChanNotFound
		}

		return fn(ctx, inputChs, outputCh)

	case handlerTypeSeparator:
		fn := h.fn.(func(ctx context.Context, input chan string, outputs []chan string) error)

		inputCh, exists := c.getChannel(h.inputs[0])
		if !exists {
			return ErrChanNotFound
		}

		outputChs := make([]chan string, len(h.outputs))

		for index, output := range h.outputs {
			ch, exists := c.getChannel(output)
			if !exists {
				return ErrChanNotFound
			}

			outputChs[index] = ch
		}

		return fn(ctx, inputCh, outputChs)

	default:
		return fmt.Errorf("%w: %s", ErrUnknownHandlerType, h.typ)
	}
}

func (c *conveyer) Send(input string, data string) error {
	c.mutex.Lock()
	channel, exists := c.channels[input]
	c.mutex.Unlock()

	if !exists {
		return ErrChanNotFound
	}

	select {
	case channel <- data:
		return nil
	default:
		return ErrChannelFull
	}
}

func (c *conveyer) Recv(output string) (string, error) {
	c.mutex.Lock()
	channel, exists := c.channels[output]
	c.mutex.Unlock()

	if !exists {
		return undefined, ErrChanNotFound
	}

	data, ok := <-channel
	if !ok {
		return undefined, nil
	}

	return data, nil
}

func (c *conveyer) RegisterDecorator(
	decoratorFn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.getOrCreateChannel(input)
	c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handler{
		typ:     handlerTypeDecorator,
		fn:      decoratorFn,
		inputs:  []string{input},
		outputs: []string{output},
	})
}

func (c *conveyer) RegisterMultiplexer(
	multiplexerFn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, input := range inputs {
		c.getOrCreateChannel(input)
	}

	c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handler{
		typ:     handlerTypeMultiplexer,
		fn:      multiplexerFn,
		inputs:  inputs,
		outputs: []string{output},
	})
}

func (c *conveyer) RegisterSeparator(
	separatorFn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.getOrCreateChannel(input)

	for _, output := range outputs {
		c.getOrCreateChannel(output)
	}

	c.handlers = append(c.handlers, handler{
		typ:     handlerTypeSeparator,
		fn:      separatorFn,
		inputs:  []string{input},
		outputs: outputs,
	})
}

func (c *conveyer) getChannel(name string) (chan string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	ch, ok := c.channels[name]

	return ch, ok
}

func (c *conveyer) getOrCreateChannel(name string) chan string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	channel, exists := c.channels[name]
	if !exists {
		channel = make(chan string, c.size)
		c.channels[name] = channel
	}

	return channel
}

func (c *conveyer) closeAllChannels() {
	for _, ch := range c.channels {
		close(ch)
	}
}

func New(size int) *conveyer {
	return &conveyer{
		channels: make(map[string]chan string),
		size:     size,
		handlers: make([]handler, 0),
	}
}
