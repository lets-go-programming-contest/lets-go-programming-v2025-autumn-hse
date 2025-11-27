package conveyer

import (
	"context"
	"fmt"
)

func (c *conveyerImpl) RegisterSeparator(
	fn func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string,
	outputs []string,
) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.getOrCreateChannel(input)
	for _, outputName := range outputs {
		c.getOrCreateChannel(outputName)
	}
	c.handlers = append(c.handlers, handlerConfig{
		handlerType: handlerSeparator,
		fn:          fn,
		inputs:      []string{input},
		outputs:     outputs,
	})
}

func (c *conveyerImpl) runSeparator(
	ctx context.Context,
	handler handlerConfig,
	inputs []chan string,
	outputs []chan string,
) error {
	separatorfn, ok := handler.fn.(func(ctx context.Context, input chan string, outputs []chan string) error)
	if !ok {
		return fmt.Errorf("invalid separator function type")
	}

	return separatorfn(ctx, inputs[0], outputs)
}
