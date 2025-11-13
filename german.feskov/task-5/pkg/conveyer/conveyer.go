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
		errGroup.Go(func() error {
			return handler.run(errGroupCtx)
		})
	}

	return errGroup.Wait()
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
	inCh := createChanIfNotExists(c.input, input, c.size)
	outCh := createChanIfNotExists(c.output, output, c.size)
	c.handlers = append(c.handlers, &decorator{fn: fn, input: inCh, output: outCh})
}

func (c *DefaultConveyer) RegisterMultiplexer(fn MultiplexerFunc, inputs []string, output string) {
	inChs := make([]chan string, len(inputs))
	for i, input := range inputs {
		inChs[i] = createChanIfNotExists(c.input, input, c.size)
	}
	outCh := createChanIfNotExists(c.output, output, c.size)
	c.handlers = append(c.handlers, &multiplexer{fn: fn, input: inChs, output: outCh})
}

func (c *DefaultConveyer) RegisterSeparator(fn SeparatorFunc, input string, outputs []string) {
	inCh := createChanIfNotExists(c.input, input, c.size)
	outChs := make([]chan string, len(outputs))
	for i, output := range outputs {
		outChs[i] = createChanIfNotExists(c.output, output, c.size)
	}
	c.handlers = append(c.handlers, &separator{fn: fn, input: inCh, output: outChs})
}

func createChanIfNotExists(container map[string]chan string, key string, size int) chan string {
	ch, ok := container[key]
	if !ok {
		ch = make(chan string, size)
	}
	return ch
}
