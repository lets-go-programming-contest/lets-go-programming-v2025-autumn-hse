package conveyer

import (
	"context"
	"fmt"
)

func (c *conveyer) runHandler(ctx context.Context, h handler) error {
	switch h.typ {
	case handlerTypeDecorator:
		return c.runDecorator(ctx, h)
	case handlerTypeMultiplexer:
		return c.runMultiplexer(ctx, h)
	case handlerTypeSeparator:
		return c.runSeparator(ctx, h)
	default:
		return fmt.Errorf("%w: %s", ErrUnknownHandlerType, h.typ)
	}
}

func (c *conveyer) runDecorator(ctx context.Context, h handler) error {
	decoratorFn, ok := h.fn.(func(context.Context, chan string, chan string) error)
	if !ok {
		return fmt.Errorf("%w: %T", ErrInvalidDecoratorType, h.fn)
	}

	inputCh, exists := c.getChannel(h.inputs[0])
	if !exists {
		return ErrChanNotFound
	}

	outputCh, exists := c.getChannel(h.outputs[0])
	if !exists {
		return ErrChanNotFound
	}

	return decoratorFn(ctx, inputCh, outputCh)
}

func (c *conveyer) runMultiplexer(ctx context.Context, h handler) error {
	multiplexerFn, ok := h.fn.(func(context.Context, []chan string, chan string) error)
	if !ok {
		return fmt.Errorf("%w: %T", ErrInvalidMultiplexerType, h.fn)
	}

	inputChs := make([]chan string, len(h.inputs))
	for i, input := range h.inputs {
		ch, exists := c.getChannel(input)
		if !exists {
			return ErrChanNotFound
		}

		inputChs[i] = ch
	}

	outputCh, exists := c.getChannel(h.outputs[0])
	if !exists {
		return ErrChanNotFound
	}

	return multiplexerFn(ctx, inputChs, outputCh)
}

func (c *conveyer) runSeparator(ctx context.Context, h handler) error {
	separatorFn, ok := h.fn.(func(context.Context, chan string, []chan string) error)
	if !ok {
		return fmt.Errorf("%w: %T", ErrInvalidSeparatorType, h.fn)
	}

	inputCh, exists := c.getChannel(h.inputs[0])
	if !exists {
		return ErrChanNotFound
	}

	outputChs := make([]chan string, len(h.outputs))
	for i, output := range h.outputs {
		ch, exists := c.getChannel(output)
		if !exists {
			return ErrChanNotFound
		}

		outputChs[i] = ch
	}

	return separatorFn(ctx, inputCh, outputChs)
}
