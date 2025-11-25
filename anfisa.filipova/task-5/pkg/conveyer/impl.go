package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

const undefined = "undefined"

var (
	errChanNotFound = errors.New("chan not found")
	errChanIsFull   = errors.New("chan is full")
	errChanIsEmpty  = errors.New("chan is empty")
)

func (c *conveyerImpl) Send(input string, data string) error {
	c.m.RLock()
	defer c.m.RUnlock()
	ch, exists := c.channels[input]
	if !exists {
		return errChanNotFound
	}
	select {
	case ch <- data:
		return nil
	default:
		return errChanIsFull
	}
}

func (c *conveyerImpl) Recv(output string) (string, error) {
	c.m.RLock()
	defer c.m.RUnlock()
	ch, exists := c.channels[output]
	if !exists {
		return "", errChanNotFound
	}
	select {
	case data, ok := <-ch:
		if !ok {
			return undefined, nil
		}
		return data, nil
	default: //empty, but open
		return "", errChanIsEmpty
	}
}

func (c *conveyerImpl) Run(ctx context.Context) error {
	c.m.Lock()
	defer c.m.Unlock()

	var wg sync.WaitGroup
	errorCh := make(chan error, len(c.handlers))
	for _, handler := range c.handlers {
		wg.Add(1)
		go func(h handlerConfig) {
			defer wg.Done()
			if err := c.runHandler(ctx, h); err != nil {
				errorCh <- err
			}
		}(handler)
	}
	go func() {
		wg.Wait()
		close(errorCh)
	}()

	select {
	case err, ok := <-errorCh:
		if ok {
			c.closeAll()
			return err
		}
		return nil
	case <-ctx.Done():
		c.closeAll()
		return ctx.Err()
	}
}

func (c *conveyerImpl) runHandler(ctx context.Context, config handlerConfig) error {
	switch config.handlerType {
	case handlerDecorator:
		fn, ok := config.fn.(func(ctx context.Context, input chan string, output chan string) error)
		if !ok {
			return fmt.Errorf("invalid decorator function type")
		}

		return fn(ctx, c.channels[config.input], c.channels[config.output])

	case handlerMultiplexer:
		fn, ok := config.fn.(func(ctx context.Context, inputs []chan string, output chan string) error)
		if !ok {
			return fmt.Errorf("invalid multiplexer function type")
		}

		inputs := make([]chan string, len(config.inputs))
		for i, name := range config.inputs {
			inputs[i] = c.channels[name]
		}

		return fn(ctx, inputs, c.channels[config.output])

	case handlerSeparator:
		fn, ok := config.fn.(func(ctx context.Context, input chan string, outputs []chan string) error)
		if !ok {
			return fmt.Errorf("invalid separator function type")
		}

		outputs := make([]chan string, len(config.outputs))
		for i, name := range config.outputs {
			outputs[i] = c.channels[name]
		}

		return fn(ctx, c.channels[config.input], outputs)

	default:
		return fmt.Errorf("unknown handler type")
	}
}

func (c *conveyerImpl) closeAll() {
	c.m.Lock()
	defer c.m.Unlock()
	for name, ch := range c.channels {
		close(ch)
		// select {
		// case _, ok := <-ch:
		// 	if ok {
		// 		close(ch)
		// 	}
		// default:
		// 	close(ch)
		// }
		delete(c.channels, name)
	}
	c.handlers = nil
}

func (c *conveyerImpl) getChannel(name string) chan string {
	if ch, exists := c.channels[name]; exists {
		return ch
	}
	ch := make(chan string, c.channelSize)
	c.channels[name] = ch
	return ch
}
