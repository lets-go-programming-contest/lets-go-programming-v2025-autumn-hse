package handlers

import (
	"errors"
	"strings"
)

func PrefixDecoratorFunc(input chan string, output chan string) error {
	for val := range input {
		if strings.Contains(val, "no decorator") {
			return errors.New("can't be decorated")
		}

		if !strings.HasPrefix(val, "decorated: ") {
			val = "decorated: " + val
		}

		output <- val
	}

	return nil
}

func SeparatorFunc(input chan string, outputs []chan string) error {
	index := 0
	for val := range input {
		outputs[index] <- val
		index = (index + 1) % len(outputs)
	}

	return nil
}

func MultiplexerFunc(inputs []chan string, output chan string) error {
	for {
		allClosed := true

		for _, ch := range inputs {
			val, ok := <-ch
			if ok {
				allClosed = false
				if strings.Contains(val, "no multiplexer") {
					continue
				}
				output <- val
			}
		}

		if allClosed {
			return nil
		}
	}
}
