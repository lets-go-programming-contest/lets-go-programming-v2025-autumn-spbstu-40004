package handlers

import (
	"context"
	"errors"
)

var ErrNoOutputChannels = errors.New("no output channels provided")

func SeparatorFunc(
	cntxt context.Context,
	input chan string,
	outputs []chan string,
) error {
	if len(outputs) == 0 {
		return ErrNoOutputChannels
	}

	currentOutput := 0

	for {
		select {
		case <-cntxt.Done():
			return nil
		case val, ok := <-input:
			if !ok {
				return nil
			}

			out := outputs[currentOutput%len(outputs)]

			out <- val

			currentOutput++
		}
	}
}
