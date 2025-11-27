package conveyer

import (
	"context"
	"errors"
	"sync"
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
	ch, err := c.getChannel(input)
	if err != nil {
		return err
	}

	if len(ch) == cap(ch) {
		return ErrChannelFull
	}
	ch <- data

	return nil
}

func (c *Conveyer) Recv(output string) (string, error) {
	ch, err := c.getChannel(output)

	if err != nil {
		return undefined, err
	}

	data, ok := <-ch
	if !ok {
		return undefined, nil
	}

	return data, nil
}

func (c *Conveyer) closeAllChannels() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for name, ch := range c.channels {
		close(ch)
		delete(c.channels, name)
	}
}

func (c *Conveyer) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	errorChan := make(chan error, 1)

	for _, handler := range c.handlers {
		wg.Add(1)

		go func(h handlerConfig) {
			defer wg.Done()

			if err := c.runHandler(ctx, h); err != nil {
				select {
				case errorChan <- err:
				default:
				}
			}
		}(handler)
	}

	select {
	case err := <-errorChan:
		c.closeAllChannels()
		wg.Wait()
		return err

	case <-ctx.Done():
		c.closeAllChannels()
		wg.Wait()
		return ctx.Err()
	}
}

func (c *Conveyer) runHandler(ctx context.Context, h handlerConfig) error {
	switch h.kind {
	case "decorator":
		fn := h.fn.(func(context.Context, chan string, chan string) error)
		input := c.getOrCreateChannel(h.inputs[0])
		output := c.getOrCreateChannel(h.outputs[0])
		return fn(ctx, input, output)

	case "multiplexer":
		fn := h.fn.(func(context.Context, []chan string, chan string) error)
		inputs := make([]chan string, len(h.inputs))
		for i, name := range h.inputs {
			inputs[i] = c.getOrCreateChannel(name)
		}
		output := c.getOrCreateChannel(h.outputs[0])
		return fn(ctx, inputs, output)

	case "separator":
		fn := h.fn.(func(context.Context, chan string, []chan string) error)
		input := c.getOrCreateChannel(h.inputs[0])
		outputs := make([]chan string, len(h.outputs))
		for i, name := range h.outputs {
			outputs[i] = c.getOrCreateChannel(name)
		}
		return fn(ctx, input, outputs)
	}
	return nil
}

func (c *Conveyer) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.getOrCreateChannel(input)
	c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handlerConfig{
		kind:    "decorator",
		fn:      fn,
		inputs:  []string{input},
		outputs: []string{output},
	})
}

func (c *Conveyer) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, inputName := range inputs {
		c.getOrCreateChannel(inputName)
	}
	c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handlerConfig{
		kind:    "multiplexer",
		fn:      fn,
		inputs:  inputs,
		outputs: []string{output},
	})
}

func (c *Conveyer) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.getOrCreateChannel(input)
	for _, outputName := range outputs {
		c.getOrCreateChannel(outputName)
	}

	c.handlers = append(c.handlers, handlerConfig{
		kind:    "separator",
		fn:      fn,
		inputs:  []string{input},
		outputs: outputs,
	})
}
