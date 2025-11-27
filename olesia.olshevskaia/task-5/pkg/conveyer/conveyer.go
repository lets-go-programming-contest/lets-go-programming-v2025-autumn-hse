package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	ErrChannelNotFound        = errors.New("chan not found")
	ErrChannelFull            = errors.New("channel full")
	ErrConveyerAlreadyRunning = errors.New("conveyer already running")
)

const (
	undefined = "undefined"
)

type Conveyer struct {
	mu       sync.Mutex
	channels map[string]chan string
	size     int
	handlers []handlerConfig
}

type handlerConfig struct {
	kind    string
	fn      interface{}
	inputs  []string
	outputs []string
}

func New(size int) *Conveyer {
	return &Conveyer{
		channels: make(map[string]chan string),
		size:     size,
		handlers: []handlerConfig{},
		mu:       sync.Mutex{},
	}
}

func (c *Conveyer) getOrCreateChannel(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ch, exists := c.channels[name]; exists {
		return ch
	}

	ch := make(chan string, c.size)
	c.channels[name] = ch

	return ch
}

func (c *Conveyer) getChannel(name string) (chan string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ch, exists := c.channels[name]
	if !exists {
		return nil, ErrChannelNotFound
	}

	return ch, nil
}

func (c *Conveyer) Send(input string, data string) error {
	channel, err := c.getChannel(input)
	if err != nil {
		return err
	}

	if len(channel) == cap(channel) {
		return ErrChannelFull
	}

	channel <- data

	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	channel, err := c.getChannel(output)
	if err != nil {
		return undefined, err
	}

	data, ok := <-channel
	if !ok {
		return undefined, nil
	}

	return data, nil
}

func (c *Conveyer) Run(ctx context.Context) error {
	defer func() {
		c.mu.Lock()
		for _, ch := range c.channels {
			close(ch)
		}
		c.mu.Unlock()
	}()

	errGroup, gctx := errgroup.WithContext(ctx)

	for _, taskItem := range c.handlers {
		ti := taskItem

		errGroup.Go(func() error {
			return c.runHandler(gctx, ti)
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("failed: %w", err)
	}

	return nil
}

func (c *Conveyer) runHandler(ctx context.Context, handler handlerConfig) error {
	switch handler.kind {
	case "decorator":
		handlerFunc, ok := handler.fn.(func(context.Context, chan string, chan string) error)
		if !ok {
			return errors.New("invalid type of handler function for one input/one output")
		}

		inputChannel := c.getOrCreateChannel(handler.inputs[0])
		outputChannel := c.getOrCreateChannel(handler.outputs[0])

		return handlerFunc(ctx, inputChannel, outputChannel)

	case "multiplexer":
		handlerFunc, ok := handler.fn.(func(context.Context, []chan string, chan string) error)
		if !ok {
			return errors.New("invalid handler function type for multiple inputs/one output")
		}

		inputChannels := make([]chan string, len(handler.inputs))

		for i, name := range handler.inputs {
			inputChannels[i] = c.getOrCreateChannel(name)
		}

		outputChannel := c.getOrCreateChannel(handler.outputs[0])

		return handlerFunc(ctx, inputChannels, outputChannel)

	case "separator":
		handlerFunc, ok := handler.fn.(func(context.Context, chan string, []chan string) error)
		if !ok {
			return errors.New("invalid handler function type for one input/multiple outputs")
		}

		inputChannel := c.getOrCreateChannel(handler.inputs[0])
		outputChannels := make([]chan string, len(handler.outputs))

		for i, name := range handler.outputs {
			outputChannels[i] = c.getOrCreateChannel(name)
		}

		return handlerFunc(ctx, inputChannel, outputChannels)
	}

	return nil
}

func (c *Conveyer) RegisterDecorator(
	handlerFunc func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.getOrCreateChannel(input)
	c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handlerConfig{
		kind:    "decorator",
		fn:      handlerFunc,
		inputs:  []string{input},
		outputs: []string{output},
	})
}

func (c *Conveyer) RegisterMultiplexer(
	handlerFunc func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, inputName := range inputs {
		c.getOrCreateChannel(inputName)
	}

	c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handlerConfig{
		kind:    "multiplexer",
		fn:      handlerFunc,
		inputs:  inputs,
		outputs: []string{output},
	})
}

func (c *Conveyer) RegisterSeparator(
	handlerFunc func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.getOrCreateChannel(input)

	for _, outputName := range outputs {
		c.getOrCreateChannel(outputName)
	}

	c.handlers = append(c.handlers, handlerConfig{
		kind:    "separator",
		fn:      handlerFunc,
		inputs:  []string{input},
		outputs: outputs,
	})
}
