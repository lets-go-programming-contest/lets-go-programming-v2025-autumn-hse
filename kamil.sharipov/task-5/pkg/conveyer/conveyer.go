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

var (
	ErrChanNotFound           = errors.New("chan not found")
	ErrConveyerAlreadyRunning = errors.New("conveyer is already running")
	ErrChannelFull            = errors.New("channel is full")
	ErrUnknownHandlerType     = errors.New("unknown handler type")
	ErrInvalidDecoratorType   = errors.New("invalid decorator function type")
	ErrInvalidMultiplexerType = errors.New("invalid multiplexer function type")
	ErrInvalidSeparatorType   = errors.New("invalid separator function type")
)

type handler struct {
	typ     handlerType
	fn      interface{}
	inputs  []string
	outputs []string
}

func (c *conveyer) Run(ctx context.Context) error {
	c.mutex.Lock()
	if c.running {
		c.mutex.Unlock()

		return ErrConveyerAlreadyRunning
	}

	c.running = true
	c.mutex.Unlock()

	defer func() {
		c.mutex.Lock()
		c.running = false
		c.closeAllChannels()
		c.mutex.Unlock()
	}()

	errGroup, ctx := errgroup.WithContext(ctx)

	for _, handler := range c.handlers {
		errGroup.Go(func() error {
			return c.runHandler(ctx, handler)
		})
	}

	if err := errGroup.Wait(); err != nil {
		return fmt.Errorf("conveyer execution failed: %w", err)
	}

	return nil
}

func (c *conveyer) Send(input string, data string) error {
	c.mutex.Lock()
	channel, exists := c.channels[input]
	c.mutex.Unlock()

	if !exists {
		return ErrChanNotFound
	}

	select {
	case channel <- data:
		return nil
	default:
		return ErrChannelFull
	}
}

func (c *conveyer) Recv(output string) (string, error) {
	c.mutex.Lock()
	channel, exists := c.channels[output]
	c.mutex.Unlock()

	if !exists {
		return undefined, ErrChanNotFound
	}

	data, ok := <-channel
	if !ok {
		return undefined, nil
	}

	return data, nil
}

func (c *conveyer) getChannel(name string) (chan string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	ch, ok := c.channels[name]

	return ch, ok
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

func (c *conveyer) closeAllChannels() {
	for _, ch := range c.channels {
		close(ch)
	}
}

func New(size int) *conveyer {
	return &conveyer{
		channels: make(map[string]chan string),
		size:     size,
		handlers: make([]handler, 0),
		mutex:    sync.Mutex{},
		running:  false,
	}
}
