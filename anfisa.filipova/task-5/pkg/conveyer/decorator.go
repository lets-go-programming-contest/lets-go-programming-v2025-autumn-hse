package conveyer

import (
	"context"
	"fmt"
)

func (c *conveyerImpl) RegisterDecorator(
	fn func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error,
	input string,
	output string,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.getOrCreateChannel(input)
	c.getOrCreateChannel(output)

	c.handlers = append(c.handlers, handlerConfig{
		handlerType: handlerDecorator,
		fn:          fn,
		inputs:      []string{input},
		outputs:     []string{output},
	})
}

func (c *conveyerImpl) runDecorator(
	ctx context.Context,
	handler handlerConfig,
	inputs []chan string,
	outputs []chan string,
) error {
	decoratorfn, ok := handler.fn.(func(ctx context.Context, input chan string, output chan string) error)
	if !ok {
		return fmt.Errorf("invalid decorator function type")
	}

	return decoratorfn(ctx, inputs[0], outputs[0])
}
