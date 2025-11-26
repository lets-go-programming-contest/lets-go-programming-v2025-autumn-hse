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
	ErrChannelMissing = errors.New("chan not found")
	ErrAlreadyRunning = errors.New("conveyer already running")
)

func New(size int) *Conveyer {
	return &Conveyer{
		channels: make(map[string]chan string),
		size:     size,
	}
}

func (c *Conveyer) get(name string) (chan string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ch, ok := c.channels[name]
	return ch, ok
}

func (c *Conveyer) getOrCreate(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()
	if ch, ok := c.channels[name]; ok {
		return ch
	}
	ch := make(chan string, c.size)
	c.channels[name] = ch
	return ch
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

	eg, _ := errgroup.WithContext(ctx)
	for _, t := range c.tasks {
		tc := t
		eg.Go(func() error { return c.exec(ctx, tc) })
	}
	return eg.Wait()
}

func (c *Conveyer) Send(name, data string) error {
	ch, ok := c.get(name)
	if !ok {
		return ErrChannelMissing
	}
	if len(ch) == cap(ch) {
		return errors.New("channel full")
	}
	ch <- data
	return nil
}

func (c *Conveyer) Recv(name string) (string, error) {
	ch, ok := c.get(name)
	if !ok {
		return "undefined", ErrChannelMissing
	}
	val, ok := <-ch
	if !ok {
		return "undefined", nil
	}
	return val, nil
}

func (c *Conveyer) exec(ctx context.Context, t task) error {
	if len(t.inputs) == 0 || len(t.outputs) == 0 {
		return ErrChannelMissing
	}

	switch t.kind {
	case "decorator":
		return t.fn.(func(context.Context, chan string, chan string) error)(ctx, c.getOrCreate(t.inputs[0]), c.getOrCreate(t.outputs[0]))
	case "multiplexer":
		ins := make([]chan string, len(t.inputs))
		for i, n := range t.inputs {
			ins[i] = c.getOrCreate(n)
		}
		return t.fn.(func(context.Context, []chan string, chan string) error)(ctx, ins, c.getOrCreate(t.outputs[0]))
	case "separator":
		outs := make([]chan string, len(t.outputs))
		for i, n := range t.outputs {
			outs[i] = c.getOrCreate(n)
		}
		return t.fn.(func(context.Context, chan string, []chan string) error)(ctx, c.getOrCreate(t.inputs[0]), outs)
	}
	return errors.New("unknown task type")
}

func (c *Conveyer) RegisterDecorator(fn func(context.Context, chan string, chan string) error, input, output string) {
	c.getOrCreate(input)
	c.getOrCreate(output)
	c.tasks = append(c.tasks, task{"decorator", fn, []string{input}, []string{output}})
}

func (c *Conveyer) RegisterMultiplexer(fn func(context.Context, []chan string, chan string) error, inputs []string, output string) {
	for _, n := range inputs {
		c.getOrCreate(n)
	}
	c.getOrCreate(output)
	c.tasks = append(c.tasks, task{"multiplexer", fn, inputs, []string{output}})
}

func (c *Conveyer) RegisterSeparator(fn func(context.Context, chan string, []chan string) error, input string, outputs []string) {
	c.getOrCreate(input)
	for _, n := range outputs {
		c.getOrCreate(n)
	}
	c.tasks = append(c.tasks, task{"separator", fn, []string{input}, outputs})
}
