package conveyer

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	ErrChanNotFound    = errors.New("chan not found")
	ErrChanFull        = errors.New("channel is full")
	ErrChanClosed      = errors.New("channel closed")
	ErrUndefined       = errors.New("undefined")
	ErrCannotDecorate  = errors.New("can't be decorated")
)

type Pipeline struct {
	bufSize int

	mu    sync.RWMutex
	pipes map[string]chan string

	stages []Stage

	cancel context.CancelFunc
}

type Stage struct {
	sType   string
	fn      interface{}
	inputs  []string
	outputs []string
}

func New(bufSize int) *Pipeline {
	return &Pipeline{
		bufSize: bufSize,
		pipes:   make(map[string]chan string),
		stages:  make([]Stage, 0),
	}
}

func (p *Pipeline) ensurePipe(id string) chan string {
	p.mu.Lock()
	defer p.mu.Unlock()

	if ch, ok := p.pipes[id]; ok {
		return ch
	}
	ch := make(chan string, p.bufSize)
	p.pipes[id] = ch
	return ch
}

func (p *Pipeline) getPipe(id string) (chan string, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	ch, ok := p.pipes[id]
	if !ok {
		return nil, ErrChanNotFound
	}
	return ch, nil
}

func (p *Pipeline) RegisterDecorator(fn func(context.Context, chan string, chan string) error, in string, out string) {
	p.ensurePipe(in)
	p.ensurePipe(out)

	p.stages = append(p.stages, Stage{
		sType:   "decorator",
		fn:      fn,
		inputs:  []string{in},
		outputs: []string{out},
	})
}

func (p *Pipeline) RegisterMultiplexer(fn func(context.Context, []chan string, chan string) error, inputs []string, output string) {
	for _, in := range inputs {
		p.ensurePipe(in)
	}
	p.ensurePipe(output)

	p.stages = append(p.stages, Stage{
		sType:   "multiplexer",
		fn:      fn,
		inputs:  inputs,
		outputs: []string{output},
	})
}

func (p *Pipeline) RegisterSeparator(fn func(context.Context, chan string, []chan string) error, input string, outputs []string) {
	p.ensurePipe(input)
	for _, out := range outputs {
		p.ensurePipe(out)
	}

	p.stages = append(p.stages, Stage{
		sType:   "separator",
		fn:      fn,
		inputs:  []string{input},
		outputs: outputs,
	})
}

// Run запускает все стадии с использованием errgroup
func (p *Pipeline) Run(ctx context.Context) error {
	ctx, p.cancel = context.WithCancel(ctx)
	g, ctx := errgroup.WithContext(ctx)

	for _, st := range p.stages {
		st := st
		g.Go(func() error {
			return p.runStageSafe(ctx, st)
		})
	}

	return g.Wait() // дождёмся всех стадий или первой ошибки
}

func (p *Pipeline) runStageSafe(ctx context.Context, st Stage) error {
	switch st.sType {

	case "decorator":
		fn := st.fn.(func(context.Context, chan string, chan string) error)
		in, _ := p.getPipe(st.inputs[0])
		out, _ := p.getPipe(st.outputs[0])
		return fn(ctx, in, out)

	case "multiplexer":
		fn := st.fn.(func(context.Context, []chan string, chan string) error)

		ins := make([]chan string, len(st.inputs))
		for i, id := range st.inputs {
			ins[i], _ = p.getPipe(id)
		}
		out, _ := p.getPipe(st.outputs[0])

		return fn(ctx, ins, out)

	case "separator":
		fn := st.fn.(func(context.Context, chan string, []chan string) error)

		in, _ := p.getPipe(st.inputs[0])
		outs := make([]chan string, len(st.outputs))
		for i, id := range st.outputs {
			outs[i], _ = p.getPipe(id)
		}

		return fn(ctx, in, outs)
	}

	return nil
}

func (p *Pipeline) Send(pipe string, data string) error {
	ch, err := p.getPipe(pipe)
	if err != nil {
		return err
	}

	select {
	case ch <- data:
		return nil
	default:
		return ErrChanFull
	}
}

func (p *Pipeline) Recv(pipe string) (string, error) {
	ch, err := p.getPipe(pipe)
	if err != nil {
		return "", err
	}

	select {
	case v, ok := <-ch:
		if !ok {
			return "", ErrChanClosed
		}
		return v, nil
	default:
		return "", ErrUndefined
	}
}

func (p *Pipeline) Stop() {
	if p.cancel != nil {
		p.cancel()
	}

	p.mu.Lock()
	for k, ch := range p.pipes {
		close(ch)
		delete(p.pipes, k)
	}
	p.mu.Unlock()
}
