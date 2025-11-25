package conveyer

import "context"

func (c *conveyerImpl) RegisterMultiplexer(
	fn func(
		ctx context.Context,
		inputs []chan string,
		output chan string,
	) error,
	inputs []string,
	output string,
) {
	c.m.Lock()
	defer c.m.Unlock()

	for _, inputName := range inputs {
		c.getChannel(inputName)
	}

	c.getChannel(output)
	c.handlers = append(c.handlers, handlerConfig{
		handlerType: handlerMultiplexer,
		fn:          fn,
		inputs:      inputs,
		output:      output,
	})
}
