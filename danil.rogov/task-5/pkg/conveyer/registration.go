package conveyer

import "context"

func (conveyer *Conveyer) RegisterDecorator(
	decoratorFunc func(context.Context, chan string, chan string) error,
	input string,
	output string,
) {
	conveyer.ensureChan(input)
	conveyer.ensureChan(output)

	conveyer.tasks = append(conveyer.tasks, task{
		name:        "decorator",
		function:    decoratorFunc,
		inputChans:  []string{input},
		outputChans: []string{output},
	})
}

func (conveyer *Conveyer) RegisterMultiplexer(
	multiplexerFunc func(context.Context, []chan string, chan string) error,
	inputs []string,
	output string,
) {
	for _, name := range inputs {
		conveyer.ensureChan(name)
	}

	conveyer.ensureChan(output)

	conveyer.tasks = append(conveyer.tasks, task{
		name:        "multiplexer",
		function:    multiplexerFunc,
		inputChans:  inputs,
		outputChans: []string{output},
	})
}

func (conveyer *Conveyer) RegisterSeparator(
	separatorFunc func(context.Context, chan string, []chan string) error,
	input string,
	outputs []string,
) {
	conveyer.ensureChan(input)

	for _, chanName := range outputs {
		conveyer.ensureChan(chanName)
	}

	conveyer.tasks = append(conveyer.tasks, task{
		name:        "separator",
		function:    separatorFunc,
		inputChans:  []string{input},
		outputChans: outputs,
	})
}
