package conveyer

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

var (
	ErrAlreadyStarted     = errors.New("already started")
	ErrUnknownHandlerType = errors.New("unknown handler type")
	ErrChanNotFound       = errors.New("chan not found")
)

func New(size int) *conveyerImpl {
	return &conveyerImpl{
		size:     size,
		chans:    make(map[string]chan string),
		handlers: []handler{},
		started:  false,
	}
}

func (c *conveyerImpl) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.started {
		c.mu.Unlock()

		return ErrAlreadyStarted
	}

	c.started = true
	c.mu.Unlock()

	group, ctx := errgroup.WithContext(ctx)

	for _, handl := range c.handlers {
		group.Go(func() error {
			c.mu.RLock()

			inChans := make([]chan string, len(handl.inputIDs))
			outChans := make([]chan string, len(handl.outputIDs))

			for i, id := range handl.inputIDs {
				inChans[i] = c.chans[id]
			}

			for i, id := range handl.outputIDs {
				outChans[i] = c.chans[id]
			}

			c.mu.RUnlock()

			switch handl.kind {
			case hDecorator:
				return handl.fnDecorator(ctx, inChans[0], outChans[0])
			case hMultiplexer:
				return handl.fnMultiplexer(ctx, inChans, outChans[0])
			case hSeparator:
				return handl.fnSeparator(ctx, inChans[0], outChans)
			default:
				return ErrUnknownHandlerType
			}
		})
	}

	err := group.Wait()

	c.mu.Lock()
	for _, ch := range c.chans {
		close(ch)
	}
	c.mu.Unlock()

	if err != nil {
		return fmt.Errorf("handler error: %w", err)
	}

	return nil
}

func (c *conveyerImpl) Send(id string, data string) error {
	c.mu.RLock()
	channel, okey := c.chans[id]
	c.mu.RUnlock()

	if !okey {
		return ErrChanNotFound
	}

	channel <- data

	return nil
}

func (c *conveyerImpl) Recv(id string) (string, error) {
	c.mu.RLock()
	channel, okey := c.chans[id]
	c.mu.RUnlock()

	if !okey {
		return "", ErrChanNotFound
	}

	v, okey := <-channel
	if !okey {
		return "undefined", nil
	}

	return v, nil
}
