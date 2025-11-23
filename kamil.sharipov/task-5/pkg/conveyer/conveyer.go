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

type handler struct {
	typ     handlerType
	fn      interface{}
	inputs  []string
	outputs []string
}

var ErrChanNotFound = errors.New("chan not found")

func (c *conveyer) Run(ctx context.Context) error {
	c.mutex.Lock()
	if c.running {
		c.mutex.Unlock()
		return errors.New("conveyer is already running")
	}
	c.running = true
	c.mutex.Unlock()

	defer func() {
		c.closeAllChannels()
		c.mutex.Lock()
		c.running = false
		c.mutex.Unlock()
	}()

	// Создаем все каналы перед запуском
	c.mutex.Lock()
	for _, h := range c.handlers {
		for _, input := range h.inputs {
			c.getOrCreateChannel(input)
		}
		for _, output := range h.outputs {
			c.getOrCreateChannel(output)
		}
	}
	c.mutex.Unlock()

	eg, ctx := errgroup.WithContext(ctx)

	for _, h := range c.handlers {
		handler := h
		eg.Go(func() error {
			return c.runHandler(ctx, handler)
		})
	}

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("conveyer finished with error: %w", err)
	}

	return nil
}

func (c *conveyer) runHandler(ctx context.Context, h handler) error {
	switch h.typ {
	case handlerTypeDecorator:
		fn := h.fn.(func(ctx context.Context, input chan string, output chan string) error)
		inputCh, err := c.getChannel(h.inputs[0])
		if err != nil {
			return err
		}
		outputCh, err := c.getChannel(h.outputs[0])
		if err != nil {
			return err
		}
		return fn(ctx, inputCh, outputCh)

	case handlerTypeMultiplexer:
		fn := h.fn.(func(ctx context.Context, inputs []chan string, output chan string) error)
		inputChs := make([]chan string, len(h.inputs))
		for i, input := range h.inputs {
			ch, err := c.getChannel(input)
			if err != nil {
				return err
			}
			inputChs[i] = ch
		}
		outputCh, err := c.getChannel(h.outputs[0])
		if err != nil {
			return err
		}
		return fn(ctx, inputChs, outputCh)

	case handlerTypeSeparator:
		fn := h.fn.(func(ctx context.Context, input chan string, outputs []chan string) error)
		inputCh, err := c.getChannel(h.inputs[0])
		if err != nil {
			return err
		}
		outputChs := make([]chan string, len(h.outputs))
		for i, output := range h.outputs {
			ch, err := c.getChannel(output)
			if err != nil {
				return err
			}
			outputChs[i] = ch
		}
		return fn(ctx, inputCh, outputChs)

	default:
		return fmt.Errorf("unknown handler type: %s", h.typ)
	}
}

func (c *conveyer) Send(input string, data string) error {
	channel, err := c.getChannel(input)

	if err != nil {
		return fmt.Errorf("%w: %s", ErrChanNotFound, input)
	}

	select {
	case channel <- data:
		return nil
	default:
		return errors.New("channel is full")
	}
}

func (c *conveyer) Recv(output string) (string, error) {
	channel, err := c.getChannel(output)

	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrChanNotFound, output)
	}

	select {
	case data, ok := <-channel:
		if !ok {
			return undefined, nil
		}
		return data, nil
	default:
		return undefined, nil
	}
}

func (c *conveyer) closeAllChannels() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for name, ch := range c.channels {
		close(ch)
		delete(c.channels, name)
	}
}

func (c *conveyer) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.handlers = append(c.handlers, handler{
		typ:     handlerTypeDecorator,
		fn:      fn,
		inputs:  []string{input},
		outputs: []string{output},
	})
}

func (c *conveyer) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	c.handlers = append(c.handlers, handler{
		typ:     handlerTypeMultiplexer,
		fn:      fn,
		inputs:  inputs,
		outputs: []string{output},
	})
}

func (c *conveyer) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.handlers = append(c.handlers, handler{
		typ:     handlerTypeSeparator,
		fn:      fn,
		inputs:  []string{input},
		outputs: outputs,
	})
}

func (c *conveyer) getChannel(name string) (chan string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	channel, exists := c.channels[name]
	if !exists {
		return nil, fmt.Errorf("%w: %s", ErrChanNotFound, name)
	}
	return channel, nil
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
