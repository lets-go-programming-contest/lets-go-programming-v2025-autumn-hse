package conveyer

import "context"

func (conveyer *Conveyer) executeDecorator(ctx context.Context, task task) error {
	if len(task.inputChans) == 0 || len(task.outputChans) == 0 {
		return ErrChanNotFound
	}

	decoratorFunc, ok := task.function.(func(context.Context, chan string, chan string) error)
	if !ok {
		return ErrInvalidDecoratorType
	}

	inputChan := conveyer.getOrCreate(task.inputChans[0])
	outputChan := conveyer.getOrCreate(task.outputChans[0])

	return decoratorFunc(ctx, inputChan, outputChan)
}

func (conveyer *Conveyer) executeMultiplexer(ctx context.Context, task task) error {
	multiplexerFunc, ok := task.function.(func(context.Context, []chan string, chan string) error)
	if !ok {
		return ErrInvalidMultiplexerType
	}

	inputChans := make([]chan string, len(task.inputChans))

	for index, name := range task.inputChans {
		inputChans[index] = conveyer.getOrCreate(name)
	}

	outputChan := conveyer.getOrCreate(task.outputChans[0])

	return multiplexerFunc(ctx, inputChans, outputChan)
}

func (conveyer *Conveyer) executeSeparator(ctx context.Context, taskItem task) error {
	separatorFunc, ok := taskItem.function.(func(context.Context, chan string, []chan string) error)
	if !ok {
		return ErrInvalidSeparatorType
	}

	outputChans := make([]chan string, len(taskItem.outputChans))

	for index, name := range taskItem.outputChans {
		outputChans[index] = conveyer.getOrCreate(name)
	}

	inputChan := conveyer.getOrCreate(taskItem.inputChans[0])

	return separatorFunc(ctx, inputChan, outputChans)
}
