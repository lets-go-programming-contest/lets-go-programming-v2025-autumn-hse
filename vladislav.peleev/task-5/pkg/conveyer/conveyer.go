package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const recvValueOnClosedChannel = "undefined"

var (
	errCannotRegisterAfterRun = errors.New("cannot register after Run")
	errAlreadyStarted         = errors.New("already started")
	errChanNotFound           = errors.New("chan not found")
)

type Conveyer interface {
	RegisterDecorator(
		handlerFunc func(context.Context, chan string, chan string) error,
		input, output string,
	) error

	RegisterMultiplexer(
		handlerFunc func(context.Context, []chan string, chan string) error,
		inputs []string,
		output string,
	) error

	RegisterSeparator(
		handlerFunc func(context.Context, chan string, []chan string) error,
		input string,
		outputs []string,
	) error

	Run(ctx context.Context) error
	Send(input, data string) error
	Recv(output string) (string, error)
}

type conveyerImpl struct {
	mu       sync.RWMutex
	channels map[string]chan string
	handlers []func(context.Context) error
	started  bool
	bufSize  int
}

func New(size int) *conveyerImpl {
	return &conveyerImpl{
		mu:       sync.RWMutex{},
		channels: make(map[string]chan string),
		handlers: nil,
		started:  false,
		bufSize:  size,
	}
}

func (c *conveyerImpl) ensureChannel(name string) chan string {
	if channel, exists := c.channels[name]; exists {
		return channel
	}

	channel := make(chan string, c.bufSize)
	c.channels[name] = channel

	return channel
}

func (c *conveyerImpl) RegisterDecorator(
	handlerFunc func(context.Context, chan string, chan string) error,
	input, output string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.started {
		return errCannotRegisterAfterRun
	}

	inChannel := c.ensureChannel(input)
	outChannel := c.ensureChannel(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return handlerFunc(ctx, inChannel, outChannel)
	})

	return nil
}

func (c *conveyerImpl) RegisterSeparator(
	handlerFunc func(context.Context, chan string, []chan string) error,
	input string, outputs []string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.started {
		return errCannotRegisterAfterRun
	}

	inChannel := c.ensureChannel(input)

	outChannels := make([]chan string, len(outputs))
	for i, name := range outputs {
		outChannels[i] = c.ensureChannel(name)
	}

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return handlerFunc(ctx, inChannel, outChannels)
	})

	return nil
}

func (c *conveyerImpl) RegisterMultiplexer(
	handlerFunc func(context.Context, []chan string, chan string) error,
	inputs []string, output string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.started {
		return errCannotRegisterAfterRun
	}

	inChannels := make([]chan string, len(inputs))
	for i, name := range inputs {
		inChannels[i] = c.ensureChannel(name)
	}

	outChannel := c.ensureChannel(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return handlerFunc(ctx, inChannels, outChannel)
	})

	return nil
}

func (c *conveyerImpl) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.started {
		c.mu.Unlock()

		return errAlreadyStarted
	}

	c.started = true
	c.mu.Unlock()

	group, groupCtx := errgroup.WithContext(ctx)

	for _, handler := range c.handlers {
		h := handler

		group.Go(func() error {
			return h(groupCtx)
		})
	}

	err := group.Wait()

	c.close()

	if err != nil {
		return fmt.Errorf("errgroup wait: %w", err)
	}

	return nil
}

func (c *conveyerImpl) close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, channel := range c.channels {
		func(ch chan string) {
			defer func() { _ = recover() }()
			close(ch)
		}(channel)
	}
}

func (c *conveyerImpl) Send(input, data string) error {
	c.mu.RLock()
	channel, exists := c.channels[input]
	c.mu.RUnlock()

	if !exists {
		return errChanNotFound
	}

	channel <- data

	return nil
}

func (c *conveyerImpl) Recv(output string) (string, error) {
	c.mu.RLock()
	channel, exists := c.channels[output]
	c.mu.RUnlock()

	if !exists {
		return "", errChanNotFound
	}

	data, ok := <-channel
	if !ok {
		return RecvValueOnClosedChannel, nil
	}

	return data, nil
}
