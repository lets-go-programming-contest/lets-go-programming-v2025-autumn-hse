package conveyer

import "context"

func (c *conveyerImpl) ensureChannel(id string) {
	if _, ok := c.chans[id]; !ok {
		c.chans[id] = make(chan string, c.size)
	}
}

func (c *conveyerImpl) RegisterDecorator(
	fnHandler func(context.Context, chan string, chan string) error,
	input string,
	output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.ensureChannel(input)
	c.ensureChannel(output)

	c.handlers = append(c.handlers, handler{
		kind:          hDecorator,
		fnDecorator:   fnHandler,
		fnMultiplexer: nil,
		fnSeparator:   nil,
		inputIDs:      []string{input},
		outputIDs:     []string{output},
	})
}

func (c *conveyerImpl) RegisterMultiplexer(
	fnHandler func(context.Context, []chan string, chan string) error,
	inputs []string,
	output string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, id := range inputs {
		c.ensureChannel(id)
	}

	c.ensureChannel(output)

	c.handlers = append(c.handlers, handler{
		kind:          hMultiplexer,
		fnDecorator:   nil,
		fnMultiplexer: fnHandler,
		fnSeparator:   nil,
		inputIDs:      inputs,
		outputIDs:     []string{output},
	})
}

func (c *conveyerImpl) RegisterSeparator(
	fnHandler func(context.Context, chan string, []chan string) error,
	input string,
	outputs []string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.ensureChannel(input)

	for _, id := range outputs {
		c.ensureChannel(id)
	}

	c.handlers = append(c.handlers, handler{
		kind:          hSeparator,
		fnDecorator:   nil,
		fnMultiplexer: nil,
		fnSeparator:   fnHandler,
		inputIDs:      []string{input},
		outputIDs:     outputs,
	})
}
