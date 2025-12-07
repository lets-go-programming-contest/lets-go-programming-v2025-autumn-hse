package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

const undefined = "undefined"

var errChanNotFound = errors.New("chan not found")

type conveyor struct {
	mu       sync.Mutex
	channels map[string]chan string
	tasks    []func(context.Context) error
	size     int
}

func New(size int) *conveyor {
	return &conveyor{
		mu:       sync.Mutex{},
		channels: make(map[string]chan string),
		tasks:    make([]func(context.Context) error, 0),
		size:     size,
	}
}

func (c *conveyor) ensure(name string) chan string {
	if channel, exists := c.channels[name]; exists {
		return channel
	}

	channel := make(chan string, c.size)
	c.channels[name] = channel

	return channel
}

func (c *conveyor) RegisterDecorator(
	taskFn func(ctx context.Context, input chan string, output chan string) error,
	inputID, outputID string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputCh := c.ensure(inputID)
	outputCh := c.ensure(outputID)

	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return taskFn(ctx, inputCh, outputCh)
	})

	return nil
}

func (c *conveyor) RegisterMultiplexer(
	taskFn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputIDs []string,
	outputID string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputChans := make([]chan string, len(inputIDs))
	for i, id := range inputIDs {
		inputChans[i] = c.ensure(id)
	}

	outputCh := c.ensure(outputID)

	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return taskFn(ctx, inputChans, outputCh)
	})

	return nil
}

func (c *conveyor) RegisterSeparator(
	taskFn func(ctx context.Context, input chan string, outputs []chan string) error,
	inputID string,
	outputIDs []string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputCh := c.ensure(inputID)
	outputChans := make([]chan string, len(outputIDs))

	for i, id := range outputIDs {
		outputChans[i] = c.ensure(id)
	}

	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return taskFn(ctx, inputCh, outputChans)
	})

	return nil
}

func (c *conveyor) Run(ctx context.Context) error {
	c.mu.Lock()
	tasks := make([]func(context.Context) error, len(c.tasks))
	copy(tasks, c.tasks)
	c.mu.Unlock()

	group, gCtx := errgroup.WithContext(ctx)

	for _, task := range tasks {
		t := task
		group.Go(func() error {
			return t(gCtx)
		})
	}

	if err := group.Wait(); err != nil {
		c.closeAll()

		return fmt.Errorf("conveyor failed: %w", err)
	}

	c.closeAll()

	return nil
}

func (c *conveyor) Send(channelID string, data string) error {
	c.mu.Lock()
	channel, exists := c.channels[channelID]
	c.mu.Unlock()

	if !exists {
		return errChanNotFound
	}
	channel <- data

	return nil
}

func (c *conveyor) Recv(channelID string) (string, error) {
	c.mu.Lock()
	channel, exists := c.channels[channelID]
	c.mu.Unlock()

	if !exists {
		return "", errChanNotFound
	}

	data, ok := <-channel
	if !ok {
		return undefined, nil
	}

	return data, nil
}

func (c *conveyor) closeAll() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, channel := range c.channels {
		close(channel)
	}
}
