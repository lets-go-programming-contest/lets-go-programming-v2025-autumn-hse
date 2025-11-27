package conveyer

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func New(size int) *conveyerImpl {
	return &conveyerImpl{
		size:  size,
		chans: make(map[string]chan string),
	}
}

func (c *conveyerImpl) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.started {
		c.mu.Unlock()
		return fmt.Errorf("already started")
	}
	c.started = true
	c.mu.Unlock()

	g, ctx := errgroup.WithContext(ctx)

	for _, h := range c.handlers {
		h := h

		g.Go(func() error {
			c.mu.RLock()
			inChans := make([]chan string, len(h.inputIDs))
			outChans := make([]chan string, len(h.outputIDs))
			for i, id := range h.inputIDs {
				inChans[i] = c.chans[id]
			}
			for i, id := range h.outputIDs {
				outChans[i] = c.chans[id]
			}
			c.mu.RUnlock()

			switch h.kind {
			case hDecorator:
				return h.fnDecorator(ctx, inChans[0], outChans[0])
			case hMultiplexer:
				return h.fnMultiplexer(ctx, inChans, outChans[0])
			case hSeparator:
				return h.fnSeparator(ctx, inChans[0], outChans)
			default:
				return fmt.Errorf("unknown handler type")
			}
		})
	}

	err := g.Wait()

	c.mu.Lock()
	for _, ch := range c.chans {
		close(ch)
	}
	c.mu.Unlock()

	return err
}

func (c *conveyerImpl) Send(id string, data string) error {
	c.mu.RLock()
	ch, ok := c.chans[id]
	c.mu.RUnlock()

	if !ok {
		return errors.New("chan not found")
	}

	ch <- data
	return nil
}

func (c *conveyerImpl) Recv(id string) (string, error) {
	c.mu.RLock()
	ch, ok := c.chans[id]
	c.mu.RUnlock()

	if !ok {
		return "", errors.New("chan not found")
	}

	v, ok := <-ch
	if !ok {
		return "undefined", nil
	}

	return v, nil
}
