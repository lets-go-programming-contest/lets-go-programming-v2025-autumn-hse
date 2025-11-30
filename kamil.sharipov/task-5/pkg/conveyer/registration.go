package conveyer

import "context"

func (c *conveyer) RegisterDecorator(
	decoratorFn func(ctx context.Context, input chan string, output chan string) error,
	input string,
	output string,
) {
	c.ensureChannel(input)
	c.ensureChannel(output)

	c.handlers = append(c.handlers, handler{
		typ:     handlerTypeDecorator,
		fn:      decoratorFn,
		inputs:  []string{input},
		outputs: []string{output},
	})
}

func (c *conveyer) RegisterMultiplexer(
	multiplexerFn func(ctx context.Context, inputs []chan string, output chan string) error,
	inputs []string,
	output string,
) {
	for _, input := range inputs {
		c.ensureChannel(input)
	}

	c.ensureChannel(output)

	c.handlers = append(c.handlers, handler{
		typ:     handlerTypeMultiplexer,
		fn:      multiplexerFn,
		inputs:  inputs,
		outputs: []string{output},
	})
}

func (c *conveyer) RegisterSeparator(
	separatorFn func(ctx context.Context, input chan string, outputs []chan string) error,
	input string,
	outputs []string,
) {
	c.ensureChannel(input)

	for _, output := range outputs {
		c.ensureChannel(output)
	}

	c.handlers = append(c.handlers, handler{
		typ:     handlerTypeSeparator,
		fn:      separatorFn,
		inputs:  []string{input},
		outputs: outputs,
	})
}
