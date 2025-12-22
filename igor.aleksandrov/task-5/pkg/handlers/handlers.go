package handlers

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var (
	ErrCantBeDecorated = errors.New("can't be decorated")
	ErrInputsEmpty     = errors.New("inputs are empty")
	ErrOutputsEmpty    = errors.New("outputs are empty")
)

const (
	noDecoratorToken   = "no decorator"
	noMultiplexerToken = "no multiplexer"
	decoratedPrefix    = "decorated: "
)

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil

		case item, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(item, noDecoratorToken) {
				return ErrCantBeDecorated
			}

			result := item
			if !strings.HasPrefix(item, decoratedPrefix) {
				result = decoratedPrefix + item
			}

			select {
			case output <- result:
			case <-ctx.Done():
				return nil
			}
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	if len(outputs) == 0 {
		return ErrOutputsEmpty
	}

	var idx int

	for {
		select {
		case <-ctx.Done():
			return nil

		case item, ok := <-input:
			if !ok {
				return nil
			}

			targetCh := outputs[idx]

			select {
			case targetCh <- item:
			case <-ctx.Done():
				return nil
			}

			idx = (idx + 1) % len(outputs)
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	if len(inputs) == 0 {
		return ErrInputsEmpty
	}

	var waitGroup sync.WaitGroup

	readCh := func(channel chan string) {
		defer waitGroup.Done()

		for {
			select {
			case <-ctx.Done():
				return

			case item, ok := <-channel:
				if !ok {
					return
				}

				if strings.Contains(item, noMultiplexerToken) {
					continue
				}

				select {
				case output <- item:
				case <-ctx.Done():
					return
				}
			}
		}
	}

	waitGroup.Add(len(inputs))

	for _, channel := range inputs {
		go readCh(channel)
	}

	waitGroup.Wait()

	return nil
}
