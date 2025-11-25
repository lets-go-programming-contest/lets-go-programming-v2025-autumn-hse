package conveyer

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

const (
	undefined = "undefined"
)

type handler interface {
	run(ctx context.Context) error
}

type DefaultConveyer struct {
	size     int
	channels map[string]chan string
	handlers []handler
}

var ErrChanNotFound = errors.New("chan not found")

func New(size int) *DefaultConveyer {
	return &DefaultConveyer{
		size:     size,
		channels: make(map[string]chan string),
		handlers: make([]handler, 0),
	}
}

func (c *DefaultConveyer) Run(ctx context.Context) error {
	defer c.close()

	errGroup, errGroupCtx := errgroup.WithContext(ctx)

	for _, handler := range c.handlers {
		h := handler

		errGroup.Go(func() error {
			if err := h.run(errGroupCtx); err != nil {
				return err
			}

			return nil
		})
	}

	err := errGroup.Wait()
	if err != nil {
		return fmt.Errorf("errgroup returned: %w", err)
	}

	return nil
}

func (c *DefaultConveyer) Send(input string, data string) error {
	channel, ok := c.channels[input]
	if !ok {
		return fmt.Errorf("send data into channel %q: %w", input, ErrChanNotFound)
	}
	channel <- data

	return nil
}

func (c *DefaultConveyer) Recv(output string) (string, error) {
	//nolint:varnamelen // ok is classic name
	channel, ok := c.channels[output]
	if !ok {
		return "", fmt.Errorf("receive data from channel %q: %w", output, ErrChanNotFound)
	}

	res, ok := <-channel
	if !ok {
		return undefined, nil
	}

	return res, nil
}

func (c *DefaultConveyer) RegisterDecorator(foo DecoratorFunc, input string, output string) {
	inCh := c.createChanIfNotExists(input)
	outCh := c.createChanIfNotExists(output)
	c.handlers = append(c.handlers, &decorator{fn: foo, input: inCh, output: outCh})
}

func (c *DefaultConveyer) RegisterMultiplexer(foo MultiplexerFunc, inputs []string, output string) {
	inChs := make([]chan string, len(inputs))
	for i, input := range inputs {
		inChs[i] = c.createChanIfNotExists(input)
	}

	outCh := c.createChanIfNotExists(output)
	c.handlers = append(c.handlers, &multiplexer{fn: foo, input: inChs, output: outCh})
}

func (c *DefaultConveyer) RegisterSeparator(foo SeparatorFunc, input string, outputs []string) {
	inCh := c.createChanIfNotExists(input)

	outChs := make([]chan string, len(outputs))
	for i, output := range outputs {
		outChs[i] = c.createChanIfNotExists(output)
	}

	c.handlers = append(c.handlers, &separator{fn: foo, input: inCh, output: outChs})
}

func (c *DefaultConveyer) createChanIfNotExists(key string) chan string {
	channel, ok := c.channels[key]
	if !ok {
		channel = make(chan string, c.size)
		c.channels[key] = channel
	}

	return channel
}

func (c *DefaultConveyer) close() {
	for _, ch := range c.channels {
		close(ch)
	}
}
