package conveyer

import (
	"context"
	"errors"
	"fmt"
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
	mutex    sync.RWMutex
}

type task struct {
	name        string
	function    any
	inputChans  []string
	outputChans []string
}

func New(size int) *Conveyer {
	return &Conveyer{
		channels: make(map[string]chan string),
		size:     size,
		tasks:    make([]task, 0),
		mutex:    sync.RWMutex{},
	}
}

func (conveyer *Conveyer) get(name string) (chan string, bool) {
	conveyer.mutex.RLock()
	defer conveyer.mutex.RUnlock()

	channel, found := conveyer.channels[name]

	return channel, found
}

func (conveyer *Conveyer) getOrCreate(name string) chan string {
	conveyer.mutex.RLock()
	defer conveyer.mutex.RUnlock()

	channel, found := conveyer.channels[name]
	if !found {
		channel = make(chan string, conveyer.size)
		conveyer.channels[name] = channel
	}

	return channel
}

func (conveyer *Conveyer) Run(ctx context.Context) error {
	defer func() {
		conveyer.mutex.RLock()
		for _, ch := range conveyer.channels {
			close(ch)
		}
		conveyer.mutex.RUnlock()
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
	if err != nil {
		return fmt.Errorf("error while executing tasks: %w", err)
	}

	return nil
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
