package conveyer

import "context"

func (c *conveyerImpl) RegisterDecorator(
	fn func(
		ctx context.Context,
		input chan string,
		output chan string,
	) error,
	input string,
	output string,
) {
	c.m.Lock()
	defer c.m.Unlock()

	c.getChannel(input)
	c.getChannel(output)

	c.handlers = append(c.handlers, handlerConfig{
		handlerType: handlerDecorator,
		fn:          fn,
		input:       input,
		output:      output,
	})
}
