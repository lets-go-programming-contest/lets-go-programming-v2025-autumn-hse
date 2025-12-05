package conveyer

import (
	"context"
	"errors"
	"sync"
)

var ErrChanNotFound = errors.New("chan not found")

type handlerFunc func(ctx context.Context) error

type Conveyer struct {
	size int

	chans map[string]chan string

	mu sync.RWMutex

	handlers []handlerFunc
}

func New(size int) *Conveyer {
	return &Conveyer{
		size:     size,
		chans:    make(map[string]chan string),
		mu:       sync.RWMutex{},
		handlers: make([]handlerFunc, 0),
	}
}

func (c *Conveyer) getOrCreate(ident string) chan string {
	c.mu.Lock()
	channelRef, ok := c.chans[ident]

	if !ok {
		channelRef = make(chan string, c.size)
		c.chans[ident] = channelRef
	}
	c.mu.Unlock()

	return channelRef
}

func (c *Conveyer) RegisterDecorator(
	handlerFunction func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	income := c.getOrCreate(input)
	outcome := c.getOrCreate(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return handlerFunction(ctx, income, outcome)
	})
}

func (c *Conveyer) RegisterMultiplexer(
	handlerFunction func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	income := make([]chan string, 0, len(inputs))
	for _, id := range inputs {
		income = append(income, c.getOrCreate(id))
	}

	outcome := c.getOrCreate(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return handlerFunction(ctx, income, outcome)
	})
}

func (c *Conveyer) RegisterSeparator(
	handlerFunction func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	income := c.getOrCreate(input)
	outcome := make([]chan string, 0, len(outputs))

	for _, id := range outputs {
		outcome = append(outcome, c.getOrCreate(id))
	}

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return handlerFunction(ctx, income, outcome)
	})
}

func (c *Conveyer) Run(ctx context.Context) error {
	var waitGroup sync.WaitGroup

	errorCh := make(chan error, 1)

	for _, handler := range c.handlers {
		waitGroup.Add(1)

		handlerCopy := handler

		go func() {
			defer waitGroup.Done()

			err := handlerCopy(ctx)
			if err != nil {
				errorCh <- err
			}
		}()
	}

	go func() {
		waitGroup.Wait()
		errorCh <- nil
	}()

	err := <-errorCh

	c.mu.RLock()
	for _, channelRef := range c.chans {
		close(channelRef)
	}
	c.mu.RUnlock()

	return err
}

func (c *Conveyer) Send(id string, data string) error {
	c.mu.RLock()
	channelRef, ok := c.chans[id]
	c.mu.RUnlock()

	if !ok {
		return ErrChanNotFound
	}

	channelRef <- data

	return nil
}

func (c *Conveyer) Recv(id string) (string, error) {
	c.mu.RLock()
	channelRef, isOpen := c.chans[id]
	c.mu.RUnlock()

	if !isOpen {
		return "", ErrChanNotFound
	}

	v, isOpen := <-channelRef
	if !isOpen {
		return "undefined", nil
	}

	return v, nil
}
