package conveyer

import (
	"context"
	"sync"
)

type Conveyer interface {
	RegisterDecorator(
		fn func(
			ctx context.Context,
			input chan string,
			output chan string,
		) error,
		input string,
		output string,
	)
	RegisterMultiplexer(
		fn func(
			ctx context.Context,
			inputs []chan string,
			output chan string,
		) error,
		inputs []string,
		output string,
	)
	RegisterSeparator(
		fn func(
			ctx context.Context,
			input chan string,
			outputs []chan string,
		) error,
		input string,
		outputs []string,
	)

	Run(ctx context.Context) error
	Send(input string, data string) error
	Recv(output string) (string, error)
}

type handlerType int

const (
	handlerDecorator handlerType = iota
	handlerMultiplexer
	handlerSeparator
)

type handlerConfig struct {
	handlerType handlerType
	fn          interface{}
	input       string
	inputs      []string
	output      string
	outputs     []string
}

type conveyerImpl struct {
	m           sync.RWMutex
	channels    map[string]chan string
	handlers    []handlerConfig
	channelSize int
}

func New(size int) Conveyer {
	return &conveyerImpl{
		channels:    make(map[string]chan string),
		channelSize: size,
	}
}
