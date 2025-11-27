package conveyer

import "context"

func (c *conveyerImpl) createChanIfNotExist(id string) {
	if _, ok := c.chans[id]; !ok {
		c.chans[id] = make(chan string, c.size)
	}
}

func (c *conveyerImpl) RegisterDecorator(
	fn func(context.Context, chan string, chan string) error,
	input string,
	output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.createChanIfNotExist(input)
	c.createChanIfNotExist(output)

	c.handlers = append(c.handlers, handler{
		kind:        hDecorator,
		fnDecorator: fn,
		inputIDs:    []string{input},
		outputIDs:   []string{output},
	})
}

func (c *conveyerImpl) RegisterMultiplexer(
	fn func(context.Context, []chan string, chan string) error,
	inputs []string,
	output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, id := range inputs {
		c.createChanIfNotExist(id)
	}
	c.createChanIfNotExist(output)

	c.handlers = append(c.handlers, handler{
		kind:          hMultiplexer,
		fnMultiplexer: fn,
		inputIDs:      inputs,
		outputIDs:     []string{output},
	})
}

func (c *conveyerImpl) RegisterSeparator(
	fn func(context.Context, chan string, []chan string) error,
	input string,
	outputs []string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.createChanIfNotExist(input)
	for _, id := range outputs {
		c.createChanIfNotExist(id)
	}

	c.handlers = append(c.handlers, handler{
		kind:        hSeparator,
		fnSeparator: fn,
		inputIDs:    []string{input},
		outputIDs:   outputs,
	})
}
