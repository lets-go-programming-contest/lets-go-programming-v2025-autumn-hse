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
	eg, ctx := errgroup.WithContext(ctx)

	for _, h := range c.handlers {
		handler := h

		eg.Go(func() error {
			switch handler.typ {
			case handlerTypeDecorator:
				fn := handler.fn.(func(ctx context.Context, input chan string, output chan string) error)
				inputCh := c.getChannel(handler.inputs[0])
				outputCh := c.getChannel(handler.outputs[0])

				return fn(ctx, inputCh, outputCh)
			case handlerTypeMultiplexer:
				fn := handler.fn.(func(ctx context.Context, inputs []chan string, output chan string) error)
				inputChs := make([]chan string, len(handler.inputs))

				for i, input := range handler.inputs {
					inputChs[i] = c.getChannel(input)
				}

				outputCh := c.getChannel(handler.outputs[0])

				return fn(ctx, inputChs, outputCh)
			case handlerTypeSeparator:
				fn := handler.fn.(func(ctx context.Context, input chan string, outputs []chan string) error)
				inputCh := c.getChannel(handler.inputs[0])
				outputChs := make([]chan string, len(handler.outputs))

				for i, output := range handler.outputs {
					outputChs[i] = c.getChannel(output)
				}

				return fn(ctx, inputCh, outputChs)
			default:
				return fmt.Errorf("unknown handler type: %s", handler.typ)
			}
		})
	}

	err := eg.Wait()
	c.closeAllChannels()

	return err
}

func (c *conveyer) Send(input string, data string) error {
	channel := c.getChannel(input)

	select {
	case channel <- data:
		return nil
	default:
		return errors.New("channel is full")
	}
}

func (c *conveyer) Recv(output string) (string, error) {
	channel := c.getChannel(output)

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
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.getChannel(input)
	c.getChannel(output)

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
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, input := range inputs {
		c.getChannel(input)
	}
	c.getChannel(output)

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
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.getChannel(input)
	for _, output := range outputs {
		c.getChannel(output)
	}

	c.handlers = append(c.handlers, handler{
		typ:     handlerTypeSeparator,
		fn:      fn,
		inputs:  []string{input},
		outputs: outputs,
	})
}

func New(size int) *conveyer {
	return &conveyer{
		channels: make(map[string]chan string),
		size:     size,
		handlers: make([]handler, 0),
	}
}

func (c *conveyer) getChannel(name string) chan string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	channel, exists := c.channels[name]
	if !exists {
		channel = make(chan string, c.size)
		c.channels[name] = channel
	}

	return channel
}
