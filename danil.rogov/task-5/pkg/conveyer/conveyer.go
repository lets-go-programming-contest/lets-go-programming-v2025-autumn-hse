package conveyer

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	ErrChanNotFound           = errors.New("chan not found")
	ErrChanIsFull             = errors.New("chan is full")
	ErrUnknownTask            = errors.New("unknown task type")
	ErrInvalidDecoratorType   = errors.New("invalid decorator function type")
	ErrInvalidMultiplexerType = errors.New("invalid multiplexer function type")
	ErrInvalidSeparatorType   = errors.New("invalid separator function type")
)

const undefined = "undefined"

type Conveyer struct {
	channels map[string]chan string
	size     int
	tasks    []task
	mutex    sync.Mutex
}

type task struct {
	name    string
	fn      any
	inputs  []string
	outputs []string
}

func New(size int) *Conveyer {
	return &Conveyer{
		channels: make(map[string]chan string),
		size:     size,
		tasks:    make([]task, 0),
		mutex:    sync.Mutex{},
	}
}

func (conveyer *Conveyer) get(name string) (chan string, bool) {
	conveyer.mutex.Lock()
	defer conveyer.mutex.Unlock()

	channel, found := conveyer.channels[name]

	return channel, found
}

func (conveyer *Conveyer) getOrCreate(name string) chan string {
	conveyer.mutex.Lock()
	defer conveyer.mutex.Unlock()

	channel, found := conveyer.channels[name]
	if found {
		return channel
	}

	channel = make(chan string, conveyer.size)
	conveyer.channels[name] = channel

	return channel
}

func (conveyer *Conveyer) Run(ctx context.Context) error {
	defer func() {
		conveyer.mutex.Lock()
		for _, ch := range conveyer.channels {
			close(ch)
		}
		conveyer.mutex.Unlock()
	}()

	errGroup, _ := errgroup.WithContext(ctx)

	for _, task := range conveyer.tasks {
		errGroup.Go(func() error {
			switch task.name {
			case "decorator":
				return conveyer.executeDecorator(ctx, task)
			case "multiplexer":
				return conveyer.executeMultiplexer(ctx, task)
			case "separator":
				return conveyer.executeSeparator(ctx, task)
			default:
				return ErrUnknownTask
			}
		})
	}

	err := errGroup.Wait()

	return err
}

func (conveyer *Conveyer) Send(name, data string) error {
	channel, found := conveyer.get(name)
	if !found {
		return ErrChanNotFound
	}

	if len(channel) == cap(channel) {
		return ErrChanIsFull
	}

	channel <- data

	return nil
}

func (conveyer *Conveyer) Recv(name string) (string, error) {
	channel, found := conveyer.get(name)
	if !found {
		return undefined, ErrChanNotFound
	}

	val, ok := <-channel
	if !ok {
		return undefined, nil
	}

	return val, nil
}

func (conveyer *Conveyer) executeDecorator(ctx context.Context, task task) error {
	if len(task.inputs) == 0 || len(task.outputs) == 0 {
		return ErrChanNotFound
	}

	decoratorFunc, ok := task.fn.(func(context.Context, chan string, chan string) error)
	if !ok {
		return ErrInvalidDecoratorType
	}

	inputChan := conveyer.getOrCreate(task.inputs[0])
	outputChan := conveyer.getOrCreate(task.outputs[0])

	return decoratorFunc(ctx, inputChan, outputChan)
}

func (conveyer *Conveyer) executeMultiplexer(ctx context.Context, task task) error {
	multiplexerFunc, ok := task.fn.(func(context.Context, []chan string, chan string) error)
	if !ok {
		return ErrInvalidMultiplexerType
	}

	inputChans := make([]chan string, len(task.inputs))

	for index, name := range task.inputs {
		inputChans[index] = conveyer.getOrCreate(name)
	}

	outputChan := conveyer.getOrCreate(task.outputs[0])

	return multiplexerFunc(ctx, inputChans, outputChan)
}

func (conveyer *Conveyer) executeSeparator(ctx context.Context, taskItem task) error {
	separatorFunc, ok := taskItem.fn.(func(context.Context, chan string, []chan string) error)
	if !ok {
		return ErrInvalidSeparatorType
	}

	outputChans := make([]chan string, len(taskItem.outputs))

	for index, name := range taskItem.outputs {
		outputChans[index] = conveyer.getOrCreate(name)
	}

	inputChan := conveyer.getOrCreate(taskItem.inputs[0])

	return separatorFunc(ctx, inputChan, outputChans)
}
