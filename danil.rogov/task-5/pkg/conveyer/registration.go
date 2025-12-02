package conveyer

import "context"

func (c *Conveyer) RegisterDecorator(
	decoratorFunc func(context.Context, chan string, chan string) error,
	input string,
	output string,
) {
	c.getOrCreate(input)
	c.getOrCreate(output)

	c.tasks = append(c.tasks, task{
		name:        "decorator",
		function:    decoratorFunc,
		inputChans:  []string{input},
		outputChans: []string{output},
	})
}

func (c *Conveyer) RegisterMultiplexer(
	multiplexerFunc func(context.Context, []chan string, chan string) error,
	inputs []string,
	output string,
) {
	for _, name := range inputs {
		c.getOrCreate(name)
	}

	c.getOrCreate(output)

	c.tasks = append(c.tasks, task{
		name:        "multiplexer",
		function:    multiplexerFunc,
		inputChans:  inputs,
		outputChans: []string{output},
	})
}

func (c *Conveyer) RegisterSeparator(
	separatorFunc func(context.Context, chan string, []chan string) error,
	input string,
	outputs []string,
) {
	c.getOrCreate(input)
	for _, chanName := range outputs {
		c.getOrCreate(chanName)
	}

	c.tasks = append(c.tasks, task{
		name:        "separator",
		function:    separatorFunc,
		inputChans:  []string{input},
		outputChans: outputs,
	})
}
