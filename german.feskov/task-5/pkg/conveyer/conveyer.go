package conveyer

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	undefined = "undefined"
)

type handler interface {
	run(context.Context) error
}

type DefaultConveyer struct {
	size     int
	input    map[string]chan string
	output   map[string]chan string
	handlers []handler
	mu       sync.Mutex
}

func New(size int) *DefaultConveyer {
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

	return err
}

func (c *DefaultConveyer) Send(input string, data string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	ch, ok := c.input[input]
	if !ok {
		return fmt.Errorf("chan not found %q", input)
	}
	ch <- data
	return nil
}

func (c *DefaultConveyer) Recv(output string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ch, ok := c.output[output]
	if !ok {
		return "", fmt.Errorf("chan not found %q", output)
	}
	res, ok := <-ch
	if !ok {
		return undefined, nil
	}
	return res, nil
}

func (c *DefaultConveyer) RegisterDecorator(fn DecoratorFunc, input string, output string) {
	inCh := c.createChanIfNotExists(input)
	outCh := c.createChanIfNotExists(output)
	c.handlers = append(c.handlers, &decorator{fn: fn, input: inCh, output: outCh})
}

func (c *DefaultConveyer) RegisterMultiplexer(fn MultiplexerFunc, inputs []string, output string) {
	inChs := make([]chan string, len(inputs))
	for i, input := range inputs {
		inChs[i] = c.createChanIfNotExists(input)
	}
	outCh := c.createChanIfNotExists(output)
	c.handlers = append(c.handlers, &multiplexer{fn: fn, input: inChs, output: outCh})
}

func (c *DefaultConveyer) RegisterSeparator(fn SeparatorFunc, input string, outputs []string) {
	inCh := c.createChanIfNotExists(input)
	outChs := make([]chan string, len(outputs))
	for i, output := range outputs {
		outChs[i] = c.createChanIfNotExists(output)
	}
	c.handlers = append(c.handlers, &separator{fn: fn, input: inCh, output: outChs})
}

func (c *DefaultConveyer) createChanIfNotExists(key string) chan string {
	ch, ok := c.input[key]
	if !ok {
		ch, ok = c.output[key]
		if !ok {
			ch = make(chan string, c.size)
			c.output[key] = ch
		}
		c.input[key] = ch
	}

	return ch
}
