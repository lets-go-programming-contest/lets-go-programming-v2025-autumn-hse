package conveyer

import "context"

func (c *Conveyer) RegisterDecorator(
	decoratorFunc func(context.Context, <-chan string, chan<- string) error,
	input, output string) {
	c.getOrCreate(input)
	c.getOrCreate(output)

	c.tasks = append(c.tasks, task{
		name:    "decorator",
		fn:      decoratorFunc,
		inputs:  []string{input},
		outputs: []string{output},
	})
}

func (c *Conveyer) RegisterMultiplexer(
	multiplexerFunc func(context.Context, []<-chan string, chan<- string) error,
	inputs []string, output string) {
	for _, name := range inputs {
		c.getOrCreate(name)
	}
	c.getOrCreate(output)

	c.tasks = append(c.tasks, task{
		name:    "multiplexer",
		fn:      multiplexerFunc,
		inputs:  inputs,
		outputs: []string{output},
	})
}

func (c *Conveyer) RegisterSeparator(
	separatorFunc func(context.Context, <-chan string, []chan<- string) error,
	input string, outputs []string) {
	c.getOrCreate(input)
	for _, name := range outputs {
		c.getOrCreate(name)
	}

	c.tasks = append(c.tasks, task{
		name:    "separator",
		fn:      separatorFunc,
		inputs:  []string{input},
		outputs: outputs,
	})
}
