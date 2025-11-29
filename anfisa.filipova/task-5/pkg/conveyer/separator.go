package conveyer

import (
	"context"
	"errors"
)

var errInvalidSeparatorFnType = errors.New("invalid separator function type")

func (c *Conveyer) RegisterSeparator(
	separatorfn func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string,
	outputs []string,
) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	c.prepareChannel(input)

	for _, outputName := range outputs {
		c.prepareChannel(outputName)
	}

	c.handlers = append(c.handlers, handlerConfig{
		handlerType: handlerSeparator,
		fn:          separatorfn,
		inputs:      []string{input},
		outputs:     outputs,
	})
}

func (c *Conveyer) runSeparator(
	ctx context.Context,
	handler handlerConfig,
	inputs []chan string,
	outputs []chan string,
) error {
	separatorfn, ok := handler.fn.(func(ctx context.Context, input chan string, outputs []chan string) error)
	if !ok {
		return errInvalidSeparatorFnType
	}

	return separatorfn(ctx, inputs[0], outputs)
}
