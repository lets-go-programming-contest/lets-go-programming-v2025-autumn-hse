package conveyer

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

type conveyor struct {
	channels map[string]chan string
	tasks    []func(context.Context) error
	size     int
}

func New(size int) *conveyor {
	return &conveyor{
		channels: make(map[string]chan string),
		tasks:    make([]func(context.Context) error, 0),
		size:     size,
	}
}

func (c *conveyor) ensure(name string) chan string {
	if ch, exists := c.channels[name]; exists {
		return ch
	}
	ch := make(chan string, c.size)
	c.channels[name] = ch
	return ch
}

func (c *conveyor) RegisterDecorator(fn func(ctx context.Context, input chan string, output chan string) error, inputID string, outputID string) error {
	inputCh := c.ensure(inputID)
	outputCh := c.ensure(outputID)
	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return fn(ctx, inputCh, outputCh)
	})
	return nil
}

func (c *conveyor) RegisterMultiplexer(fn func(ctx context.Context, inputs []chan string, output chan string) error, inputIDs []string, outputID string) error {
	inputChans := make([]chan string, len(inputIDs))
	for i, id := range inputIDs {
		inputChans[i] = c.ensure(id)
	}
	outputCh := c.ensure(outputID)
	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return fn(ctx, inputChans, outputCh)
	})
	return nil
}

func (c *conveyor) RegisterSeparator(fn func(ctx context.Context, input chan string, outputs []chan string) error, inputID string, outputIDs []string) error {
	inputCh := c.ensure(inputID)
	outputChans := make([]chan string, len(outputIDs))
	for i, id := range outputIDs {
		outputChans[i] = c.ensure(id)
	}
	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return fn(ctx, inputCh, outputChans)
	})
	return nil
}

func (c *conveyor) Run(ctx context.Context) error {
	g, gCtx := errgroup.WithContext(ctx)
	for _, task := range c.tasks {
		t := task
		g.Go(func() error {
			return t(gCtx)
		})
	}
	return g.Wait()
}

func (c *conveyor) Send(channelID string, data string) error {
	ch, exists := c.channels[channelID]
	if !exists {
		return fmt.Errorf("chan not found")
	}
	ch <- data
	return nil
}

func (c *conveyor) Recv(channelID string) (string, error) {
	ch, exists := c.channels[channelID]
	if !exists {
		return "", fmt.Errorf("chan not found")
	}
	data, ok := <-ch
	if !ok {
		return "undefined", nil
	}
	return data, nil
}
