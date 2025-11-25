package conveyer

import "context"

func (c *conveyerImpl) RegisterSeparator(
	fn func(
		ctx context.Context,
		input chan string,
		outputs []chan string,
	) error,
	input string,
	outputs []string,
) {
	c.m.Lock()
	defer c.m.Unlock()

	c.getChannel(input)
	for _, outputName := range outputs {
		c.getChannel(outputName)
	}
	c.handlers = append(c.handlers, handlerConfig{
		handlerType: handlerSeparator,
		fn:          fn,
		input:       input,
		outputs:     outputs,
	})
}
