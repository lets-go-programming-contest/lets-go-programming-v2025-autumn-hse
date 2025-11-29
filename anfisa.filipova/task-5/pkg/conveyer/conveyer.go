package conveyer

import (
	"sync"
)

type handlerType int

const (
	handlerDecorator handlerType = iota
	handlerMultiplexer
	handlerSeparator
)

type handlerConfig struct {
	handlerType handlerType
	fn          interface{}
	inputs      []string
	outputs     []string
}

type Conveyer struct {
	mutex       sync.RWMutex
	channels    map[string]chan string
	handlers    []handlerConfig
	channelSize int
}

func New(size int) *Conveyer {
	return &Conveyer{
		mutex:       sync.RWMutex{},
		channels:    make(map[string]chan string),
		handlers:    make([]handlerConfig, 0),
		channelSize: size,
	}
}
