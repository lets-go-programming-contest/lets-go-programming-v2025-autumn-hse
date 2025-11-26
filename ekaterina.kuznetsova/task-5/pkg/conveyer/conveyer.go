package conveyer

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Conveyer struct {
	channels map[string]chan string
	size     int
	tasks    []task
	mu       sync.Mutex
	active   bool
}

type task struct {
	kind    string
	fn      interface{}
	inputs  []string
	outputs []string
}

var (
	ErrChannelMissing = errors.New("channel missing")
	ErrAlreadyRunning = errors.New("conveyer already running")
)

func New(size int) *Conveyer {
	return &Conveyer{
		channels: make(map[string]chan string),
		size:     size,
		tasks:    make([]task, 0),
	}
}

func (c *Conveyer) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.active {
		c.mu.Unlock()
		return ErrAlreadyRunning
	}
	c.active = true
	c.mu.Unlock()

	defer func() {
		c.mu.Lock()
		c.active = false
		for _, ch := range c.channels {
			close(ch)
		}
		c.mu.Unlock()
	}()

	eg, ctx := errgroup.WithContext(ctx)
	for _, t := range c.tasks {
		taskCopy := t
		eg.Go(func() error {
			return c.executeTask(ctx, taskCopy)
		})
	}

	return eg.Wait()
}

func (c *Conveyer) Send(name, data string) error {
	c.mu.Lock()
	ch, ok := c.channels[name]
	c.mu.Unlock()
	if !ok {
		return ErrChannelMissing
	}

	select {
	case ch <- data:
		return nil
	default:
		return errors.New("channel full")
	}
}

func (c *Conveyer) Recv(name string) (string, error) {
	c.mu.Lock()
	ch, ok := c.channels[name]
	c.mu.Unlock()
	if !ok {
		return "undefined", ErrChannelMissing
	}

	val, ok := <-ch
	if !ok {
		return "undefined", nil
	}
	return val, nil
}

func (c *Conveyer) getOrCreate(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()
	ch, ok := c.channels[name]
	if !ok {
		ch = make(chan string, c.size)
		c.channels[name] = ch
	}
	return ch
}

func (c *Conveyer) executeTask(ctx context.Context, t task) error {
	switch t.kind {
	case "decorator":
		fn, ok := t.fn.(func(context.Context, chan string, chan string) error)
		if !ok {
			return errors.New("invalid decorator type")
		}
		return fn(ctx, c.getOrCreate(t.inputs[0]), c.getOrCreate(t.outputs[0]))
	case "multiplexer":
		fn, ok := t.fn.(func(context.Context, []chan string, chan string) error)
		if !ok {
			return errors.New("invalid multiplexer type")
		}
		ins := make([]chan string, len(t.inputs))
		for i, n := range t.inputs {
			ins[i] = c.getOrCreate(n)
		}
		return fn(ctx, ins, c.getOrCreate(t.outputs[0]))
	case "separator":
		fn, ok := t.fn.(func(context.Context, chan string, []chan string) error)
		if !ok {
			return errors.New("invalid separator type")
		}
		outs := make([]chan string, len(t.outputs))
		for i, n := range t.outputs {
			outs[i] = c.getOrCreate(n)
		}
		return fn(ctx, c.getOrCreate(t.inputs[0]), outs)
	}
	return errors.New("unknown task type")
}

func (c *Conveyer) RegisterDecorator(fn func(context.Context, chan string, chan string) error, input, output string) {
	c.getOrCreate(input)
	c.getOrCreate(output)
	c.tasks = append(c.tasks, task{
		kind:    "decorator",
		fn:      fn,
		inputs:  []string{input},
		outputs: []string{output},
	})
}

func (c *Conveyer) RegisterMultiplexer(fn func(context.Context, []chan string, chan string) error, inputs []string, output string) {
	for _, in := range inputs {
		c.getOrCreate(in)
	}
	c.getOrCreate(output)
	c.tasks = append(c.tasks, task{
		kind:    "multiplexer",
		fn:      fn,
		inputs:  inputs,
		outputs: []string{output},
	})
}

func (c *Conveyer) RegisterSeparator(fn func(context.Context, chan string, []chan string) error, input string, outputs []string) {
	c.getOrCreate(input)
	for _, out := range outputs {
		c.getOrCreate(out)
	}
	c.tasks = append(c.tasks, task{
		kind:    "separator",
		fn:      fn,
		inputs:  []string{input},
		outputs: outputs,
	})
}
