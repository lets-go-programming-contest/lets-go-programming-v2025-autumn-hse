package conveyer

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Conveyer struct {
	channels map[string]chan string
	size     int
	tasks    []task
	mu       sync.Mutex
}

type task struct {
	kind    string
	fn      interface{}
	inputs  []string
	outputs []string
}

var (
	ErrChanNotFound    = errors.New("chan not found")
	ErrChannelFull     = errors.New("channel full")
	ErrUnknownTask     = errors.New("unknown task type")
	ErrInvalidTaskFunc = errors.New("invalid task function")
)

var undefined = "undefined"

func New(size int) *Conveyer {
	return &Conveyer{
		channels: make(map[string]chan string),
		size:     size,
		tasks:    []task{},
		mu:       sync.Mutex{},
	}
}

func (c *Conveyer) get(name string) (chan string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	channel, found := c.channels[name]

	return channel, found
}

func (c *Conveyer) getOrCreate(name string) chan string {
	c.mu.Lock()
	defer c.mu.Unlock()

	channel, found := c.channels[name]
	if found {
		return channel
	}

	channel = make(chan string, c.size)
	c.channels[name] = channel

	return channel
}

func (c *Conveyer) Run(ctx context.Context) error {
	defer func() {
		c.mu.Lock()
		for _, ch := range c.channels {
			close(ch)
		}
		c.mu.Unlock()
	}()

	errGroup, _ := errgroup.WithContext(ctx)
	for _, taskItem := range c.tasks {
		errGroup.Go(func() error {
			return c.exec(ctx, taskItem)
		})
	}

	err := errGroup.Wait()
	if err != nil {
		return fmt.Errorf("run tasks failed: %w", err)
	}

	return nil
}

func (c *Conveyer) Send(name, data string) error {
	channel, found := c.get(name)
	if !found {
		return ErrChanNotFound
	}

	if len(channel) == cap(channel) {
		return ErrChannelFull
	}

	channel <- data

	return nil
}

func (c *Conveyer) Recv(name string) (string, error) {
	channel, found := c.get(name)
	if !found {
		return undefined, ErrChanNotFound
	}

	val, ok := <-channel
	if !ok {
		return undefined, nil
	}

	return val, nil
}

func (c *Conveyer) exec(ctx context.Context, taskItem task) error {
	if len(taskItem.inputs) == 0 || len(taskItem.outputs) == 0 {
		return ErrChanNotFound
	}

	switch taskItem.kind {
	case "decorator":
		decoratorFn, ok := taskItem.fn.(func(context.Context, chan string, chan string) error)
		if !ok {
			return ErrInvalidTaskFunc
		}

		inputChannel := c.getOrCreate(taskItem.inputs[0])
		outputChannel := c.getOrCreate(taskItem.outputs[0])

		return decoratorFn(ctx, inputChannel, outputChannel)

	case "multiplexer":
		multiplexerFn, ok := taskItem.fn.(func(context.Context, []chan string, chan string) error)
		if !ok {
			return ErrInvalidTaskFunc
		}

		ins := make([]chan string, len(taskItem.inputs))

		for index, name := range taskItem.inputs {
			ins[index] = c.getOrCreate(name)
		}

		outputChannel := c.getOrCreate(taskItem.outputs[0])

		return multiplexerFn(ctx, ins, outputChannel)

	case "separator":
		separatorFn, ok := taskItem.fn.(func(context.Context, chan string, []chan string) error)
		if !ok {
			return ErrInvalidTaskFunc
		}

		outs := make([]chan string, len(taskItem.outputs))

		for index, name := range taskItem.outputs {
			outs[index] = c.getOrCreate(name)
		}

		inputChannel := c.getOrCreate(taskItem.inputs[0])

		return separatorFn(ctx, inputChannel, outs)
	}

	return ErrUnknownTask
}

func (c *Conveyer) RegisterDecorator(
	decFunc func(context.Context, chan string, chan string) error,
	input string,
	output string,
) {
	c.getOrCreate(input)
	c.getOrCreate(output)

	c.tasks = append(c.tasks, task{
		kind:    "decorator",
		fn:      decFunc,
		inputs:  []string{input},
		outputs: []string{output},
	})
}

func (c *Conveyer) RegisterMultiplexer(
	decFunc func(context.Context, []chan string, chan string) error,
	inputs []string,
	output string,
) {
	for _, name := range inputs {
		c.getOrCreate(name)
	}

	c.getOrCreate(output)

	c.tasks = append(c.tasks, task{
		kind:    "multiplexer",
		fn:      decFunc,
		inputs:  inputs,
		outputs: []string{output},
	})
}

func (c *Conveyer) RegisterSeparator(
	decFunc func(context.Context, chan string, []chan string) error,
	input string,
	outputs []string,
) {
	c.getOrCreate(input)

	for _, name := range outputs {
		c.getOrCreate(name)
	}

	c.tasks = append(c.tasks, task{
		kind:    "separator",
		fn:      decFunc,
		inputs:  []string{input},
		outputs: outputs,
	})
}
