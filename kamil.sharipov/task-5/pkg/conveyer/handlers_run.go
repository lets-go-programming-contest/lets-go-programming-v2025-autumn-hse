package conveyer

import (
	"context"
	"fmt"
)

func (c *conveyer) runHandler(ctx context.Context, hnd handler) error {
	switch hnd.typ {
	case handlerTypeDecorator:
		return c.runDecorator(ctx, hnd)
	case handlerTypeMultiplexer:
		return c.runMultiplexer(ctx, hnd)
	case handlerTypeSeparator:
		return c.runSeparator(ctx, hnd)
	default:
		return fmt.Errorf("%w: %s", ErrUnknownHandlerType, hnd.typ)
	}
}

func (c *conveyer) runDecorator(ctx context.Context, hnd handler) error {
	decoratorFn, ok := hnd.fn.(func(context.Context, chan string, chan string) error)
	if !ok {
		return fmt.Errorf("%w: %T", ErrInvalidDecoratorType, hnd.fn)
	}

	inputCh, exists := c.getChannel(hnd.inputs[0])
	if !exists {
		return ErrChanNotFound
	}

	outputCh, exists := c.getChannel(hnd.outputs[0])
	if !exists {
		return ErrChanNotFound
	}

	return decoratorFn(ctx, inputCh, outputCh)
}

func (c *conveyer) runMultiplexer(ctx context.Context, hnd handler) error {
	multiplexerFn, ok := hnd.fn.(func(context.Context, []chan string, chan string) error)
	if !ok {
		return fmt.Errorf("%w: %T", ErrInvalidMultiplexerType, hnd.fn)
	}

	inputChs := make([]chan string, len(hnd.inputs))

	for index, input := range hnd.inputs {
		ch, exists := c.getChannel(input)
		if !exists {
			return ErrChanNotFound
		}

		inputChs[index] = ch
	}

	outputCh, exists := c.getChannel(hnd.outputs[0])
	if !exists {
		return ErrChanNotFound
	}

	return multiplexerFn(ctx, inputChs, outputCh)
}

func (c *conveyer) runSeparator(ctx context.Context, hnd handler) error {
	separatorFn, ok := hnd.fn.(func(context.Context, chan string, []chan string) error)
	if !ok {
		return fmt.Errorf("%w: %T", ErrInvalidSeparatorType, hnd.fn)
	}

	inputCh, exists := c.getChannel(hnd.inputs[0])
	if !exists {
		return ErrChanNotFound
	}

	outputChs := make([]chan string, len(hnd.outputs))

	for index, output := range hnd.outputs {
		ch, exists := c.getChannel(output)
		if !exists {
			return ErrChanNotFound
		}

		outputChs[index] = ch
	}

	return separatorFn(ctx, inputCh, outputChs)
}
