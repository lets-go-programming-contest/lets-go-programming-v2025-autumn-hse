package conveyer

import (
	"context"
	"errors"
	"sync"

	"golang.org/x/sync/errgroup"
)

var (
	ErrPipeNotFound = errors.New("pipe not found")
	ErrPipeFull     = errors.New("pipe is full")
	ErrPipeClosed   = errors.New("pipe closed")
	ErrNoData       = errors.New("no data available")
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
		return nil, ErrPipeNotFound
	}
	return ch, nil
}

func (p *Pipeline) RegisterDecorator(fn func(context.Context, chan string, chan string) error, in, out string) {
	p.ensurePipe(in)
	p.ensurePipe(out)
	p.stages = append(p.stages, Stage{
		sType:   "decorator",
		fn:      fn,
		inputs:  []string{in},
		outputs: []string{out},
	})
}

func (p *Pipeline) RegisterMultiplexer(fn func(context.Context, []chan string, chan string) error, ins []string, out string) {
	for _, id := range ins {
		p.ensurePipe(id)
	}
	p.ensurePipe(out)
	p.stages = append(p.stages, Stage{
		sType:   "multiplexer",
		fn:      fn,
		inputs:  ins,
		outputs: []string{out},
	})
}

func (p *Pipeline) RegisterSeparator(fn func(context.Context, chan string, []chan string) error, in string, outs []string) {
	p.ensurePipe(in)
	for _, id := range outs {
		p.ensurePipe(id)
	}
	p.stages = append(p.stages, Stage{
		sType:   "separator",
		fn:      fn,
		inputs:  []string{in},
		outputs: outs,
	})
}

// Run запускает все стадии через errgroup
func (p *Pipeline) Run(ctx context.Context) error {
	ctx, p.cancel = context.WithCancel(ctx)
	g, ctx := errgroup.WithContext(ctx)

	for _, st := range p.stages {
		st := st // локальная копия
		g.Go(func() error {
			return p.runStage(ctx, st)
		})
	}

	err := g.Wait()

	// После завершения закрываем все каналы
	p.mu.Lock()
	for id, ch := range p.pipes {
		close(ch)
		delete(p.pipes, id)
	}
	p.mu.Unlock()

	return err
}

func (p *Pipeline) runStage(ctx context.Context, st Stage) error {
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

func (p *Pipeline) Send(pipe, data string) error {
	ch, err := p.getPipe(pipe)
	if err != nil {
		return err
	}
	select {
	case ch <- data:
		return nil
	default:
		return ErrPipeFull
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
			return "", ErrPipeClosed
		}
		return v, nil
	default:
		return "", ErrNoData
	}
}

func (p *Pipeline) Stop() {
	if p.cancel != nil {
		p.cancel()
	}
}
