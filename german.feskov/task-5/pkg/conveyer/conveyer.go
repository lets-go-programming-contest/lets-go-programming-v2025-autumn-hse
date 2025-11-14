package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

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
	input    map[string]chan string
	output   map[string]chan string
	handlers []handler
	mu       sync.Mutex
}

var ErrChanNotFound = errors.New("chan no found")

func New(size int) *DefaultConveyer {
	//nolint:exhaustruct // sync.Mutex dont need to initialize
	return &DefaultConveyer{
		size:     size,
		input:    make(map[string]chan string),
		output:   make(map[string]chan string),
		handlers: make([]handler, 0),
	}
}

func (c *DefaultConveyer) Run(ctx context.Context) error {
	errGroup, errGroupCtx := errgroup.WithContext(ctx)

	for _, handler := range c.handlers {
		h := handler

		errGroup.Go(func() error {
			err := h.run(errGroupCtx)

			return err
		})
	}

	err := errGroup.Wait()

	c.mu.Lock()
	defer c.mu.Unlock()

	for key, ch := range c.input {
		if _, ok := c.output[key]; !ok {
			close(ch)
		}
	}

	for _, ch := range c.output {
		close(ch)
	}

	//nolint:wrapcheck // errgorup.Wait return errors from handlers
	return err
}

func (c *DefaultConveyer) Send(input string, data string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	channel, ok := c.input[input]
	if !ok {
		return fmt.Errorf("str %q: %w", input, ErrChanNotFound)
	}
	channel <- data

	return nil
}

func (c *DefaultConveyer) Recv(output string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	//nolint:varnamelen // ok is classic name
	channel, ok := c.output[output]
	if !ok {
		return "", fmt.Errorf("str %q: %w", output, ErrChanNotFound)
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
	channel, ok := c.input[key]
	if !ok {
		channel, ok = c.output[key]
		if !ok {
			channel = make(chan string, c.size)
			c.output[key] = channel
		}

		c.input[key] = channel
	}

	return channel
}
