package conveyer

import (
	"context"
	"errors"
	"sync"
)

const Undefined = "undefined"

type Conveyer interface {
	RegisterDecorator(fn func(context.Context, chan string, chan string) error, input, output string) error
	RegisterMultiplexer(fn func(context.Context, []chan string, chan string) error, inputs []string, output string) error
	RegisterSeparator(fn func(context.Context, chan string, []chan string) error, input string, outputs []string) error
	Run(ctx context.Context) error
	Send(input, data string) error
	Recv(output string) (string, error)
}

type conveyerImpl struct {
	mu sync.RWMutex

	channels map[string]chan string

	handlers []func(context.Context) error

	started bool
	stopped bool

	bufSize int
}

func New(size int) Conveyer {
	return &conveyerImpl{
		channels: make(map[string]chan string),
		bufSize:  size,
	}
}

func (c *conveyerImpl) ensureChannel(name string) chan string {
	if ch, ok := c.channels[name]; ok {
		return ch
	}
	ch := make(chan string, c.bufSize)
	c.channels[name] = ch
	return ch
}

func (c *conveyerImpl) RegisterDecorator(
	fn func(context.Context, chan string, chan string) error,
	input, output string,
) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.started {
		return errors.New("cannot register after Run")
	}

	inCh := c.ensureChannel(input)
	outCh := c.ensureChannel(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, inCh, outCh)
	})

	return nil
}

func (c *conveyerImpl) RegisterSeparator(
	fn func(context.Context, chan string, []chan string) error,
	input string, outputs []string,
) error {

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.started {
		return errors.New("cannot register after Run")
	}

	inCh := c.ensureChannel(input)
	var outChs []chan string
	for _, o := range outputs {
		outChs = append(outChs, c.ensureChannel(o))
	}

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, inCh, outChs)
	})

	return nil
}

func (c *conveyerImpl) RegisterMultiplexer(
	fn func(context.Context, []chan string, chan string) error,
	inputs []string, output string,
) error {

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.started {
		return errors.New("cannot register after Run")
	}

	var inChs []chan string
	for _, inp := range inputs {
		inChs = append(inChs, c.ensureChannel(inp))
	}

	outCh := c.ensureChannel(output)

	c.handlers = append(c.handlers, func(ctx context.Context) error {
		return fn(ctx, inChs, outCh)
	})

	return nil
}

func (c *conveyerImpl) Run(ctx context.Context) error {
	c.mu.Lock()
	if c.started {
		c.mu.Unlock()
		return errors.New("Run already started")
	}
	c.started = true
	c.mu.Unlock()

	var wg sync.WaitGroup
	errChan := make(chan error, len(c.handlers))

	// запуск обработчиков
	for _, h := range c.handlers {
		wg.Add(1)
		go func(hh func(context.Context) error) {
			defer wg.Done()
			if err := hh(ctx); err != nil {
				select {
				case errChan <- err:
				default:
				}
			}
		}(h)
	}

	// ожидание завершения
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		c.stop()
		return ctx.Err()

	case err := <-errChan:
		c.stop()
		return err

	case <-done:
		c.stop()
		return nil
	}
}

func (c *conveyerImpl) stop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.stopped {
		return
	}
	c.stopped = true

	for _, ch := range c.channels {
		close(ch)
	}
}

func (c *conveyerImpl) Send(input, data string) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.stopped {
		return errors.New("conveyer stopped")
	}

	ch, ok := c.channels[input]
	if !ok {
		return errors.New("chan not found")
	}

	select {
	case ch <- data:
	default: // по заданию — если буфер заполнен, молча игнорировать
	}
	return nil
}

func (c *conveyerImpl) Recv(output string) (string, error) {
	c.mu.RLock()
	ch, ok := c.channels[output]
	c.mu.RUnlock()

	if !ok {
		return "", errors.New("chan not found")
	}

	data, ok := <-ch
	if !ok {
		// канал закрыт, данных больше нет
		return Undefined, nil
	}

	return data, nil
}
