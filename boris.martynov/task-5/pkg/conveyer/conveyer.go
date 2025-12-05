package conveyer

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

var ErrChanNotFound = errors.New("chan not found")

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

func (c *conveyor) RegisterDecorator(
	fnDecorator func(ctx context.Context, input chan string, output chan string) error,
	inputID string,
	outputID string,
) error {
	inputCh := c.ensure(inputID)
	outputCh := c.ensure(outputID)

	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return fnDecorator(ctx, inputCh, outputCh)
	})

	return nil
}

func (c *conveyor) RegisterMultiplexer(
	fnMultiplexer func(ctx context.Context, inputs []chan string, output chan string) error,
	inputIDs []string,
	outputID string,
) error {
	inputChans := make([]chan string, len(inputIDs))
	for i, id := range inputIDs {
		inputChans[i] = c.ensure(id)
	}

	outputCh := c.ensure(outputID)

	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return fnMultiplexer(ctx, inputChans, outputCh)
	})

	return nil
}

func (c *conveyor) RegisterSeparator(
	fnSeparator func(ctx context.Context, input chan string, outputs []chan string) error,
	inputID string,
	outputIDs []string,
) error {
	inputCh := c.ensure(inputID)

	outputChans := make([]chan string, len(outputIDs))
	for i, id := range outputIDs {
		outputChans[i] = c.ensure(id)
	}

	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return fnSeparator(ctx, inputCh, outputChans)
	})

	return nil
}

func (c *conveyor) Run(ctx context.Context) error {
	defer c.close()

	group, gCtx := errgroup.WithContext(ctx)

	for _, task := range c.tasks {
		t := task

		group.Go(func() error { return t(gCtx) })
	}

	err := group.Wait()
	if err != nil {
		return fmt.Errorf("error groyp return: %w", err)
	}

	return nil
}

func (c *conveyor) Send(channelID string, data string) error {
	ch, exists := c.channels[channelID]
	if !exists {
		return ErrChanNotFound
	}

	ch <- data

	return nil
}

func (c *conveyor) Recv(channelID string) (string, error) {
	ch, exists := c.channels[channelID]
	if !exists {
		return "", ErrChanNotFound
	}

	data, ok := <-ch
	if !ok {
		return "undefined", nil
	}

	return data, nil
}

func (c *conveyor) ensure(name string) chan string {
	if ch, exists := c.channels[name]; exists {
		return ch
	}

	ch := make(chan string, c.size)
	c.channels[name] = ch

	return ch
}

func (c *conveyor) close() {
	for _, ch := range c.channels {
		close(ch)
	}
}
