package conveyer

import (
	"context"
	"errors"
)

var errInvalidMultiplexerFnType = errors.New("invalid multiplexer function type")

func (c *conveyerImpl) RegisterMultiplexer(
	fn func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string,
	output string,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, inputName := range inputs {
		c.getOrCreateChannel(inputName)
	}

	c.getOrCreateChannel(output)
	c.handlers = append(c.handlers, handlerConfig{
		handlerType: handlerMultiplexer,
		fn:          fn,
		inputs:      inputs,
		outputs:     []string{output},
	})
}

func (c *conveyerImpl) runMultiplexer(
	ctx context.Context,
	handler handlerConfig,
	inputs []chan string,
	outputs []chan string,
) error {
	multiplexerfn, ok := handler.fn.(func(ctx context.Context, inputs []chan string, output chan string) error)
	if !ok {
		return errInvalidMultiplexerFnType
	}

	return multiplexerfn(ctx, inputs, outputs[0])
}
