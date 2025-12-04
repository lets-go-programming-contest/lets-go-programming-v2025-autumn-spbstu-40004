package handlers

import (
	"context"
	"errors"
	"strings"
)

var (
	errCantBeDecorated = errors.New("can't be decorated")
	errInputsEmpty     = errors.New("inputs are empty")
	errOutputsEmpty    = errors.New("outputs are empty")
)

const (
	noDecoratorSub   = "no decorator"
	noMultiplexerSub = "no multiplexer"
	decoratedSub     = "decorated: "
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, noDecoratorSub) {
				return errCantBeDecorated
			}

			result := data
			if !strings.HasPrefix(data, decoratedSub) {
				result = decoratedSub + data
			}

			select {
			case <-ctx.Done():
				return nil
			case output <- result:
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	outputLength := len(outputs)

	if outputLength == 0 {
		return errOutputsEmpty
	}

	currentIndex := 0

	for {
		select {
		case data, ok := <-input:
			if !ok {
				return nil
			}

			select {
			case <-ctx.Done():
				return nil
			case outputs[currentIndex] <- data:
			}

			currentIndex = (currentIndex + 1) % outputLength
		case <-ctx.Done():
			return nil
		}
	}
}
