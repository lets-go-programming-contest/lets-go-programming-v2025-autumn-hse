package conveyer

import (
	"context"
	"sync"
)

type handlerType int

const (
	hDecorator handlerType = iota
	hMultiplexer
	hSeparator
)

type handler struct {
	kind handlerType

	fnDecorator   func(context.Context, chan string, chan string) error
	fnMultiplexer func(context.Context, []chan string, chan string) error
	fnSeparator   func(context.Context, chan string, []chan string) error

	inputIDs  []string
	outputIDs []string
}

type conveyerImpl struct {
	size     int
	mu       sync.RWMutex
	chans    map[string]chan string
	handlers []handler
	started  bool
}
