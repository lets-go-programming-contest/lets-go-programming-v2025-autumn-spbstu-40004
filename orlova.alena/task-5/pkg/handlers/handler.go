package handlers

import (
	"context"
	"errors"
	"strings"
)

var ErrCannotDecorate = errors.New("can't be decorated")

func PrefixDecoratorFunc(ctx context.Context, input chan string, output chan string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			if strings.Contains(data, "no decorator") {
				return ErrCannotDecorate
			}

			if !strings.HasPrefix(data, "decorated: ") {
				data = "decorated: " + data
			}

			output <- data
		}
	}
}

func SeparatorFunc(ctx context.Context, input chan string, outputs []chan string) error {
	iter := 0

	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-input:
			if !ok {
				return nil
			}

			outputs[iter%len(outputs)] <- data

			iter++
		}
	}
}

func MultiplexerFunc(ctx context.Context, inputs []chan string, output chan string) error {
	merged := make(chan string)

	for _, inp := range inputs {
		go func(ch chan string) {
			for {
				select {
				case <-ctx.Done():
					return
				case data, ok := <-ch:
					if !ok {
						return
					}

					merged <- data
				}
			}
		}(inp)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case data, ok := <-merged:
			if !ok {
				return nil
			}

			if strings.Contains(data, "no multiplexer") {
				continue
			}

			output <- data
		}
	}
}
