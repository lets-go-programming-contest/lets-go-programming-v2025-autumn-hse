package conveyer

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type conveyor struct {
	mu      sync.RWMutex
	inputs  map[string]chan string
	outputs map[string]chan string
	tasks   []func(context.Context) error
	size    int
}

func New(size int) *conveyor {
	return &conveyor{
		inputs:  make(map[string]chan string),
		outputs: make(map[string]chan string),
		tasks:   make([]func(context.Context) error, 0),
		size:    size,
	}
}

func (c *conveyor) ensureInput(name string) chan string {
	if ch, exists := c.inputs[name]; exists {
		return ch
	}
	ch := make(chan string, c.size)
	c.inputs[name] = ch
	return ch
}

func (c *conveyor) ensureOutput(name string) chan string {
	if ch, exists := c.outputs[name]; exists {
		return ch
	}
	ch := make(chan string, c.size)
	c.outputs[name] = ch
	return ch
}

func (c *conveyor) RegisterDecorator(
	fn func(ctx context.Context, input chan string, output chan string) error,
	inputID string,
	outputID string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputCh := c.ensureInput(inputID)
	outputCh := c.ensureOutput(outputID)

	task := func(ctx context.Context) error {
		return fn(ctx, inputCh, outputCh)
	}

	c.tasks = append(c.tasks, task)
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
		inputChans[i] = c.ensureInput(id)
	}
	outputCh := c.ensureOutput(outputID)

	task := func(ctx context.Context) error {
		return fn(ctx, inputChans, outputCh)
	}

	c.tasks = append(c.tasks, task)
	return nil
}

func (c *conveyor) RegisterSeparator(
	fn func(ctx context.Context, input chan string, outputs []chan string) error,
	inputID string,
	outputIDs []string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	inputCh := c.ensureInput(inputID)

	outputChans := make([]chan string, len(outputIDs))
	for i, id := range outputIDs {
		outputChans[i] = c.ensureOutput(id)
	}

	task := func(ctx context.Context) error {
		return fn(ctx, inputCh, outputChans)
	}

	c.tasks = append(c.tasks, task)
	return nil
}

func (c *conveyor) Run(ctx context.Context) error {
	c.mu.RLock()
	tasks := c.tasks
	c.mu.RUnlock()

	if len(tasks) == 0 {
		<-ctx.Done()
		return ctx.Err()
	}

	g, gCtx := errgroup.WithContext(ctx)

	for _, task := range tasks {
		t := task
		g.Go(func() error {
			return t(gCtx)
		})
	}

	return g.Wait()
}
func (c *conveyor) Send(inputID string, data string) error {
	c.mu.RLock()
	ch, exists := c.inputs[inputID]
	c.mu.RUnlock()

	if !exists {
		return fmt.Errorf("chan not found")
	}

	select {
	case ch <- data:
		return nil
	default:
		return fmt.Errorf("input channel '%s' is closed or full", inputID)
	}
}

func (c *conveyor) Recv(outputID string) (string, error) {
	c.mu.RLock()
	ch, exists := c.outputs[outputID]
	c.mu.RUnlock()

	if !exists {
		return "", fmt.Errorf("chan not found")
	}

	data, ok := <-ch
	if !ok {
		return "undefined", nil
	}
	return data, nil
}
