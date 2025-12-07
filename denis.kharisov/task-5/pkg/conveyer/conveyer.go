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

func (c *conveyor) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	inputID, outputID string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputCh := c.ensure(inputID)
	outputCh := c.ensure(outputID)

	c.tasks = append(c.tasks, func(ctx context.Context) error {
		return fn(ctx, inputCh, outputCh)
	})
	return nil
}

func (c *conveyor) RegisterMultiplexer(
	fn func(ctx context.Context, inputs []chan string, output chan string) error,
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
		return fn(ctx, inputChans, outputCh)
	})
	return nil
}

func (c *conveyor) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
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
		return fn(ctx, inputCh, outputChans)
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
	ch, exists := c.channels[channelID]
	c.mu.Unlock()

	if !exists {
		return errChanNotFound
	}
	ch <- data
	return nil
}

func (c *conveyor) Recv(channelID string) (string, error) {
	c.mu.Lock()
	ch, exists := c.channels[channelID]
	c.mu.Unlock()

	if !exists {
		return "", errChanNotFound
	}

	data, ok := <-ch
	if !ok {
		return undefined, nil
	}
	return data, nil
}

func (c *conveyor) closeAll() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, ch := range c.channels {
		close(ch)
	}
}