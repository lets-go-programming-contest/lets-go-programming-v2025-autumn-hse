package conveyer

import (
	"context"
	"errors"
)

var errInvalidDecoratorFnType = errors.New("invalid decorator function type")

func (c *conveyerImpl) RegisterDecorator(
	decoratorfn func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error,
	input string,
	output string,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.prepareChannel(input)
	c.prepareChannel(output)

	c.handlers = append(c.handlers, handlerConfig{
		handlerType: handlerDecorator,
		fn:          decoratorfn,
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
		return errInvalidDecoratorFnType
	}

	return decoratorfn(ctx, inputs[0], outputs[0])
}
