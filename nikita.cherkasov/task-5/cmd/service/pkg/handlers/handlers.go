package handlers

import (
	"context"
	"errors"
	"strings"
)

var ErrNoDecorator = errors.New("can't be decorated")

const (
	noDecoratorData        = "no decorator"
	textForDecoratorString = "decorated: "
	noMultiplexerData      = "no multiplexer"
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	defer close(output)

	for {
		select {
		case <-ctx.Done():
			return nil

		case line, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(line, noDecoratorData) {
				return ErrNoDecorator
			}

			if !strings.HasPrefix(line, textForDecoratorString) {
				line = textForDecoratorString + line
			}

			select {
			case <-ctx.Done():
				return nil
			case output <- line:
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	defer func() {
		for _, outputChannel := range outputs {
			close(outputChannel)
		}
	}()

	outputCount := len(outputs)
	currentIndex := 0

	for {
		select {
		case <-ctx.Done():
			return nil

		case line, okay := <-input:
			if !okay {
				return nil
			}

			if outputCount == 0 {
				continue
			}

			select {
			case <-ctx.Done():
				return nil
			case outputs[currentIndex%outputCount] <- line:
				currentIndex++
			}
		}
	}
}
