package handlers

import (
	"context"
)

func SeparatorFunc(
	ctx context.Context,
	inputChannel chan string,
	outputChannels []chan string,
) error {
	if len(outputChannels) == 0 {
		return separatorReadHelper(ctx, inputChannel)
	}

	return separatorHelper(ctx, inputChannel, outputChannels)
}

func separatorReadHelper(
	ctx context.Context,
	inputChannel chan string,
) error {
	for {
		select {
		case <-ctx.Done():
			return nil

		case _, receivedOk := <-inputChannel:
			if !receivedOk {
				return nil
			}
		}
	}
}

func separatorHelper(
	ctx context.Context,
	inputChannel chan string,
	outputChannels []chan string,
) error {
	index := 0
	totalOutputs := len(outputChannels)

	for {
		select {
		case <-ctx.Done():
			return nil

		case data, receivedOk := <-inputChannel:
			if !receivedOk {
				for _, outputChannel := range outputChannels {
					close(outputChannel)
				}

				return nil
			}

			currentOutput := outputChannels[index]

			index++
			if index >= totalOutputs {
				index = 0
			}

			select {
			case <-ctx.Done():
				return nil
			case currentOutput <- data:
			}
		}
	}
}
