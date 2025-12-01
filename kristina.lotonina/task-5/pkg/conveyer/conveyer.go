package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

const (
	ErrChanNotFound    = "chan not found"
	ErrUndefined       = "undefined"
	ErrConveyerStopped = "conveyer stopped"
)

type conveyerImpl struct {
	size     int
	channels map[string]chan string
	handlers []handler
	mutex    sync.RWMutex
	wg       sync.WaitGroup
	cancel   context.CancelFunc
}

type handler struct {
	handlerType string
	fn          interface{}
	inputs      []string
	outputs     []string
}

func New(size int) *conveyerImpl {
	return &conveyerImpl{
		size:     size,
		channels: make(map[string]chan string),
		handlers: make([]handler, 0),
	}
}

func (c *conveyerImpl) getOrCreateChannel(id string) chan string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if ch, exists := c.channels[id]; exists {
		return ch
	}

	ch := make(chan string, c.size)
	c.channels[id] = ch
	return ch
}

func (c *conveyerImpl) getChannel(id string) (chan string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if ch, exists := c.channels[id]; exists {
		return ch, nil
	}

	return nil, errors.New(ErrChanNotFound)
}

func (c *conveyerImpl) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	inputChan := c.getOrCreateChannel(input)
	outputChan := c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handler{
		handlerType: "decorator",
		fn:          fn,
		inputs:      []string{input},
		outputs:     []string{output},
	})
}

func (c *conveyerImpl) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	inputChans := make([]chan string, len(inputs))
	for i, input := range inputs {
		inputChans[i] = c.getOrCreateChannel(input)
	}
	outputChan := c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handler{
		handlerType: "multiplexer",
		fn:          fn,
		inputs:      inputs,
		outputs:     []string{output},
	})
}

func (c *conveyerImpl) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	inputChan := c.getOrCreateChannel(input)
	outputChans := make([]chan string, len(outputs))
	for i, output := range outputs {
		outputChans[i] = c.getOrCreateChannel(output)
	}

	c.handlers = append(c.handlers, handler{
		handlerType: "separator",
		fn:          fn,
		inputs:      []string{input},
		outputs:     outputs,
	})
}

func (c *conveyerImpl) Run(ctx context.Context) error {
	ctx, c.cancel = context.WithCancel(ctx)
	defer c.cancel()

	for _, h := range c.handlers {
		c.wg.Add(1)
		go func(handler handler) {
			defer c.wg.Done()
			c.runHandler(ctx, handler)
		}(h)
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (c *conveyerImpl) runHandler(ctx context.Context, h handler) {
	switch h.handlerType {
	case "decorator":
		if fn, ok := h.fn.(func(ctx context.Context, input chan string, output chan string) error); ok {
			inputChan, _ := c.getChannel(h.inputs[0])
			outputChan, _ := c.getChannel(h.outputs[0])
			fn(ctx, inputChan, outputChan)
		}
	case "multiplexer":
		if fn, ok := h.fn.(func(ctx context.Context, inputs []chan string, output chan string) error); ok {
			inputChans := make([]chan string, len(h.inputs))
			for i, input := range h.inputs {
				inputChans[i], _ = c.getChannel(input)
			}
			outputChan, _ := c.getChannel(h.outputs[0])
			fn(ctx, inputChans, outputChan)
		}
	case "separator":
		if fn, ok := h.fn.(func(ctx context.Context, input chan string, outputs []chan string) error); ok {
			inputChan, _ := c.getChannel(h.inputs[0])
			outputChans := make([]chan string, len(h.outputs))
			for i, output := range h.outputs {
				outputChans[i], _ = c.getChannel(output)
			}
			fn(ctx, inputChan, outputChans)
		}
	}
}

func (c *conveyerImpl) Send(input string, data string) error {
	ch, err := c.getChannel(input)
	if err != nil {
		return err
	}

	select {
	case ch <- data:
		return nil
	default:
		return errors.New("channel is full")
	}
}

func (c *conveyerImpl) Recv(output string) (string, error) {
	ch, err := c.getChannel(output)
	if err != nil {
		return "", err
	}

	select {
	case data, ok := <-ch:
		if !ok {
			return ErrUndefined, nil
		}
		return data, nil
	default:
		return "", errors.New("no data available")
	}
}

func (c *conveyerImpl) Stop() {
	if c.cancel != nil {
		c.cancel()
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	for id, ch := range c.channels {
		close(ch)
		delete(c.channels, id)
	}

	c.wg.Wait()
}