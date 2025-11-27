package conveyer

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

const undefined = "undefined"

var (
	errChanNotFound       = errors.New("chan not found")
	errChanIsFull         = errors.New("chan is full")
	errUnknownHandlerType = errors.New("unknown handler type")
)

func (c *conveyerImpl) Send(input string, data string) error {
	channel, exists := c.channels[input]
	if !exists {
		return errChanNotFound
	}

	select {
	case channel <- data:
		return nil
	default:
		return errChanIsFull
	}
}

func (c *conveyerImpl) Recv(output string) (string, error) {
	channel, exists := c.channels[output]
	if !exists {
		return undefined, errChanNotFound
	}

	data, ok := <-channel
	if !ok {
		return undefined, nil
	}

	return data, nil

}

func (c *conveyerImpl) Run(ctx context.Context) error {
	defer c.closeAll()
	errGroup, ctx := errgroup.WithContext(ctx)
	for _, handler := range c.handlers {
		errGroup.Go(func() error {
			return c.runHandler(ctx, handler)
		})
	}

	err := errGroup.Wait()
	if err != nil {
		return fmt.Errorf("conveyer error: %w", err)
	}

	return nil
}

func (c *conveyerImpl) runHandler(ctx context.Context, handler handlerConfig) error {
	inputChannels := make([]chan string, len(handler.inputs))
	for i, name := range handler.inputs {
		inputChannels[i] = c.channels[name]
	}

	outputChannels := make([]chan string, len(handler.outputs))
	for i, name := range handler.outputs {
		outputChannels[i] = c.channels[name]
	}

	switch handler.handlerType {
	case handlerDecorator:
		return c.runDecorator(ctx, handler, inputChannels, outputChannels)
	case handlerMultiplexer:
		return c.runMultiplexer(ctx, handler, inputChannels, outputChannels)
	case handlerSeparator:
		return c.runSeparator(ctx, handler, inputChannels, outputChannels)
	default:
		return errUnknownHandlerType
	}
}

func (c *conveyerImpl) closeAll() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, channel := range c.channels {
		close(channel)
	}

	c.handlers = nil
}

func (c *conveyerImpl) getOrCreateChannel(name string) chan string {
	if channel, exists := c.channels[name]; exists {
		return channel
	}

	channel := make(chan string, c.channelSize)
	c.channels[name] = channel

	return channel
}
