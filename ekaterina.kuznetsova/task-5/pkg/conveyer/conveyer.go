package conveyer

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Conveyer struct {
	channels map[string]chan string
	handlers []func(ctx context.Context) error
	size     int
	mu       sync.Mutex
	closed   bool
}

func New(size int) *Conveyer {
	return &Conveyer{
		channels: make(map[string]chan string),
		size:     size,
	}
}

func (conv *Conveyer) get(name string) chan string {
	conv.mu.Lock()
	defer conv.mu.Unlock()

	ch, ok := conv.channels[name]
	if !ok {
		ch = make(chan string, conv.size)
		conv.channels[name] = ch
	}

	return ch
}

var closeRegistry = struct {
	mu   sync.Mutex
	once map[chan string]*sync.Once
}{
	once: make(map[chan string]*sync.Once),
}

func safeClose(ch chan string) {
	closeRegistry.mu.Lock()
	o, ok := closeRegistry.once[ch]
	if !ok {
		o = &sync.Once{}
		closeRegistry.once[ch] = o
	}
	closeRegistry.mu.Unlock()

	o.Do(func() { close(ch) })
}

func (conv *Conveyer) closeAll() {
	conv.mu.Lock()
	defer conv.mu.Unlock()

	if conv.closed {
		return
	}
	conv.closed = true

	for _, ch := range conv.channels {
		safeClose(ch)
	}
}

func (conv *Conveyer) Run(ctx context.Context) error {
	defer conv.closeAll()

	errGroup, gCtx := errgroup.WithContext(ctx)

	for _, handler := range conv.handlers {
		h := handler

		runHandler := func() error {
			return h(gCtx)
		}
		errGroup.Go(runHandler)
	}

	return errGroup.Wait()
}

func (conv *Conveyer) Send(input string, data string) error {
	conv.mu.Lock()
	ch, ok := conv.channels[input]
	conv.mu.Unlock()

	if !ok {
		return errors.New("chan not found")
	}

	select {
	case ch <- data:
		return nil
	case <-context.Background().Done():
		return errors.New("send canceled")
	}
}

func (conv *Conveyer) Recv(output string) (string, error) {
	conv.mu.Lock()
	ch, ok := conv.channels[output]
	conv.mu.Unlock()

	if !ok {
		return "", errors.New("chan not found")
	}

	val, ok_ := <-ch
	if !ok_ {
		return "undefined", nil
	}

	return val, nil
}

func (conv *Conveyer) RegisterDecorator(
	fn func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error,
	input string,
	output string,
) {
	in := conv.get(input)
	out := conv.get(output)

	handler := func(ctx context.Context) error {
		return fn(ctx, in, out)
	}

	conv.handlers = append(conv.handlers, handler)
}

func (conv *Conveyer) RegisterMultiplexer(
	fn func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string,
	output string,
) {
	out := conv.get(output)

	var inCh []chan string
	for _, name := range inputs {
		inCh = append(inCh, conv.get(name))
	}

	handler := func(ctx context.Context) error {
		return fn(ctx, inCh, out)
	}

	conv.handlers = append(conv.handlers, handler)

}

func (conv *Conveyer) RegisterSeparator(
	fn func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string,
	outputs []string,
) {
	in := conv.get(input)

	var outCh []chan string

	for _, name := range outputs {
		outCh = append(outCh, conv.get(name))
	}

	handler := func(ctx context.Context) error {
		return fn(ctx, in, outCh)
	}

	conv.handlers = append(conv.handlers, handler)

}
